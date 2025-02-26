package elrondapi

// // Declare the function signatures (see [cgo](https://golang.org/cmd/cgo/)).
//
// #include <stdlib.h>
// typedef unsigned char uint8_t;
// typedef int int32_t;
//
// extern int32_t		v1_4_bigIntNew(void* context, long long smallValue);
//
// extern int32_t		v1_4_bigIntUnsignedByteLength(void* context, int32_t reference);
// extern int32_t		v1_4_bigIntSignedByteLength(void* context, int32_t reference);
// extern int32_t		v1_4_bigIntGetUnsignedBytes(void* context, int32_t reference, int32_t byteOffset);
// extern int32_t		v1_4_bigIntGetSignedBytes(void* context, int32_t reference, int32_t byteOffset);
// extern void			v1_4_bigIntSetUnsignedBytes(void* context, int32_t destination, int32_t byteOffset, int32_t byteLength);
// extern void			v1_4_bigIntSetSignedBytes(void* context, int32_t destination, int32_t byteOffset, int32_t byteLength);
//
// extern int32_t		v1_4_bigIntIsInt64(void* context, int32_t reference);
// extern long long		v1_4_bigIntGetInt64(void* context, int32_t reference);
// extern void			v1_4_bigIntSetInt64(void* context, int32_t destination, long long value);
// extern long long		v1_4_bigIntGetOrCreateInt64(void* context, int32_t reference);
//
// extern void			v1_4_bigIntAdd(void* context, int32_t destination, int32_t op1, int32_t op2);
// extern void			v1_4_bigIntSub(void* context, int32_t destination, int32_t op1, int32_t op2);
// extern void			v1_4_bigIntMul(void* context, int32_t destination, int32_t op1, int32_t op2);
// extern void			v1_4_bigIntTDiv(void* context, int32_t destination, int32_t op1, int32_t op2);
// extern void			v1_4_bigIntTMod(void* context, int32_t destination, int32_t op1, int32_t op2);
// extern void			v1_4_bigIntEDiv(void* context, int32_t destination, int32_t op1, int32_t op2);
// extern void			v1_4_bigIntEMod(void* context, int32_t destination, int32_t op1, int32_t op2);
//
// extern void			v1_4_bigIntPow(void* context, int32_t destination, int32_t op1, int32_t op2);
// extern int32_t		v1_4_bigIntLog2(void* context, int32_t op);
// extern void			v1_4_bigIntSqrt(void* context, int32_t destination, int32_t op);
//
// extern void			v1_4_bigIntAbs(void* context, int32_t destination, int32_t op);
// extern void			v1_4_bigIntNeg(void* context, int32_t destination, int32_t op);
// extern int32_t		v1_4_bigIntSign(void* context, int32_t op);
// extern int32_t		v1_4_bigIntCmp(void* context, int32_t op1, int32_t op2);
//
// extern void			v1_4_bigIntNot(void* context, int32_t destination, int32_t op);
// extern void			v1_4_bigIntAnd(void* context, int32_t destination, int32_t op1, int32_t op2);
// extern void			v1_4_bigIntOr(void* context, int32_t destination, int32_t op1, int32_t op2);
// extern void			v1_4_bigIntXor(void* context, int32_t destination, int32_t op1, int32_t op2);
// extern void			v1_4_bigIntShr(void* context, int32_t destination, int32_t op, int32_t bits);
// extern void			v1_4_bigIntShl(void* context, int32_t destination, int32_t op, int32_t bits);
//
// extern void			v1_4_bigIntFinishUnsigned(void* context, int32_t reference);
// extern void			v1_4_bigIntFinishSigned(void* context, int32_t reference);
// extern int32_t		v1_4_bigIntStorageStoreUnsigned(void *context, int32_t keyOffset, int32_t keyLength, int32_t source);
// extern int32_t		v1_4_bigIntStorageLoadUnsigned(void *context, int32_t keyOffset, int32_t keyLength, int32_t destination);
// extern void			v1_4_bigIntGetUnsignedArgument(void *context, int32_t id, int32_t destination);
// extern void			v1_4_bigIntGetSignedArgument(void *context, int32_t id, int32_t destination);
// extern void			v1_4_bigIntGetCallValue(void *context, int32_t destination);
// extern void			v1_4_bigIntGetESDTCallValue(void *context, int32_t destination);
// extern void			v1_4_bigIntGetESDTCallValueByIndex(void *context, int32_t destination, int32_t index);
// extern void			v1_4_bigIntGetESDTExternalBalance(void *context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen, long long nonce, int32_t result);
// extern void			v1_4_bigIntGetExternalBalance(void *context, int32_t addressOffset, int32_t result);
import "C"

