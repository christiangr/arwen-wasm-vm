package host

import (
	"fmt"
	"sync"

	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/arwen"
	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/arwen/contexts"
	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/arwen/cryptoapi"
	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/arwen/elrondapi"
	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/config"
	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/crypto"
	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/crypto/factory"
	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/wasmer"
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	logger "github.com/ElrondNetwork/elrond-go-logger"
	"github.com/ElrondNetwork/elrond-vm-common"
)

var log = logger.GetOrCreate("arwen/host")

// MaximumWasmerInstanceCount represents the maximum number of Wasmer instances that can be active at the same time
var MaximumWasmerInstanceCount = uint64(10)

// TryFunction corresponds to the try() part of a try / catch block
type TryFunction func()

// CatchFunction corresponds to the catch() part of a try / catch block
type CatchFunction func(error)

// vmHost implements HostContext interface.
type vmHost struct {
	cryptoHook   crypto.VMCrypto
	mutExecution sync.RWMutex

	ethInput []byte

	blockchainContext   arwen.BlockchainContext
	runtimeContext      arwen.RuntimeContext
	outputContext       arwen.OutputContext
	meteringContext     arwen.MeteringContext
	storageContext      arwen.StorageContext
	managedTypesContext arwen.ManagedTypesContext

	gasSchedule          config.GasScheduleMap
	scAPIMethods         *wasmer.Imports
	builtInFuncContainer vmcommon.BuiltInFunctionContainer
	esdtTransferParser   vmcommon.ESDTTransferParser
}

// NewArwenVM creates a new Arwen vmHost
func NewArwenVM(
	blockChainHook vmcommon.BlockchainHook,
	hostParameters *arwen.VMHostParameters,
) (arwen.VMHost, error) {

	if check.IfNil(blockChainHook) {
		return nil, arwen.ErrNilBlockChainHook
	}
	if hostParameters == nil {
		return nil, arwen.ErrNilHostParameters
	}
	if check.IfNil(hostParameters.ESDTTransferParser) {
		return nil, arwen.ErrNilESDTTransferParser
	}
	if check.IfNil(hostParameters.BuiltInFuncContainer) {
		return nil, arwen.ErrNilBuiltInFunctionsContainer
	}

	cryptoHook := factory.NewVMCrypto()
	host := &vmHost{
		cryptoHook:           cryptoHook,
		meteringContext:      nil,
		runtimeContext:       nil,
		blockchainContext:    nil,
		storageContext:       nil,
		managedTypesContext:  nil,
		gasSchedule:          hostParameters.GasSchedule,
		scAPIMethods:         nil,
		builtInFuncContainer: hostParameters.BuiltInFuncContainer,
		esdtTransferParser:   hostParameters.ESDTTransferParser,
	}

	var err error

	imports, err := elrondapi.ElrondEIImports()
	if err != nil {
		return nil, err
	}

	imports, err = elrondapi.BigIntImports(imports)
	if err != nil {
		return nil, err
	}

	imports, err = elrondapi.SmallIntImports(imports)
	if err != nil {
		return nil, err
	}

	imports, err = elrondapi.ManagedBufferImports(imports)
	if err != nil {
		return nil, err
	}

	imports, err = cryptoapi.CryptoImports(imports)
	if err != nil {
		return nil, err
	}

	err = wasmer.SetImports(imports)
	if err != nil {
		return nil, err
	}

	host.scAPIMethods = imports

	host.blockchainContext, err = contexts.NewBlockchainContext(host, blockChainHook)
	if err != nil {
		return nil, err
	}

	host.runtimeContext, err = contexts.NewRuntimeContext(
		host,
		hostParameters.VMType,
		host.builtInFuncContainer,
	)
	if err != nil {
		return nil, err
	}

	host.meteringContext, err = contexts.NewMeteringContext(host, hostParameters.GasSchedule, hostParameters.BlockGasLimit)
	if err != nil {
		return nil, err
	}

	host.outputContext, err = contexts.NewOutputContext(host)
	if err != nil {
		return nil, err
	}

	host.storageContext, err = contexts.NewStorageContext(host, blockChainHook, hostParameters.ElrondProtectedKeyPrefix)
	if err != nil {
		return nil, err
	}

	host.managedTypesContext, err = contexts.NewManagedTypesContext(host)
	if err != nil {
		return nil, err
	}

	gasCostConfig, err := config.CreateGasConfig(host.gasSchedule)
	if err != nil {
		return nil, err
	}

	host.runtimeContext.SetMaxInstanceCount(MaximumWasmerInstanceCount)

	opcodeCosts := gasCostConfig.WASMOpcodeCost.ToOpcodeCostsArray()
	wasmer.SetOpcodeCosts(&opcodeCosts)

	host.initContexts()

	return host, nil
}

