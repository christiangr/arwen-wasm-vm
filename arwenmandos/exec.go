package arwenmandos

import (
	"fmt"

	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/arwen"
	arwenHost "github.com/ElrondNetwork/arwen-wasm-vm/v1_4/arwen/host"
	gasSchedules "github.com/ElrondNetwork/arwen-wasm-vm/v1_4/arwenmandos/gasSchedules"
	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/config"
	mc "github.com/ElrondNetwork/arwen-wasm-vm/v1_4/mandos-go/controller"
	er "github.com/ElrondNetwork/arwen-wasm-vm/v1_4/mandos-go/expression/reconstructor"
	fr "github.com/ElrondNetwork/arwen-wasm-vm/v1_4/mandos-go/fileresolver"
	mj "github.com/ElrondNetwork/arwen-wasm-vm/v1_4/mandos-go/json/model"
	worldhook "github.com/ElrondNetwork/arwen-wasm-vm/v1_4/mock/world"
	logger "github.com/ElrondNetwork/elrond-go-logger"
	vmi "github.com/ElrondNetwork/elrond-vm-common"
	"github.com/ElrondNetwork/elrond-vm-common/parsers"
)

var log = logger.GetOrCreate("arwen/mandos")

// TestVMType is the VM type argument we use in tests.
var TestVMType = []byte{0, 0}

// ArwenTestExecutor parses, interprets and executes both .test.json tests and .scen.json scenarios with Arwen.
type ArwenTestExecutor struct {
	World                   *worldhook.MockWorld
	vm                      vmi.VMExecutionHandler
	checkGas                bool
	mandosGasScheduleLoaded bool
	fileResolver            fr.FileResolver
	exprReconstructor       er.ExprReconstructor
}

var _ mc.TestExecutor = (*ArwenTestExecutor)(nil)
var _ mc.ScenarioExecutor = (*ArwenTestExecutor)(nil)

// NewArwenTestExecutor prepares a new ArwenTestExecutor instance.
func NewArwenTestExecutor() (*ArwenTestExecutor, error) {
	world := worldhook.NewMockWorld()

	gasScheduleMap := config.MakeGasMapForTests()
	err := world.InitBuiltinFunctions(gasScheduleMap)
	if err != nil {
		return nil, err
	}

	blockGasLimit := uint64(10000000)
	esdtTransferParser, _ := parsers.NewESDTTransferParser(worldhook.WorldMarshalizer)
	vm, err := arwenHost.NewArwenVM(world, &arwen.VMHostParameters{
		VMType:                   TestVMType,
		BlockGasLimit:            blockGasLimit,
		GasSchedule:              gasScheduleMap,
		BuiltInFuncContainer:     world.BuiltinFuncs.Container,
		ElrondProtectedKeyPrefix: []byte(ElrondProtectedKeyPrefix),
		ESDTTransferParser:       esdtTransferParser,
	})
	if err != nil {
		return nil, err
	}

	return &ArwenTestExecutor{
		World:                   world,
		vm:                      vm,
		checkGas:                true,
		mandosGasScheduleLoaded: false,
		fileResolver:            nil,
		exprReconstructor:       er.ExprReconstructor{},
	}, nil
}

// GetVM yields a reference to the VMExecutionHandler used.
func (ae *ArwenTestExecutor) GetVM() vmi.VMExecutionHandler {
	return ae.vm
}

func (ae *ArwenTestExecutor) gasScheduleMapFromMandos(mandosGasSchedule mj.GasSchedule) (config.GasScheduleMap, error) {
	switch mandosGasSchedule {
	case mj.GasScheduleDefault:
		return gasSchedules.LoadGasScheduleConfig(gasSchedules.GetV3())
	case mj.GasScheduleDummy:
		return config.MakeGasMapForTests(), nil
	case mj.GasScheduleV1:
		return gasSchedules.LoadGasScheduleConfig(gasSchedules.GetV1())
	case mj.GasScheduleV2:
		return gasSchedules.LoadGasScheduleConfig(gasSchedules.GetV2())
	case mj.GasScheduleV3:
		return gasSchedules.LoadGasScheduleConfig(gasSchedules.GetV3())
	default:
		return nil, fmt.Errorf("unknown mandos GasSchedule: %d", mandosGasSchedule)
	}
}

// SetMandosGasSchedule updates the gas costs based on the mandos scenario config
// only changes the gas schedule once,
// this prevents subsequent gasSchedule declarations in externalSteps to overwrite
func (ae *ArwenTestExecutor) SetMandosGasSchedule(newGasSchedule mj.GasSchedule) error {
	if ae.mandosGasScheduleLoaded {
		return nil
	}
	gasSchedule, err := ae.gasScheduleMapFromMandos(newGasSchedule)
	if err != nil {
		return err
	}
	ae.mandosGasScheduleLoaded = true
	ae.vm.GasScheduleChange(gasSchedule)
	return nil
}