import (
	"math/big"
	"unsafe"

	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/arwen"
	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/math"
	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/wasmer"
	twos "github.com/ElrondNetwork/big-int-util/twos-complement"
)

// BigIntImports creates a new wasmer.Imports populated with the BigInt API methods
func BigIntImports(imports *wasmer.Imports) (*wasmer.Imports, error) {
	imports = imports.Namespace("env")

	imports, err := imports.Append("bigIntNew", v1_4_bigIntNew, C.v1_4_bigIntNew)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntUnsignedByteLength", v1_4_bigIntUnsignedByteLength, C.v1_4_bigIntUnsignedByteLength)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntSignedByteLength", v1_4_bigIntSignedByteLength, C.v1_4_bigIntSignedByteLength)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntGetUnsignedBytes", v1_4_bigIntGetUnsignedBytes, C.v1_4_bigIntGetUnsignedBytes)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntGetSignedBytes", v1_4_bigIntGetSignedBytes, C.v1_4_bigIntGetSignedBytes)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntSetUnsignedBytes", v1_4_bigIntSetUnsignedBytes, C.v1_4_bigIntSetUnsignedBytes)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntSetSignedBytes", v1_4_bigIntSetSignedBytes, C.v1_4_bigIntSetSignedBytes)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntIsInt64", v1_4_bigIntIsInt64, C.v1_4_bigIntIsInt64)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntGetInt64", v1_4_bigIntGetInt64, C.v1_4_bigIntGetInt64)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntGetOrCreateInt64", v1_4_bigIntGetOrCreateInt64, C.v1_4_bigIntGetOrCreateInt64)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntSetInt64", v1_4_bigIntSetInt64, C.v1_4_bigIntSetInt64)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntAdd", v1_4_bigIntAdd, C.v1_4_bigIntAdd)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntSub", v1_4_bigIntSub, C.v1_4_bigIntSub)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntMul", v1_4_bigIntMul, C.v1_4_bigIntMul)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntTDiv", v1_4_bigIntTDiv, C.v1_4_bigIntTDiv)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntTMod", v1_4_bigIntTMod, C.v1_4_bigIntTMod)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntEDiv", v1_4_bigIntEDiv, C.v1_4_bigIntEDiv)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntEMod", v1_4_bigIntEMod, C.v1_4_bigIntEMod)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntSqrt", v1_4_bigIntSqrt, C.v1_4_bigIntSqrt)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntPow", v1_4_bigIntPow, C.v1_4_bigIntPow)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntLog2", v1_4_bigIntLog2, C.v1_4_bigIntLog2)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntAbs", v1_4_bigIntAbs, C.v1_4_bigIntAbs)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntNeg", v1_4_bigIntNeg, C.v1_4_bigIntNeg)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntSign", v1_4_bigIntSign, C.v1_4_bigIntSign)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntCmp", v1_4_bigIntCmp, C.v1_4_bigIntCmp)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntNot", v1_4_bigIntNot, C.v1_4_bigIntNot)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntAnd", v1_4_bigIntAnd, C.v1_4_bigIntAnd)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntOr", v1_4_bigIntOr, C.v1_4_bigIntOr)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntXor", v1_4_bigIntXor, C.v1_4_bigIntXor)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntShr", v1_4_bigIntShr, C.v1_4_bigIntShr)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntShl", v1_4_bigIntShl, C.v1_4_bigIntShl)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntFinishUnsigned", v1_4_bigIntFinishUnsigned, C.v1_4_bigIntFinishUnsigned)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntFinishSigned", v1_4_bigIntFinishSigned, C.v1_4_bigIntFinishSigned)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntStorageStoreUnsigned", v1_4_bigIntStorageStoreUnsigned, C.v1_4_bigIntStorageStoreUnsigned)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntStorageLoadUnsigned", v1_4_bigIntStorageLoadUnsigned, C.v1_4_bigIntStorageLoadUnsigned)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntGetUnsignedArgument", v1_4_bigIntGetUnsignedArgument, C.v1_4_bigIntGetUnsignedArgument)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntGetSignedArgument", v1_4_bigIntGetSignedArgument, C.v1_4_bigIntGetSignedArgument)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntGetCallValue", v1_4_bigIntGetCallValue, C.v1_4_bigIntGetCallValue)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntGetESDTCallValue", v1_4_bigIntGetESDTCallValue, C.v1_4_bigIntGetESDTCallValue)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntGetESDTExternalBalance", v1_4_bigIntGetESDTExternalBalance, C.v1_4_bigIntGetESDTExternalBalance)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntGetExternalBalance", v1_4_bigIntGetExternalBalance, C.v1_4_bigIntGetExternalBalance)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("bigIntGetESDTCallValueByIndex", v1_4_bigIntGetESDTCallValueByIndex, C.v1_4_bigIntGetESDTCallValueByIndex)
	if err != nil {
		return nil, err
	}

	return imports, nil
}

