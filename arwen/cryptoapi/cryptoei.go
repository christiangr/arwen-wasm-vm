package cryptoapi

// // Declare the function signatures (see [cgo](https://golang.org/cmd/cgo/)).
//
// #include <stdlib.h>
// typedef unsigned char uint8_t;
// typedef int int32_t;
//
// extern int32_t sha256(void* context, int32_t dataOffset, int32_t length, int32_t resultOffset);
// extern int32_t keccak256(void *context, int32_t dataOffset, int32_t length, int32_t resultOffset);
// extern int32_t ripemd160(void *context, int32_t dataOffset, int32_t length, int32_t resultOffset);
// extern int32_t verifyBLS(void *context, int32_t keyOffset, int32_t messageOffset, int32_t messageLength, int32_t sigOffset);
// extern int32_t verifyEd25519(void *context, int32_t keyOffset, int32_t messageOffset, int32_t messageLength, int32_t sigOffset);
// extern int32_t verifySecp256k1(void *context, int32_t keyOffset, int32_t keyLength, int32_t messageOffset, int32_t messageLength, int32_t sigOffset);
// extern void addEC( void *context, int32_t destination1, int32_t destination2, int32_t fieldOrder, int32_t basePointOrder, int32_t eqConstant, int32_t xBasePoint, int32_t yBasePoint, int32_t sizeOfField, int32_t fstPointX, int32_t fstPointY, int32_t sndPointX, int32_t sndPointY);
// extern void doubleEC( void *context, int32_t destination1, int32_t destination2, int32_t fieldOrder, int32_t basePointOrder, int32_t eqConstant, int32_t xBasePoint, int32_t yBasePoint, int32_t sizeOfField, int32_t pointX, int32_t pointY);
// extern int32_t isOnCurveEC( void *context, int32_t fieldOrder, int32_t basePointOrder, int32_t eqConstant, int32_t xBasePoint, int32_t yBasePoint, int32_t sizeOfField, int32_t pointX, int32_t pointY);
// extern void scalarBaseMultEC ( void *context, int32_t destination1, int32_t destination2, int32_t fieldOrder, int32_t basePointOrder, int32_t eqConstant, int32_t xBasePoint, int32_t yBasePoint, int32_t sizeOfField, int32_t kOffset, int32_t length);
// extern void scalarMultEC ( void *context, int32_t destination1, int32_t destination2, int32_t fieldOrder, int32_t basePointOrder, int32_t eqConstant, int32_t xBasePoint, int32_t yBasePoint, int32_t sizeOfField, int32_t pointX, int32_t pointY, int32_t kOffset, int32_t length);
import "C"

import (
	"math/big"
	"unsafe"

	"github.com/ElrondNetwork/arwen-wasm-vm/arwen"
	elliptic_curve "github.com/ElrondNetwork/arwen-wasm-vm/crypto/elliptic_curves"
	"github.com/ElrondNetwork/arwen-wasm-vm/math"
	"github.com/ElrondNetwork/arwen-wasm-vm/wasmer"
)

const blsPublicKeyLength = 96
const blsSignatureLength = 48
const ed25519PublicKeyLength = 32
const ed25519SignatureLength = 64
const secp256k1CompressedPublicKeyLength = 33
const secp256k1UncompressedPublicKeyLength = 65
const secp256k1SignatureLength = 64

// CryptoImports adds some crypto imports to the Wasmer Imports map
func CryptoImports(imports *wasmer.Imports) (*wasmer.Imports, error) {
	imports = imports.Namespace("env")
	imports, err := imports.Append("sha256", sha256, C.sha256)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("keccak256", keccak256, C.keccak256)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("ripemd160", ripemd160, C.ripemd160)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("verifyBLS", verifyBLS, C.verifyBLS)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("verifyEd25519", verifyEd25519, C.verifyEd25519)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("verifySecp256k1", verifySecp256k1, C.verifySecp256k1)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("addEC", addEC, C.addEC)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("doubleEC", doubleEC, C.doubleEC)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("isOnCurveEC", isOnCurveEC, C.isOnCurveEC)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("scalarBaseMultEC", scalarBaseMultEC, C.scalarBaseMultEC)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("scalarMultEC", scalarMultEC, C.scalarMultEC)
	if err != nil {
		return nil, err
	}

	return imports, nil
}

