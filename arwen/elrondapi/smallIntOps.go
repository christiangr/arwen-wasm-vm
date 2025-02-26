package elrondapi

// // Declare the function signatures (see [cgo](https://golang.org/cmd/cgo/)).
//
// #include <stdlib.h>
// typedef unsigned char uint8_t;
// typedef int int32_t;
//
// extern long long v1_4_smallIntGetUnsignedArgument(void *context, int32_t id);
// extern long long v1_4_smallIntGetSignedArgument(void *context, int32_t id);
//
// extern void			v1_4_smallIntFinishUnsigned(void* context, long long value);
// extern void			v1_4_smallIntFinishSigned(void* context, long long value);
//
// extern int32_t		v1_4_smallIntStorageStoreUnsigned(void *context, int32_t keyOffset, int32_t keyLength, long long value);
// extern int32_t		v1_4_smallIntStorageStoreSigned(void *context, int32_t keyOffset, int32_t keyLength, long long value);
// extern long long v1_4_smallIntStorageLoadUnsigned(void *context, int32_t keyOffset, int32_t keyLength);
// extern long long v1_4_smallIntStorageLoadSigned(void *context, int32_t keyOffset, int32_t keyLength);
//
// extern long long v1_4_int64getArgument(void *context, int32_t id);
// extern int32_t		v1_4_int64storageStore(void *context, int32_t keyOffset, int32_t keyLength , long long value);
// extern long long v1_4_int64storageLoad(void *context, int32_t keyOffset, int32_t keyLength );
// extern void			v1_4_int64finish(void* context, long long value);
//
import "C"

import (
	"math/big"
	"unsafe"

	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/arwen"
	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/wasmer"
	twos "github.com/ElrondNetwork/big-int-util/twos-complement"
)

// SmallIntImports creates a new wasmer.Imports populated with the small int (int64/uint64) API methods
func SmallIntImports(imports *wasmer.Imports) (*wasmer.Imports, error) {
	imports = imports.Namespace("env")

	imports, err := imports.Append("smallIntGetUnsignedArgument", v1_4_smallIntGetUnsignedArgument, C.v1_4_smallIntGetUnsignedArgument)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("smallIntGetSignedArgument", v1_4_smallIntGetSignedArgument, C.v1_4_smallIntGetSignedArgument)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("smallIntFinishUnsigned", v1_4_smallIntFinishUnsigned, C.v1_4_smallIntFinishUnsigned)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("smallIntFinishSigned", v1_4_smallIntFinishSigned, C.v1_4_smallIntFinishSigned)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("smallIntStorageStoreUnsigned", v1_4_smallIntStorageStoreUnsigned, C.v1_4_smallIntStorageStoreUnsigned)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("smallIntStorageStoreSigned", v1_4_smallIntStorageStoreSigned, C.v1_4_smallIntStorageStoreSigned)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("smallIntStorageLoadUnsigned", v1_4_smallIntStorageLoadUnsigned, C.v1_4_smallIntStorageLoadUnsigned)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("smallIntStorageLoadSigned", v1_4_smallIntStorageLoadSigned, C.v1_4_smallIntStorageLoadSigned)
	if err != nil {
		return nil, err
	}

	// the last are just for backwards compatibility:

	imports, err = imports.Append("int64getArgument", v1_4_int64getArgument, C.v1_4_int64getArgument)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("int64storageStore", v1_4_int64storageStore, C.v1_4_int64storageStore)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("int64storageLoad", v1_4_int64storageLoad, C.v1_4_int64storageLoad)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("int64finish", v1_4_int64finish, C.v1_4_int64finish)
	if err != nil {
		return nil, err
	}

	return imports, nil
}

//export v1_4_smallIntGetUnsignedArgument
func v1_4_smallIntGetUnsignedArgument(context unsafe.Pointer, id int32) int64 {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.Int64GetArgument
	metering.UseGas(gasToUse)

	args := runtime.Arguments()
	if id < 0 || id >= int32(len(args)) {
		arwen.WithFault(arwen.ErrArgIndexOutOfRange, context, runtime.ElrondAPIErrorShouldFailExecution())
		return 0
	}

	arg := args[id]
	argBigInt := big.NewInt(0).SetBytes(arg)
	if !argBigInt.IsUint64() {
		arwen.WithFault(arwen.ErrArgOutOfRange, context, runtime.ElrondAPIErrorShouldFailExecution())
		return 0
	}
	return int64(argBigInt.Uint64())
}