//export v1_4_bigIntGetUnsignedArgument
func v1_4_bigIntGetUnsignedArgument(context unsafe.Pointer, id int32, destinationHandle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntGetUnsignedArgument
	metering.UseGas(gasToUse)

	args := runtime.Arguments()
	if int32(len(args)) <= id {
		return
	}

	value := managedType.GetBigIntOrCreate(destinationHandle)

	value.SetBytes(args[id])
}

//export v1_4_bigIntGetSignedArgument
func v1_4_bigIntGetSignedArgument(context unsafe.Pointer, id int32, destinationHandle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntGetSignedArgument
	metering.UseGas(gasToUse)

	args := runtime.Arguments()
	if int32(len(args)) <= id {
		return
	}

	value := managedType.GetBigIntOrCreate(destinationHandle)

	twos.SetBytes(value, args[id])
}

//export v1_4_bigIntStorageStoreUnsigned
func v1_4_bigIntStorageStoreUnsigned(context unsafe.Pointer, keyOffset int32, keyLength int32, sourceHandle int32) int32 {
	managedType := arwen.GetManagedTypesContext(context)
	runtime := arwen.GetRuntimeContext(context)
	storage := arwen.GetStorageContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntStorageStoreUnsigned
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if arwen.WithFault(err, context, runtime.BigIntAPIErrorShouldFailExecution()) {
		return -1
	}

	value := managedType.GetBigIntOrCreate(sourceHandle)
	bytes := value.Bytes()

	storageStatus, err := storage.SetStorage(key, bytes)
	if arwen.WithFault(err, context, runtime.BigIntAPIErrorShouldFailExecution()) {
		return -1
	}

	return int32(storageStatus)
}

//export v1_4_bigIntStorageLoadUnsigned
func v1_4_bigIntStorageLoadUnsigned(context unsafe.Pointer, keyOffset int32, keyLength int32, destinationHandle int32) int32 {
	managedType := arwen.GetManagedTypesContext(context)
	runtime := arwen.GetRuntimeContext(context)
	storage := arwen.GetStorageContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntStorageLoadUnsigned
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if arwen.WithFault(err, context, runtime.BigIntAPIErrorShouldFailExecution()) {
		return -1
	}

	bytes := storage.GetStorage(key)

	value := managedType.GetBigIntOrCreate(destinationHandle)
	value.SetBytes(bytes)

	return int32(len(bytes))
}

//export v1_4_bigIntGetCallValue
func v1_4_bigIntGetCallValue(context unsafe.Pointer, destinationHandle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntGetCallValue
	metering.UseGas(gasToUse)

	value := managedType.GetBigIntOrCreate(destinationHandle)
	value.Set(runtime.GetVMInput().CallValue)
}

//export v1_4_bigIntGetESDTCallValue
func v1_4_bigIntGetESDTCallValue(context unsafe.Pointer, destination int32) {
	_ = failIfMoreThanOneESDTTransfer(context)
	v1_4_bigIntGetESDTCallValueByIndex(context, destination, 0)
}