const maxBigIntByteLenForNormalCost = 32

func useExtraGasForOperations(metering arwen.MeteringContext, values []*big.Int) {
	for _, val := range values {
		byteLen := val.BitLen() / 8
		if byteLen > maxBigIntByteLenForNormalCost {
			metering.UseGas(math.MulUint64(uint64(byteLen), metering.GasSchedule().BaseOperationCost.DataCopyPerByte))
		}
	}
}

//export sha256
func sha256(context unsafe.Pointer, dataOffset int32, length int32, resultOffset int32) int32 {
	runtime := arwen.GetRuntimeContext(context)
	crypto := arwen.GetCryptoContext(context)
	metering := arwen.GetMeteringContext(context)

	memLoadGas := math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(length))
	gasToUse := math.AddUint64(metering.GasSchedule().CryptoAPICost.SHA256, memLoadGas)
	metering.UseGas(gasToUse)

	data, err := runtime.MemLoad(dataOffset, length)
	if arwen.WithFault(err, context, runtime.CryptoAPIErrorShouldFailExecution()) {
		return 1
	}

	result, err := crypto.Sha256(data)
	if err != nil {
		return 1
	}

	err = runtime.MemStore(resultOffset, result)
	if arwen.WithFault(err, context, runtime.CryptoAPIErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

//export keccak256
func keccak256(context unsafe.Pointer, dataOffset int32, length int32, resultOffset int32) int32 {
	runtime := arwen.GetRuntimeContext(context)
	crypto := arwen.GetCryptoContext(context)
	metering := arwen.GetMeteringContext(context)

	memLoadGas := math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(length))
	gasToUse := math.AddUint64(metering.GasSchedule().CryptoAPICost.Keccak256, memLoadGas)
	metering.UseGas(gasToUse)

	data, err := runtime.MemLoad(dataOffset, length)
	if arwen.WithFault(err, context, runtime.CryptoAPIErrorShouldFailExecution()) {
		return 1
	}

	result, err := crypto.Keccak256(data)
	if err != nil {
		return 1
	}

	err = runtime.MemStore(resultOffset, result)
	if arwen.WithFault(err, context, runtime.CryptoAPIErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

//export ripemd160
func ripemd160(context unsafe.Pointer, dataOffset int32, length int32, resultOffset int32) int32 {
	runtime := arwen.GetRuntimeContext(context)
	crypto := arwen.GetCryptoContext(context)
	metering := arwen.GetMeteringContext(context)

	memLoadGas := math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(length))
	gasToUse := math.AddUint64(metering.GasSchedule().CryptoAPICost.Ripemd160, memLoadGas)
	metering.UseGas(gasToUse)

	data, err := runtime.MemLoad(dataOffset, length)
	if arwen.WithFault(err, context, runtime.CryptoAPIErrorShouldFailExecution()) {
		return 1
	}

	result, err := crypto.Ripemd160(data)
	if err != nil {
		return 1
	}

	err = runtime.MemStore(resultOffset, result)
	if arwen.WithFault(err, context, runtime.CryptoAPIErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

//export verifyBLS
func verifyBLS(
	context unsafe.Pointer,
	keyOffset int32,
	messageOffset int32,
	messageLength int32,
	sigOffset int32,
) int32 {
	runtime := arwen.GetRuntimeContext(context)
	crypto := arwen.GetCryptoContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().CryptoAPICost.VerifyBLS
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, blsPublicKeyLength)
	if arwen.WithFault(err, context, runtime.CryptoAPIErrorShouldFailExecution()) {
		return 1
	}

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(messageLength))
	metering.UseGas(gasToUse)

	message, err := runtime.MemLoad(messageOffset, messageLength)
	if arwen.WithFault(err, context, runtime.CryptoAPIErrorShouldFailExecution()) {
		return 1
	}

	sig, err := runtime.MemLoad(sigOffset, blsSignatureLength)
	if arwen.WithFault(err, context, runtime.CryptoAPIErrorShouldFailExecution()) {
		return 1
	}

	invalidSigErr := crypto.VerifyBLS(key, message, sig)
	if invalidSigErr != nil {
		return -1
	}

	return 0
}

//export verifyEd25519
func verifyEd25519(
	context unsafe.Pointer,
	keyOffset int32,
	messageOffset int32,
	messageLength int32,
	sigOffset int32,
) int32 {
	runtime := arwen.GetRuntimeContext(context)
	crypto := arwen.GetCryptoContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().CryptoAPICost.VerifyEd25519
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, ed25519PublicKeyLength)
	if arwen.WithFault(err, context, runtime.CryptoAPIErrorShouldFailExecution()) {
		return 1
	}

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(messageLength))
	metering.UseGas(gasToUse)

	message, err := runtime.MemLoad(messageOffset, messageLength)
	if arwen.WithFault(err, context, runtime.CryptoAPIErrorShouldFailExecution()) {
		return 1
	}

	sig, err := runtime.MemLoad(sigOffset, ed25519SignatureLength)
	if arwen.WithFault(err, context, runtime.CryptoAPIErrorShouldFailExecution()) {
		return 1
	}

	invalidSigErr := crypto.VerifyEd25519(key, message, sig)
	if invalidSigErr != nil {
		return -1
	}

	return 0
}

