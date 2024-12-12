package blockinputs

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	gethstate "github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/ethereum/go-ethereum/triedb/hashdb"
	"github.com/kkrt-labs/kakarot-controller/pkg/ethereum"
	"github.com/kkrt-labs/kakarot-controller/pkg/ethereum/evm"
	"github.com/kkrt-labs/kakarot-controller/pkg/ethereum/state"
	"github.com/kkrt-labs/kakarot-controller/pkg/log"
	"github.com/kkrt-labs/kakarot-controller/pkg/tag"
)

// Executor is the interface for EVM execution on provable inputs.
// It runs a full "execution + final state validation" of the block
// It is primarily meant to validate that the provable inputs are correct and enable proper EVM execution.
type Executor interface {
	// Execute runs a full EVM block execution on provable inputs
	Execute(ctx context.Context, inputs *ProvableInputs) (*core.ProcessResult, error)
}

type BaseExecutor struct{}

// NewExecutor creates a new instance of the BaseExecutor.
func NewExecutor() *BaseExecutor {
	return &BaseExecutor{}
}

// Execute runs the ProvableBlockInputs data for the EVM prover engine.
func (e *BaseExecutor) Execute(ctx context.Context, inputs *ProvableInputs) (*core.ProcessResult, error) {
	ctx = tag.WithComponent(ctx, "execute")
	ctx = tag.WithTags(
		ctx,
		tag.Key("chain.id").String(inputs.ChainConfig.ChainID.String()),
		tag.Key("block.number").Int64(inputs.Block.Number.ToInt().Int64()),
		tag.Key("block.hash").String(inputs.Block.Hash.Hex()),
	)

	res, err := e.execute(ctx, inputs)
	if err != nil {
		log.LoggerFromContext(ctx).WithError(err).Errorf("Provable execution failed")
		return res, err
	}

	log.LoggerFromContext(ctx).Infof("Provable execution succeeded")

	return res, err
}

type executorContext struct {
	ctx     context.Context
	stateDB gethstate.Database
	hc      *core.HeaderChain
}

func (e *BaseExecutor) execute(ctx context.Context, inputs *ProvableInputs) (*core.ProcessResult, error) {
	log.LoggerFromContext(ctx).Infof("Process provable execution...")

	execCtx, err := e.prepareContext(ctx, inputs)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare execution context: %v", err)
	}

	err = e.preparePreState(execCtx, inputs)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare pre-state: %v", err)
	}

	execParams, err := e.prepareExecParams(execCtx, inputs)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare execution exec params: %v", err)
	}

	return e.execEVM(execCtx, execParams)
}

func (e *BaseExecutor) prepareContext(ctx context.Context, inputs *ProvableInputs) (*executorContext, error) {
	log.LoggerFromContext(ctx).Debug("Prepare context...")

	// --- Create necessary database and chain instances ---
	db := rawdb.NewMemoryDatabase()
	trieDB := triedb.NewDatabase(db, &triedb.Config{HashDB: &hashdb.Config{}})
	stateDB := state.NewModifiedTrieDatabase(trieDB, nil) // We use a modified trie database to track trie modifications

	hc, err := ethereum.NewChain(inputs.ChainConfig, stateDB)
	if err != nil {
		return nil, fmt.Errorf("failed to create chain: %v", err)
	}

	return &executorContext{
		ctx:     ctx,
		stateDB: stateDB,
		hc:      hc,
	}, nil
}

func (e *BaseExecutor) preparePreState(ctx *executorContext, inputs *ProvableInputs) error {
	log.LoggerFromContext(ctx.ctx).Info("Prepare pre-state...")

	// -- Preload the ancestors of the block into database ---
	ethereum.WriteHeaders(ctx.stateDB.TrieDB().Disk(), inputs.Ancestors...)

	// --- Preload the account bytecodes into the database ---
	codes := make([][]byte, 0)
	for _, code := range inputs.Codes {
		codes = append(codes, code)
	}
	ethereum.WriteCodes(ctx.stateDB.TrieDB().Disk(), codes...)

	// -- Preload the pre-state nodes to database ---
	nodes := make([][]byte, 0)
	for _, node := range inputs.PreState {
		b, err := hexutil.Decode(node)
		if err != nil {
			return fmt.Errorf("failed to decode node %v: %v", node, err)
		}
		nodes = append(nodes, b)
	}
	ethereum.WriteNodesToHashDB(ctx.stateDB.TrieDB().Disk(), nodes...)

	return nil
}

func (e *BaseExecutor) prepareExecParams(ctx *executorContext, inputs *ProvableInputs) (*evm.ExecParams, error) {
	log.LoggerFromContext(ctx.ctx).Debug("Prepare execution parameters...")

	parentHeader := inputs.Ancestors[0]
	preState, err := gethstate.New(parentHeader.Root, ctx.stateDB)
	if err != nil {
		return nil, fmt.Errorf("failed to create pre-state from parent root %v: %v", parentHeader.Root, err)
	}

	return &evm.ExecParams{
		VMConfig: &vm.Config{
			StatelessSelfValidation: true,
		},
		Block:    inputs.Block.Block(),
		Validate: true, // We validate the block execution to ensure the result and final state are correct
		Chain:    ctx.hc,
		State:    preState,
	}, nil
}

func (e *BaseExecutor) execEVM(ctx *executorContext, execParams *evm.ExecParams) (*core.ProcessResult, error) {
	log.LoggerFromContext(ctx.ctx).Infof("Execute EVM...")

	res, err := evm.ExecutorWithTags("evm")(evm.ExecutorWithLog()(evm.NewExecutor())).Execute(ctx.ctx, execParams)
	if err != nil {
		return res, fmt.Errorf("failed to execute block: %v", err)
	}

	return res, nil
}