//export v1_4_bigIntGetESDTCallValueByIndex
func v1_4_bigIntGetESDTCallValueByIndex(context unsafe.Pointer, destinationHandle int32, index int32) {
	managedType := arwen.GetManagedTypesContext(context)
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntGetCallValue
	metering.UseGas(gasToUse)

	value := managedType.GetBigIntOrCreate(destinationHandle)
	esdtTransfer := getESDTTransferFromInput(runtime.GetVMInput(), index)
	if esdtTransfer != nil {
		value.Set(esdtTransfer.ESDTValue)
	} else {
		value.Set(big.NewInt(0))
	}
}

//export v1_4_bigIntGetExternalBalance
func v1_4_bigIntGetExternalBalance(context unsafe.Pointer, addressOffset int32, result int32) {
	managedType := arwen.GetManagedTypesContext(context)
	runtime := arwen.GetRuntimeContext(context)
	blockchain := arwen.GetBlockchainContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntGetExternalBalance
	metering.UseGas(gasToUse)

	address, err := runtime.MemLoad(addressOffset, arwen.AddressLen)
	if arwen.WithFault(err, context, runtime.BigIntAPIErrorShouldFailExecution()) {
		return
	}

	balance := blockchain.GetBalance(address)
	value := managedType.GetBigIntOrCreate(result)

	value.SetBytes(balance)
}

//export v1_4_bigIntGetESDTExternalBalance
func v1_4_bigIntGetESDTExternalBalance(context unsafe.Pointer, addressOffset int32, tokenIDOffset int32, tokenIDLen int32, nonce int64, resultHandle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntGetExternalBalance
	metering.UseGas(gasToUse)

	esdtData, err := getESDTDataFromBlockchainHook(context, addressOffset, tokenIDOffset, tokenIDLen, nonce)
	if arwen.WithFault(err, context, runtime.BigIntAPIErrorShouldFailExecution()) {
		return
	}
	if esdtData == nil {
		return
	}

	value := managedType.GetBigIntOrCreate(resultHandle)
	value.Set(esdtData.Value)
}

//export v1_4_bigIntNew
func v1_4_bigIntNew(context unsafe.Pointer, smallValue int64) int32 {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntNew
	metering.UseGas(gasToUse)

	return managedType.PutBigInt(smallValue)
}

//export v1_4_bigIntUnsignedByteLength
func v1_4_bigIntUnsignedByteLength(context unsafe.Pointer, referenceHandle int32) int32 {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntUnsignedByteLength
	metering.UseGas(gasToUse)

	value, err := managedType.GetBigInt(referenceHandle)
	if arwen.WithFault(err, context, runtime.BigIntAPIErrorShouldFailExecution()) {
		return -1
	}

	bytes := value.Bytes()
	return int32(len(bytes))
}

//export v1_4_bigIntSignedByteLength
func v1_4_bigIntSignedByteLength(context unsafe.Pointer, referenceHandle int32) int32 {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSignedByteLength
	metering.UseGas(gasToUse)

	value, err := managedType.GetBigInt(referenceHandle)
	if arwen.WithFault(err, context, runtime.BigIntAPIErrorShouldFailExecution()) {
		return 0
	}

	bytes := twos.ToBytes(value)
	return int32(len(bytes))
}

//export v1_4_bigIntGetUnsignedBytes
func v1_4_bigIntGetUnsignedBytes(context unsafe.Pointer, referenceHandle int32, byteOffset int32) int32 {
	managedType := arwen.GetManagedTypesContext(context)
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntGetUnsignedBytes
	metering.UseGas(gasToUse)

	value, err := managedType.GetBigInt(referenceHandle)
	if arwen.WithFault(err, context, runtime.BigIntAPIErrorShouldFailExecution()) {
		return -1
	}
	bytes := value.Bytes()

	err = runtime.MemStore(byteOffset, bytes)
	if arwen.WithFault(err, context, runtime.BigIntAPIErrorShouldFailExecution()) {
		return -1
	}

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(len(bytes)))
	metering.UseGas(gasToUse)

	return int32(len(bytes))
}