//export verifySecp256k1
func verifySecp256k1(
	context unsafe.Pointer,
	keyOffset int32,
	keyLength int32,
	messageOffset int32,
	messageLength int32,
	sigOffset int32,
) int32 {
	runtime := arwen.GetRuntimeContext(context)
	crypto := arwen.GetCryptoContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().CryptoAPICost.VerifySecp256k1
	metering.UseGas(gasToUse)

	if keyLength != secp256k1CompressedPublicKeyLength && keyLength != secp256k1UncompressedPublicKeyLength {
		arwen.WithFault(arwen.ErrInvalidPublicKeySize, context, runtime.ElrondAPIErrorShouldFailExecution())
		return 1
	}

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if arwen.WithFault(err, context, runtime.CryptoAPIErrorShouldFailExecution()) {
		return 1
	}

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(messageLength))
	metering.UseGas(gasToUse)

	message, err := runtime.MemLoad(messageOffset, messageLength)
	if arwen.WithFault(err, context, runtime.CryptoAPIErrorShouldFailExecution()) {
		return 1
	}

	// read the 2 leading bytes first
	// byte1: 0x30, header
	// byte2: the remaining buffer length
	const sigHeaderLength = 2
	sigHeader, err := runtime.MemLoad(sigOffset, sigHeaderLength)
	if arwen.WithFault(err, context, runtime.CryptoAPIErrorShouldFailExecution()) {
		return 1
	}
	sigLength := int32(sigHeader[1]) + sigHeaderLength
	sig, err := runtime.MemLoad(sigOffset, sigLength)
	if arwen.WithFault(err, context, runtime.CryptoAPIErrorShouldFailExecution()) {
		return 1
	}

	invalidSigErr := crypto.VerifySecp256k1(key, message, sig)
	if invalidSigErr != nil {
		return -1
	}

	return 0
}

//export addEC
func addEC(
	context unsafe.Pointer,
	destination1 int32,
	destination2 int32,
	fieldOrder int32,
	basePointOrder int32,
	eqConstant int32,
	xBasePoint int32,
	yBasePoint int32,
	sizeOfField int32,
	fstPointX int32,
	fstPointY int32,
	sndPointX int32,
	sndPointY int32,
) {
	bigInt := arwen.GetBigIntContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := uint64(0)
	metering.UseGas(gasToUse)

	dest1, dest2, P := bigInt.GetThree(destination1, destination2, fieldOrder)
	N, B, Gx := bigInt.GetThree(basePointOrder, eqConstant, xBasePoint)
	Gy, x1, y1 := bigInt.GetThree(yBasePoint, fstPointX, fstPointY)
	x2, y2 := bigInt.GetTwo(sndPointX, sndPointY)
	useExtraGasForOperations(metering, []*big.Int{dest1, dest2, P, N, B, Gx, Gy, x1, y1, x2, y2})

	dest1, dest2 = elliptic_curve.Add(P, N, B, Gx, Gy, int(sizeOfField), x1, y1, x2, y2)
}