// GetVersion returns the Arwen version string
func (host *vmHost) GetVersion() string {
	return arwen.ArwenVersion
}

// Crypto returns the VMCrypto instance of the host
func (host *vmHost) Crypto() crypto.VMCrypto {
	return host.cryptoHook
}

// Blockchain returns the BlockchainContext instance of the host
func (host *vmHost) Blockchain() arwen.BlockchainContext {
	return host.blockchainContext
}

// Runtime returns the RuntimeContext instance of the host
func (host *vmHost) Runtime() arwen.RuntimeContext {
	return host.runtimeContext
}

// Output returns the OutputContext instance of the host
func (host *vmHost) Output() arwen.OutputContext {
	return host.outputContext
}

// Metering returns the MeteringContext instance of the host
func (host *vmHost) Metering() arwen.MeteringContext {
	return host.meteringContext
}

// Storage returns the StorageContext instance of the host
func (host *vmHost) Storage() arwen.StorageContext {
	return host.storageContext
}

// BigInt returns the BigIntContext instance of the host
func (host *vmHost) ManagedTypes() arwen.ManagedTypesContext {
	return host.managedTypesContext
}

// GetContexts returns the main contexts of the host
func (host *vmHost) GetContexts() (
	arwen.ManagedTypesContext,
	arwen.BlockchainContext,
	arwen.MeteringContext,
	arwen.OutputContext,
	arwen.RuntimeContext,
	arwen.StorageContext,
) {
	return host.managedTypesContext,
		host.blockchainContext,
		host.meteringContext,
		host.outputContext,
		host.runtimeContext,
		host.storageContext
}

// InitState resets the contexts of the host and reconfigures its flags
func (host *vmHost) InitState() {
	host.initContexts()
}

func (host *vmHost) initContexts() {
	host.ClearContextStateStack()
	host.managedTypesContext.InitState()
	host.outputContext.InitState()
	host.meteringContext.InitState()
	host.runtimeContext.InitState()
	host.storageContext.InitState()
	host.ethInput = nil
}

// ClearContextStateStack cleans the state stacks of all the contexts of the host
func (host *vmHost) ClearContextStateStack() {
	host.managedTypesContext.ClearStateStack()
	host.outputContext.ClearStateStack()
	host.meteringContext.ClearStateStack()
	host.runtimeContext.ClearStateStack()
	host.storageContext.ClearStateStack()
}

// Clean closes the currently running Wasmer instance
func (host *vmHost) Clean() {
	host.runtimeContext.CleanWasmerInstance()
}

// GetAPIMethods returns the EEI as a set of imports for Wasmer
func (host *vmHost) GetAPIMethods() *wasmer.Imports {
	return host.scAPIMethods
}

// GasScheduleChange applies a new gas schedule to the host
func (host *vmHost) GasScheduleChange(newGasSchedule config.GasScheduleMap) {
	host.mutExecution.Lock()
	defer host.mutExecution.Unlock()

	host.gasSchedule = newGasSchedule
	gasCostConfig, err := config.CreateGasConfig(newGasSchedule)
	if err != nil {
		log.Error("cannot apply new gas config", "err", err)
		return
	}

	opcodeCosts := gasCostConfig.WASMOpcodeCost.ToOpcodeCostsArray()
	wasmer.SetOpcodeCosts(&opcodeCosts)

	host.meteringContext.SetGasSchedule(newGasSchedule)
}