//export v1_4_bigIntGetSignedBytes
func v1_4_bigIntGetSignedBytes(context unsafe.Pointer, referenceHandle int32, byteOffset int32) int32 {
	managedType := arwen.GetManagedTypesContext(context)
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntGetSignedBytes
	metering.UseGas(gasToUse)

	value, err := managedType.GetBigInt(referenceHandle)
	if arwen.WithFault(err, context, runtime.BigIntAPIErrorShouldFailExecution()) {
		return -1
	}
	bytes := twos.ToBytes(value)

	err = runtime.MemStore(byteOffset, bytes)
	if arwen.WithFault(err, context, runtime.BigIntAPIErrorShouldFailExecution()) {
		return -1
	}

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(len(bytes)))
	metering.UseGas(gasToUse)

	return int32(len(bytes))
}

//export v1_4_bigIntSetUnsignedBytes
func v1_4_bigIntSetUnsignedBytes(context unsafe.Pointer, destinationHandle int32, byteOffset int32, byteLength int32) {
	managedType := arwen.GetManagedTypesContext(context)
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSetUnsignedBytes
	metering.UseGas(gasToUse)

	bytes, err := runtime.MemLoad(byteOffset, byteLength)
	if arwen.WithFault(err, context, runtime.BigIntAPIErrorShouldFailExecution()) {
		return
	}

	value := managedType.GetBigIntOrCreate(destinationHandle)
	value.SetBytes(bytes)

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(len(bytes)))
	metering.UseGas(gasToUse)
}

//export v1_4_bigIntSetSignedBytes
func v1_4_bigIntSetSignedBytes(context unsafe.Pointer, destinationHandle int32, byteOffset int32, byteLength int32) {
	managedType := arwen.GetManagedTypesContext(context)
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSetSignedBytes
	metering.UseGas(gasToUse)

	bytes, err := runtime.MemLoad(byteOffset, byteLength)
	if arwen.WithFault(err, context, runtime.BigIntAPIErrorShouldFailExecution()) {
		return
	}

	value := managedType.GetBigIntOrCreate(destinationHandle)
	twos.SetBytes(value, bytes)

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(len(bytes)))
	metering.UseGas(gasToUse)
}

//export v1_4_bigIntIsInt64
func v1_4_bigIntIsInt64(context unsafe.Pointer, destinationHandle int32) int32 {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntIsInt64
	metering.UseGas(gasToUse)

	value, err := managedType.GetBigInt(destinationHandle)
	if arwen.WithFault(err, context, runtime.BigIntAPIErrorShouldFailExecution()) {
		return -1
	}
	if value.IsInt64() {
		return 1
	}
	return 0
}

//export v1_4_bigIntGetInt64
func v1_4_bigIntGetInt64(context unsafe.Pointer, destinationHandle int32) int64 {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntGetInt64
	metering.UseGas(gasToUse)

	value, err := managedType.GetBigInt(destinationHandle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
	}
	return value.Int64()
}

//export v1_4_bigIntGetOrCreateInt64
func v1_4_bigIntGetOrCreateInt64(context unsafe.Pointer, destinationHandle int32) int64 {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntGetInt64
	metering.UseGas(gasToUse)

	value := managedType.GetBigIntOrCreate(destinationHandle)
	return value.Int64()
}

//export v1_4_bigIntSetInt64
func v1_4_bigIntSetInt64(context unsafe.Pointer, destinationHandle int32, value int64) {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSub
	metering.UseGas(gasToUse)

	dest := managedType.GetBigIntOrCreate(destinationHandle)
	dest.SetInt64(value)
}