//export doubleEC
func doubleEC(
	context unsafe.Pointer,
	destination1 int32,
	destination2 int32,
	fieldOrder int32,
	basePointOrder int32,
	eqConstant int32,
	xBasePoint int32,
	yBasePoint int32,
	sizeOfField int32,
	pointX int32,
	pointY int32,
) {
	bigInt := arwen.GetBigIntContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := uint64(0)
	metering.UseGas(gasToUse)

	dest1, dest2, P := bigInt.GetThree(destination1, destination2, fieldOrder)
	N, B, Gx := bigInt.GetThree(basePointOrder, eqConstant, xBasePoint)
	Gy, x1, y1 := bigInt.GetThree(yBasePoint, pointX, pointY)
	useExtraGasForOperations(metering, []*big.Int{dest1, dest2, P, N, B, Gx, Gy, x1, y1})

	dest1, dest2 = elliptic_curve.Double(P, N, B, Gx, Gy, int(sizeOfField), x1, y1)
}

//export isOnCurveEC
func isOnCurveEC(
	context unsafe.Pointer,
	fieldOrder int32,
	basePointOrder int32,
	eqConstant int32,
	xBasePoint int32,
	yBasePoint int32,
	sizeOfField int32,
	pointX int32,
	pointY int32,
) int32 {
	bigInt := arwen.GetBigIntContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := uint64(0)
	metering.UseGas(gasToUse)

	x, y, P := bigInt.GetThree(pointX, pointY, fieldOrder)
	N, B, Gx := bigInt.GetThree(basePointOrder, eqConstant, xBasePoint)
	Gy := bigInt.GetOne(yBasePoint)
	useExtraGasForOperations(metering, []*big.Int{P, N, B, Gx, Gy, x, y})

	if elliptic_curve.IsOnCurve(P, N, B, Gx, Gy, int(sizeOfField), x, y) {
		return 1
	}

	return 0

}

//export scalarBaseMultEC
func scalarBaseMultEC(
	context unsafe.Pointer,
	destination1 int32,
	destination2 int32,
	fieldOrder int32,
	basePointOrder int32,
	eqConstant int32,
	xBasePoint int32,
	yBasePoint int32,
	sizeOfField int32,
	kOffset int32,
	length int32,
) {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)
	bigInt := arwen.GetBigIntContext(context)

	gasToUse := uint64(0)
	metering.UseGas(gasToUse)

	k, err := runtime.MemLoad(kOffset, length)
	if arwen.WithFault(err, context, runtime.CryptoAPIErrorShouldFailExecution()) {
		return
	}

	dest1, dest2, P := bigInt.GetThree(destination1, destination2, fieldOrder)
	N, B, Gx := bigInt.GetThree(basePointOrder, eqConstant, xBasePoint)
	Gy := bigInt.GetOne(yBasePoint)
	useExtraGasForOperations(metering, []*big.Int{dest1, dest2, P, N, B, Gx, Gy})

	dest1, dest2 = elliptic_curve.ScalarBaseMult(P, N, B, Gx, Gy, int(sizeOfField), k)
}

//export scalarMultEC
func scalarMultEC(
	context unsafe.Pointer,
	destination1 int32,
	destination2 int32,
	fieldOrder int32,
	basePointOrder int32,
	eqConstant int32,
	xBasePoint int32,
	yBasePoint int32,
	sizeOfField int32,
	pointX int32,
	pointY int32,
	kOffset int32,
	length int32,
) {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)
	bigInt := arwen.GetBigIntContext(context)

	gasToUse := uint64(0)
	metering.UseGas(gasToUse)

	k, err := runtime.MemLoad(kOffset, length)
	if arwen.WithFault(err, context, runtime.CryptoAPIErrorShouldFailExecution()) {
		return
	}

	dest1, dest2, P := bigInt.GetThree(destination1, destination2, fieldOrder)
	N, B, Gx := bigInt.GetThree(basePointOrder, eqConstant, xBasePoint)
	Gy, x, y := bigInt.GetThree(yBasePoint, pointX, pointY)
	useExtraGasForOperations(metering, []*big.Int{dest1, dest2, P, N, B, Gx, Gy, x, y})

	dest1, dest2 = elliptic_curve.ScalarMult(P, N, B, Gx, Gy, int(sizeOfField), x, y, k)
}