// GetGasScheduleMap returns the currently stored gas schedule
func (host *vmHost) GetGasScheduleMap() config.GasScheduleMap {
	return host.gasSchedule
}

// RunSmartContractCreate executes the deployment of a new contract
func (host *vmHost) RunSmartContractCreate(input *vmcommon.ContractCreateInput) (vmOutput *vmcommon.VMOutput, err error) {
	host.mutExecution.RLock()
	defer host.mutExecution.RUnlock()

	log.Trace("RunSmartContractCreate begin", "len(code)", len(input.ContractCode), "metadata", input.ContractCodeMetadata)

	try := func() {
		vmOutput = host.doRunSmartContractCreate(input)
	}

	catch := func(caught error) {
		err = caught
		log.Error("RunSmartContractCreate", "error", err)
	}

	TryCatch(try, catch, "arwen.RunSmartContractCreate")
	if vmOutput != nil {
		log.Trace("RunSmartContractCreate end", "returnCode", vmOutput.ReturnCode, "returnMessage", vmOutput.ReturnMessage)
	}

	return
}

// RunSmartContractCall executes the call of an existing contract
func (host *vmHost) RunSmartContractCall(input *vmcommon.ContractCallInput) (vmOutput *vmcommon.VMOutput, err error) {
	host.mutExecution.RLock()
	defer host.mutExecution.RUnlock()

	log.Trace("RunSmartContractCall begin", "function", input.Function)

	tryUpgrade := func() {
		vmOutput = host.doRunSmartContractUpgrade(input)
	}

	tryCall := func() {
		vmOutput = host.doRunSmartContractCall(input)
	}

	catch := func(caught error) {
		err = caught
		log.Error("RunSmartContractCall", "error", err)
	}

	isUpgrade := input.Function == arwen.UpgradeFunctionName
	if isUpgrade {
		TryCatch(tryUpgrade, catch, "arwen.RunSmartContractUpgrade")
	} else {
		TryCatch(tryCall, catch, "arwen.RunSmartContractCall")
	}

	return
}

// TryCatch simulates a try/catch block using golang's recover() functionality
func TryCatch(try TryFunction, catch CatchFunction, catchFallbackMessage string) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%s, panic: %v", catchFallbackMessage, r)
			}

			catch(err)
		}
	}()

	try()
}

// AreInSameShard returns true if the provided addresses are part of the same shard
func (host *vmHost) AreInSameShard(leftAddress []byte, rightAddress []byte) bool {
	blockchain := host.Blockchain()
	leftShard := blockchain.GetShardOfAddress(leftAddress)
	rightShard := blockchain.GetShardOfAddress(rightAddress)

	return leftShard == rightShard
}

// IsInterfaceNil returns true if there is no value under the interface
func (host *vmHost) IsInterfaceNil() bool {
	return host == nil
}

// SetRuntimeContext sets the runtimeContext for this host, used in tests
func (host *vmHost) SetRuntimeContext(runtime arwen.RuntimeContext) {
	host.runtimeContext = runtime
}

// GetRuntimeErrors obtains the cumultated error object after running the SC
func (host *vmHost) GetRuntimeErrors() error {
	if host.runtimeContext != nil {
		return host.runtimeContext.GetAllErrors()
	}
	return nil
}

// SetBuiltInFunctionsContainer sets the built in function container - only for testing
func (host *vmHost) SetBuiltInFunctionsContainer(builtInFuncs vmcommon.BuiltInFunctionContainer) {
	if check.IfNil(builtInFuncs) {
		return
	}
	host.builtInFuncContainer = builtInFuncs
}