//export v1_4_bigIntAdd
func v1_4_bigIntAdd(context unsafe.Pointer, destinationHandle, op1Handle, op2Handle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext((context))

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSub
	metering.UseGas(gasToUse)

	dest := managedType.GetBigIntOrCreate(destinationHandle)
	a, b, err := managedType.GetTwoBigInt(op1Handle, op2Handle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	managedType.ConsumeGasForBigIntCopy(dest, a, b)
	dest.Add(a, b)
}

//export v1_4_bigIntSub
func v1_4_bigIntSub(context unsafe.Pointer, destinationHandle, op1Handle, op2Handle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSub
	metering.UseGas(gasToUse)

	dest := managedType.GetBigIntOrCreate(destinationHandle)
	a, b, err := managedType.GetTwoBigInt(op1Handle, op2Handle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	managedType.ConsumeGasForBigIntCopy(dest, a, b)
	dest.Sub(a, b)
}

//export v1_4_bigIntMul
func v1_4_bigIntMul(context unsafe.Pointer, destinationHandle, op1Handle, op2Handle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntMul
	metering.UseGas(gasToUse)

	dest := managedType.GetBigIntOrCreate(destinationHandle)
	a, b, err := managedType.GetTwoBigInt(op1Handle, op2Handle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	managedType.ConsumeGasForBigIntCopy(dest, a, b)
	dest.Mul(a, b)
}

//export v1_4_bigIntTDiv
func v1_4_bigIntTDiv(context unsafe.Pointer, destinationHandle, op1Handle, op2Handle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntTDiv
	metering.UseGas(gasToUse)

	dest := managedType.GetBigIntOrCreate(destinationHandle)
	a, b, err := managedType.GetTwoBigInt(op1Handle, op2Handle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	managedType.ConsumeGasForBigIntCopy(dest, a, b)
	if b.Sign() == 0 {
		runtime := arwen.GetRuntimeContext(context)
		arwen.WithFault(arwen.ErrDivZero, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	dest.Quo(a, b) // Quo implements truncated division (like Go)
}

//export v1_4_bigIntTMod
func v1_4_bigIntTMod(context unsafe.Pointer, destinationHandle, op1Handle, op2Handle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSub
	metering.UseGas(gasToUse)

	dest := managedType.GetBigIntOrCreate(destinationHandle)
	a, b, err := managedType.GetTwoBigInt(op1Handle, op2Handle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	managedType.ConsumeGasForBigIntCopy(dest, a, b)
	if b.Sign() == 0 {
		runtime := arwen.GetRuntimeContext(context)
		arwen.WithFault(arwen.ErrDivZero, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	dest.Rem(a, b) // Rem implements truncated modulus (like Go)
}

//export v1_4_bigIntEDiv
func v1_4_bigIntEDiv(context unsafe.Pointer, destinationHandle, op1Handle, op2Handle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSub
	metering.UseGas(gasToUse)

	dest := managedType.GetBigIntOrCreate(destinationHandle)
	a, b, err := managedType.GetTwoBigInt(op1Handle, op2Handle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	managedType.ConsumeGasForBigIntCopy(dest, a, b)
	if b.Sign() == 0 {
		runtime := arwen.GetRuntimeContext(context)
		arwen.WithFault(arwen.ErrDivZero, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	dest.Div(a, b) // Div implements Euclidean division (unlike Go)
}

//export v1_4_bigIntEMod
func v1_4_bigIntEMod(context unsafe.Pointer, destinationHandle, op1Handle, op2Handle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSub
	metering.UseGas(gasToUse)

	dest := managedType.GetBigIntOrCreate(destinationHandle)
	a, b, err := managedType.GetTwoBigInt(op1Handle, op2Handle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	managedType.ConsumeGasForBigIntCopy(dest, a, b)
	if b.Sign() == 0 {
		runtime := arwen.GetRuntimeContext(context)
		arwen.WithFault(arwen.ErrDivZero, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	dest.Mod(a, b) // Mod implements Euclidean division (unlike Go)
}

//export v1_4_bigIntSqrt
func v1_4_bigIntSqrt(context unsafe.Pointer, destinationHandle, opHandle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSqrt
	metering.UseGas(gasToUse)

	dest := managedType.GetBigIntOrCreate(destinationHandle)
	a, err := managedType.GetBigInt(opHandle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	managedType.ConsumeGasForBigIntCopy(dest, a)
	if a.Sign() < 0 {
		runtime := arwen.GetRuntimeContext(context)
		arwen.WithFault(arwen.ErrBadLowerBounds, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	dest.Sqrt(a)
}

//export v1_4_bigIntPow
func v1_4_bigIntPow(context unsafe.Pointer, destinationHandle, op1Handle, op2Handle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntPow
	metering.UseGas(gasToUse)

	dest := managedType.GetBigIntOrCreate(destinationHandle)
	a, b, err := managedType.GetTwoBigInt(op1Handle, op2Handle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}

	//this calculates the length of the result in bytes
	lengthOfResult := big.NewInt(0).Div(big.NewInt(0).Mul(b, big.NewInt(int64(a.BitLen()))), big.NewInt(8))

	managedType.ConsumeGasForThisBigIntNumberOfBytes(lengthOfResult)
	managedType.ConsumeGasForBigIntCopy(a, b)

	if b.Sign() < 0 {
		runtime := arwen.GetRuntimeContext(context)
		arwen.WithFault(arwen.ErrBadLowerBounds, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}

	dest.Exp(a, b, nil)
}

//export v1_4_bigIntLog2
func v1_4_bigIntLog2(context unsafe.Pointer, op1Handle int32) int32 {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntLog
	metering.UseGas(gasToUse)

	a, err := managedType.GetBigInt(op1Handle)
	if arwen.WithFault(err, context, runtime.BigIntAPIErrorShouldFailExecution()) {
		return -1
	}
	managedType.ConsumeGasForBigIntCopy(a)
	if a.Sign() < 0 {
		runtime := arwen.GetRuntimeContext(context)
		arwen.WithFault(arwen.ErrBadLowerBounds, context, runtime.BigIntAPIErrorShouldFailExecution())
		return -1
	}

	return int32(a.BitLen() - 1)
}

//export v1_4_bigIntAbs
func v1_4_bigIntAbs(context unsafe.Pointer, destinationHandle, opHandle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSub
	metering.UseGas(gasToUse)

	dest := managedType.GetBigIntOrCreate(destinationHandle)
	a, err := managedType.GetBigInt(opHandle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	managedType.ConsumeGasForBigIntCopy(dest, a)
	dest.Abs(a)
}

//export v1_4_bigIntNeg
func v1_4_bigIntNeg(context unsafe.Pointer, destinationHandle, opHandle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSub
	metering.UseGas(gasToUse)

	dest := managedType.GetBigIntOrCreate(destinationHandle)
	a, err := managedType.GetBigInt(opHandle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	managedType.ConsumeGasForBigIntCopy(dest, a)
	dest.Neg(a)
}

//export v1_4_bigIntSign
func v1_4_bigIntSign(context unsafe.Pointer, opHandle int32) int32 {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSign
	metering.UseGas(gasToUse)

	a, err := managedType.GetBigInt(opHandle)
	if arwen.WithFault(err, context, runtime.BigIntAPIErrorShouldFailExecution()) {
		return -2
	}
	managedType.ConsumeGasForBigIntCopy(a)
	return int32(a.Sign())
}

//export v1_4_bigIntCmp
func v1_4_bigIntCmp(context unsafe.Pointer, op1Handle, op2Handle int32) int32 {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntCmp
	metering.UseGas(gasToUse)

	a, b, err := managedType.GetTwoBigInt(op1Handle, op2Handle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
		return -2
	}
	managedType.ConsumeGasForBigIntCopy(a, b)
	return int32(a.Cmp(b))
}

//export v1_4_bigIntNot
func v1_4_bigIntNot(context unsafe.Pointer, destinationHandle, opHandle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSub
	metering.UseGas(gasToUse)

	dest := managedType.GetBigIntOrCreate(destinationHandle)
	a, err := managedType.GetBigInt(opHandle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	managedType.ConsumeGasForBigIntCopy(dest, a)
	if a.Sign() < 0 {
		runtime := arwen.GetRuntimeContext(context)
		arwen.WithFault(arwen.ErrBitwiseNegative, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	dest.Not(a)
}

//export v1_4_bigIntAnd
func v1_4_bigIntAnd(context unsafe.Pointer, destinationHandle, op1Handle, op2Handle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSub
	metering.UseGas(gasToUse)

	dest := managedType.GetBigIntOrCreate(destinationHandle)
	a, b, err := managedType.GetTwoBigInt(op1Handle, op2Handle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	managedType.ConsumeGasForBigIntCopy(a, b)
	if a.Sign() < 0 || b.Sign() < 0 {
		runtime := arwen.GetRuntimeContext(context)
		arwen.WithFault(arwen.ErrBitwiseNegative, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	dest.And(a, b)
}

//export v1_4_bigIntOr
func v1_4_bigIntOr(context unsafe.Pointer, destinationHandle, op1Handle, op2Handle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSub
	metering.UseGas(gasToUse)

	dest := managedType.GetBigIntOrCreate(destinationHandle)
	a, b, err := managedType.GetTwoBigInt(op1Handle, op2Handle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	managedType.ConsumeGasForBigIntCopy(a, b)
	if a.Sign() < 0 || b.Sign() < 0 {
		runtime := arwen.GetRuntimeContext(context)
		arwen.WithFault(arwen.ErrBitwiseNegative, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	dest.Or(a, b)
}

//export v1_4_bigIntXor
func v1_4_bigIntXor(context unsafe.Pointer, destinationHandle, op1Handle, op2Handle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSub
	metering.UseGas(gasToUse)

	dest := managedType.GetBigIntOrCreate(destinationHandle)
	a, b, err := managedType.GetTwoBigInt(op1Handle, op2Handle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	managedType.ConsumeGasForBigIntCopy(a, b)
	if a.Sign() < 0 || b.Sign() < 0 {
		runtime := arwen.GetRuntimeContext(context)
		arwen.WithFault(arwen.ErrBitwiseNegative, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	dest.Xor(a, b)
}

//export v1_4_bigIntShr
func v1_4_bigIntShr(context unsafe.Pointer, destinationHandle, opHandle, bits int32) {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSub
	metering.UseGas(gasToUse)

	dest := managedType.GetBigIntOrCreate(destinationHandle)
	a, err := managedType.GetBigInt(opHandle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	managedType.ConsumeGasForBigIntCopy(a)
	if a.Sign() < 0 || bits < 0 {
		runtime := arwen.GetRuntimeContext(context)
		arwen.WithFault(arwen.ErrShiftNegative, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	dest.Rsh(a, uint(bits))
	managedType.ConsumeGasForBigIntCopy(dest)
}

//export v1_4_bigIntShl
func v1_4_bigIntShl(context unsafe.Pointer, destinationHandle, opHandle, bits int32) {
	managedType := arwen.GetManagedTypesContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntSub
	metering.UseGas(gasToUse)

	dest := managedType.GetBigIntOrCreate(destinationHandle)
	a, err := managedType.GetBigInt(opHandle)
	if err != nil {
		arwen.WithFault(arwen.ErrNoBigIntUnderThisHandle, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	managedType.ConsumeGasForBigIntCopy(a)
	if a.Sign() < 0 || bits < 0 {
		runtime := arwen.GetRuntimeContext(context)
		arwen.WithFault(arwen.ErrShiftNegative, context, runtime.BigIntAPIErrorShouldFailExecution())
		return
	}
	dest.Lsh(a, uint(bits))
	managedType.ConsumeGasForBigIntCopy(dest)
}

//export v1_4_bigIntFinishUnsigned
func v1_4_bigIntFinishUnsigned(context unsafe.Pointer, referenceHandle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	output := arwen.GetOutputContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntFinishUnsigned
	metering.UseGas(gasToUse)

	value, err := managedType.GetBigInt(referenceHandle)
	if arwen.WithFault(err, context, runtime.BigIntAPIErrorShouldFailExecution()) {
		return
	}
	bigIntBytes := value.Bytes()
	output.Finish(bigIntBytes)

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.PersistPerByte, uint64(len(value.Bytes())))
	metering.UseGas(gasToUse)
}

//export v1_4_bigIntFinishSigned
func v1_4_bigIntFinishSigned(context unsafe.Pointer, referenceHandle int32) {
	managedType := arwen.GetManagedTypesContext(context)
	output := arwen.GetOutputContext(context)
	metering := arwen.GetMeteringContext(context)
	runtime := arwen.GetRuntimeContext(context)

	gasToUse := metering.GasSchedule().BigIntAPICost.BigIntFinishSigned
	metering.UseGas(gasToUse)

	value, err := managedType.GetBigInt(referenceHandle)
	if arwen.WithFault(err, context, runtime.BigIntAPIErrorShouldFailExecution()) {
		return
	}
	bigInt2cBytes := twos.ToBytes(value)
	output.Finish(bigInt2cBytes)

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.PersistPerByte, uint64(len(bigInt2cBytes)))
	metering.UseGas(gasToUse)
}