//export v1_4_smallIntGetSignedArgument
func v1_4_smallIntGetSignedArgument(context unsafe.Pointer, id int32) int64 {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.Int64GetArgument
	metering.UseGas(gasToUse)

	args := runtime.Arguments()
	if id < 0 || id >= int32(len(args)) {
		arwen.WithFault(arwen.ErrArgIndexOutOfRange, context, runtime.ElrondAPIErrorShouldFailExecution())
		return 0
	}

	arg := args[id]
	argBigInt := twos.SetBytes(big.NewInt(0), arg)
	if !argBigInt.IsInt64() {
		arwen.WithFault(arwen.ErrArgOutOfRange, context, runtime.ElrondAPIErrorShouldFailExecution())
		return 0
	}
	return argBigInt.Int64()
}

//export v1_4_smallIntFinishUnsigned
func v1_4_smallIntFinishUnsigned(context unsafe.Pointer, value int64) {
	output := arwen.GetOutputContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.Int64Finish
	metering.UseGas(gasToUse)

	valueBytes := big.NewInt(0).SetUint64(uint64(value)).Bytes()
	output.Finish(valueBytes)
}

//export v1_4_smallIntFinishSigned
func v1_4_smallIntFinishSigned(context unsafe.Pointer, value int64) {
	output := arwen.GetOutputContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.Int64Finish
	metering.UseGas(gasToUse)

	valueBytes := twos.ToBytes(big.NewInt(value))
	output.Finish(valueBytes)
}

//export v1_4_smallIntStorageStoreUnsigned
func v1_4_smallIntStorageStoreUnsigned(context unsafe.Pointer, keyOffset int32, keyLength int32, value int64) int32 {
	runtime := arwen.GetRuntimeContext(context)
	storage := arwen.GetStorageContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.Int64StorageStore
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	valueBytes := big.NewInt(0).SetUint64(uint64(value)).Bytes()
	storageStatus, err := storage.SetStorage(key, valueBytes)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	return int32(storageStatus)
}

//export v1_4_smallIntStorageStoreSigned
func v1_4_smallIntStorageStoreSigned(context unsafe.Pointer, keyOffset int32, keyLength int32, value int64) int32 {
	runtime := arwen.GetRuntimeContext(context)
	storage := arwen.GetStorageContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.Int64StorageStore
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	valueBytes := twos.ToBytes(big.NewInt(value))
	storageStatus, err := storage.SetStorage(key, valueBytes)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	return int32(storageStatus)
}

//export v1_4_smallIntStorageLoadUnsigned
func v1_4_smallIntStorageLoadUnsigned(context unsafe.Pointer, keyOffset int32, keyLength int32) int64 {
	runtime := arwen.GetRuntimeContext(context)
	storage := arwen.GetStorageContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.Int64StorageLoad
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 0
	}

	data := storage.GetStorage(key)
	valueBigInt := big.NewInt(0).SetBytes(data)
	if !valueBigInt.IsUint64() {
		arwen.WithFault(arwen.ErrStorageValueOutOfRange, context, runtime.ElrondAPIErrorShouldFailExecution())
		return 0
	}

	return int64(valueBigInt.Uint64())
}

//export v1_4_smallIntStorageLoadSigned
func v1_4_smallIntStorageLoadSigned(context unsafe.Pointer, keyOffset int32, keyLength int32) int64 {
	runtime := arwen.GetRuntimeContext(context)
	storage := arwen.GetStorageContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.Int64StorageLoad
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 0
	}

	data := storage.GetStorage(key)
	valueBigInt := twos.SetBytes(big.NewInt(0), data)
	if !valueBigInt.IsInt64() {
		arwen.WithFault(arwen.ErrStorageValueOutOfRange, context, runtime.ElrondAPIErrorShouldFailExecution())
		return 0
	}

	return valueBigInt.Int64()
}

//export v1_4_int64getArgument
func v1_4_int64getArgument(context unsafe.Pointer, id int32) int64 {
	// backwards compatibility
	return v1_4_smallIntGetSignedArgument(context, id)
}

//export v1_4_int64finish
func v1_4_int64finish(context unsafe.Pointer, value int64) {
	// backwards compatibility
	v1_4_smallIntFinishSigned(context, value)
}

//export v1_4_int64storageStore
func v1_4_int64storageStore(context unsafe.Pointer, keyOffset int32, keyLength int32, value int64) int32 {
	// backwards compatibility
	return v1_4_smallIntStorageStoreUnsigned(context, keyOffset, keyLength, value)
}

//export v1_4_int64storageLoad
func v1_4_int64storageLoad(context unsafe.Pointer, keyOffset int32, keyLength int32) int64 {
	// backwards compatibility
	return v1_4_smallIntStorageLoadUnsigned(context, keyOffset, keyLength)
}
