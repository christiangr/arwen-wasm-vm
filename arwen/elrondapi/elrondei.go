package elrondapi

// // Declare the function signatures (see [cgo](https://golang.org/cmd/cgo/)).
//
// #include <stdlib.h>
// typedef unsigned char uint8_t;
// typedef int int32_t;
//
// extern void		v1_4_getSCAddress(void *context, int32_t resultOffset);
// extern void		v1_4_getOwnerAddress(void *context, int32_t resultOffset);
// extern int32_t	v1_4_getShardOfAddress(void *context, int32_t addressOffset);
// extern int32_t	v1_4_isSmartContract(void *context, int32_t addressOffset);
// extern void		v1_4_getExternalBalance(void *context, int32_t addressOffset, int32_t resultOffset);
// extern int32_t	v1_4_blockHash(void *context, long long nonce, int32_t resultOffset);
// extern int32_t	v1_4_transferValue(void *context, int32_t dstOffset, int32_t valueOffset, int32_t dataOffset, int32_t length);
// extern int32_t	v1_4_transferESDTExecute(void *context, int32_t dstOffset, int32_t tokenIDOffset, int32_t tokenIdLen, int32_t valueOffset, long long gasLimit, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t	v1_4_transferESDTNFTExecute(void *context, int32_t dstOffset, int32_t tokenIDOffset, int32_t tokenIdLen, int32_t valueOffset, long long nonce, long long gasLimit, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t	v1_4_multiTransferESDTNFTExecute(void *context, int32_t dstOffset, int32_t numTokenTransfers, int32_t tokenTransfersArgsLengthOffset, int32_t tokenTransferDataOffset, long long gasLimit, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t	v1_4_transferValueExecute(void *context, int32_t dstOffset, int32_t valueOffset, long long gasLimit, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t	v1_4_getArgumentLength(void *context, int32_t id);
// extern int32_t	v1_4_getArgument(void *context, int32_t id, int32_t argOffset);
// extern int32_t	v1_4_getFunction(void *context, int32_t functionOffset);
// extern int32_t	v1_4_getNumArguments(void *context);
// extern int32_t	v1_4_storageStore(void *context, int32_t keyOffset, int32_t keyLength , int32_t dataOffset, int32_t dataLength);
// extern int32_t	v1_4_storageLoadLength(void *context, int32_t keyOffset, int32_t keyLength );
// extern int32_t	v1_4_storageLoad(void *context, int32_t keyOffset, int32_t keyLength , int32_t dataOffset);
// extern int32_t	v1_4_storageLoadFromAddress(void *context, int32_t addressOffset, int32_t keyOffset, int32_t keyLength , int32_t dataOffset);
// extern void		v1_4_getCaller(void *context, int32_t resultOffset);
// extern void		v1_4_checkNoPayment(void *context);
// extern int32_t	v1_4_callValue(void *context, int32_t resultOffset);
// extern int32_t	v1_4_getESDTValue(void *context, int32_t resultOffset);
// extern int32_t	v1_4_getESDTTokenName(void *context, int32_t resultOffset);
// extern long long	v1_4_getESDTTokenNonce(void *context);
// extern int32_t	v1_4_getESDTTokenType(void *context);
// extern int32_t	v1_4_getCallValueTokenName(void *context, int32_t callValueOffset, int32_t tokenNameOffset);
// extern int32_t	v1_4_getESDTValueByIndex(void *context, int32_t resultOffset, int32_t index);
// extern int32_t	v1_4_getESDTTokenNameByIndex(void *context, int32_t resultOffset, int32_t index);
// extern long long	v1_4_getESDTTokenNonceByIndex(void *context, int32_t index);
// extern int32_t	v1_4_getESDTTokenTypeByIndex(void *context, int32_t index);
// extern int32_t	v1_4_getCallValueTokenNameByIndex(void *context, int32_t callValueOffset, int32_t tokenNameOffset, int32_t index);
// extern int32_t	v1_4_getNumESDTTransfers(void *context);
// extern long long v1_4_getCurrentESDTNFTNonce(void *context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen);
// extern void		v1_4_writeLog(void *context, int32_t pointer, int32_t length, int32_t topicPtr, int32_t numTopics);
// extern void		v1_4_writeEventLog(void *context, int32_t numTopics, int32_t topicLengthsOffset, int32_t topicOffset, int32_t dataOffset, int32_t dataLength);
// extern void		v1_4_returnData(void* context, int32_t dataOffset, int32_t length);
// extern void		v1_4_signalError(void* context, int32_t messageOffset, int32_t messageLength);
// extern long long v1_4_getGasLeft(void *context);
// extern int32_t	v1_4_getESDTBalance(void *context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen, long long nonce, int32_t resultOffset);
// extern int32_t	v1_4_getESDTNFTNameLength(void *context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen, long long nonce);
// extern int32_t	v1_4_getESDTNFTAttributeLength(void *context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen, long long nonce);
// extern int32_t	v1_4_getESDTNFTURILength(void *context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen, long long nonce);
// extern int32_t	v1_4_getESDTTokenData(void *context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen, long long nonce, int32_t valueOffset, int32_t propertiesOffset, int32_t hashOffset, int32_t nameOffset, int32_t attributesOffset, int32_t creatorOffset, int32_t royaltiesOffset, int32_t urisOffset);
//
// extern int32_t	v1_4_executeOnDestContext(void *context, long long gas, int32_t addressOffset, int32_t valueOffset, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t	v1_4_executeOnDestContextByCaller(void *context, long long gas, int32_t addressOffset, int32_t valueOffset, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t	v1_4_executeOnSameContext(void *context, long long gas, int32_t addressOffset, int32_t valueOffset, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t	v1_4_delegateExecution(void *context, long long gas, int32_t addressOffset, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t	v1_4_executeReadOnly(void *context, long long gas, int32_t addressOffset, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t	v1_4_createContract(void *context, long long gas, int32_t valueOffset, int32_t codeOffset, int32_t codeMetadataOffset, int32_t length, int32_t resultOffset, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t	v1_4_deployFromSourceContract(void *context, long long gas, int32_t valueOffset, int32_t addressOffset, int32_t codeMetadataOffset, int32_t resultOffset, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern void		v1_4_upgradeContract(void *context, int32_t dstOffset, long long gas, int32_t valueOffset, int32_t codeOffset, int32_t codeMetadataOffset, int32_t length, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern void		v1_4_upgradeFromSourceContract(void *context, int32_t dstOffset, long long gas, int32_t valueOffset, int32_t addressOffset, int32_t codeMetadataOffset, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern void		v1_4_asyncCall(void *context, int32_t dstOffset, int32_t valueOffset, int32_t dataOffset, int32_t length);
// extern void		v1_4_createAsyncCall(void *context, int32_t identifierOffset, int32_t identifierLength, int32_t dstOffset, int32_t valueOffset, int32_t dataOffset, int32_t length, int32_t successCallback, int32_t successLength, int32_t errorCallback, int32_t errorLength, long long gas);
// extern int32_t	v1_4_setAsyncContextCallback(void *context, int32_t identifierOffset, int32_t identifierLength, int32_t callback, int32_t callbackLength);
//
// extern int32_t	v1_4_getNumReturnData(void *context);
// extern int32_t	v1_4_getReturnDataSize(void *context, int32_t resultID);
// extern int32_t	v1_4_getReturnData(void *context, int32_t resultID, int32_t dataOffset);
//
// extern int32_t	v1_4_setStorageLock(void *context, int32_t keyOffset, int32_t keyLength, long long lockTimestamp);
// extern long long v1_4_getStorageLock(void *context, int32_t keyOffset, int32_t keyLength);
// extern int32_t	v1_4_isStorageLocked(void *context, int32_t keyOffset, int32_t keyLength);
// extern int32_t	v1_4_clearStorageLock(void *context, int32_t keyOffset, int32_t keyLength);
//
// extern long long v1_4_getBlockTimestamp(void *context);
// extern long long v1_4_getBlockNonce(void *context);
// extern long long v1_4_getBlockRound(void *context);
// extern long long v1_4_getBlockEpoch(void *context);
// extern void		v1_4_getBlockRandomSeed(void *context, int32_t resultOffset);
// extern void		v1_4_getStateRootHash(void *context, int32_t resultOffset);
//
// extern long long v1_4_getPrevBlockTimestamp(void *context);
// extern long long v1_4_getPrevBlockNonce(void *context);
// extern long long v1_4_getPrevBlockRound(void *context);
// extern long long v1_4_getPrevBlockEpoch(void *context);
// extern void		v1_4_getPrevBlockRandomSeed(void *context, int32_t resultOffset);
// extern void		v1_4_getOriginalTxHash(void *context, int32_t resultOffset);
import "C"

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"unsafe"

	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/arwen"
	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/math"
	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/wasmer"
	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/data/esdt"
	"github.com/ElrondNetwork/elrond-go-core/data/vm"
	logger "github.com/ElrondNetwork/elrond-go-logger"
	vmcommon "github.com/ElrondNetwork/elrond-vm-common"
	"github.com/ElrondNetwork/elrond-vm-common/parsers"
)

var logEEI = logger.GetOrCreate("arwen/eei")

func getESDTTransferFromInput(vmInput *vmcommon.VMInput, index int32) *vmcommon.ESDTTransfer {
	esdtTransfers := vmInput.ESDTTransfers
	if int32(len(esdtTransfers))-1 < index {
		return nil
	}
	return esdtTransfers[index]
}

func failIfMoreThanOneESDTTransfer(context unsafe.Pointer) bool {
	runtime := arwen.GetRuntimeContext(context)
	if len(runtime.GetVMInput().ESDTTransfers) > 1 {
		return arwen.WithFault(arwen.ErrTooManyESDTTransfers, context, true)
	}
	return false
}

// ElrondEIImports creates a new wasmer.Imports populated with the ElrondEI API methods
func ElrondEIImports() (*wasmer.Imports, error) {
	imports := wasmer.NewImports()
	imports = imports.Namespace("env")

	imports, err := imports.Append("getSCAddress", v1_4_getSCAddress, C.v1_4_getSCAddress)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getOwnerAddress", v1_4_getOwnerAddress, C.v1_4_getOwnerAddress)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getShardOfAddress", v1_4_getShardOfAddress, C.v1_4_getShardOfAddress)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("isSmartContract", v1_4_isSmartContract, C.v1_4_isSmartContract)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getExternalBalance", v1_4_getExternalBalance, C.v1_4_getExternalBalance)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getBlockHash", v1_4_blockHash, C.v1_4_blockHash)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("transferValue", v1_4_transferValue, C.v1_4_transferValue)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("transferESDTExecute", v1_4_transferESDTExecute, C.v1_4_transferESDTExecute)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("transferESDTNFTExecute", v1_4_transferESDTNFTExecute, C.v1_4_transferESDTNFTExecute)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("multiTransferESDTNFTExecute", v1_4_multiTransferESDTNFTExecute, C.v1_4_multiTransferESDTNFTExecute)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("transferValueExecute", v1_4_transferValueExecute, C.v1_4_transferValueExecute)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("asyncCall", v1_4_asyncCall, C.v1_4_asyncCall)
	if err != nil {
		return nil, err
	}

	// imports, err = imports.Append("createAsyncCall", createAsyncCall, C.createAsyncCall)
	// if err != nil {
	// 	return nil, err
	// }

	// imports, err = imports.Append("setAsyncContextCallback", setAsyncContextCallback, C.setAsyncContextCallback)
	// if err != nil {
	// 	return nil, err
	// }

	imports, err = imports.Append("getArgumentLength", v1_4_getArgumentLength, C.v1_4_getArgumentLength)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getArgument", v1_4_getArgument, C.v1_4_getArgument)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getFunction", v1_4_getFunction, C.v1_4_getFunction)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getNumArguments", v1_4_getNumArguments, C.v1_4_getNumArguments)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("storageStore", v1_4_storageStore, C.v1_4_storageStore)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("storageLoadLength", v1_4_storageLoadLength, C.v1_4_storageLoadLength)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("storageLoad", v1_4_storageLoad, C.v1_4_storageLoad)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("storageLoadFromAddress", v1_4_storageLoadFromAddress, C.v1_4_storageLoadFromAddress)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getStorageLock", v1_4_getStorageLock, C.v1_4_getStorageLock)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("setStorageLock", v1_4_setStorageLock, C.v1_4_setStorageLock)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("isStorageLocked", v1_4_isStorageLocked, C.v1_4_isStorageLocked)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("clearStorageLock", v1_4_clearStorageLock, C.v1_4_clearStorageLock)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getCaller", v1_4_getCaller, C.v1_4_getCaller)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("checkNoPayment", v1_4_checkNoPayment, C.v1_4_checkNoPayment)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getCallValue", v1_4_callValue, C.v1_4_callValue)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTValue", v1_4_getESDTValue, C.v1_4_getESDTValue)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTTokenName", v1_4_getESDTTokenName, C.v1_4_getESDTTokenName)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTTokenType", v1_4_getESDTTokenType, C.v1_4_getESDTTokenType)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTTokenNonce", v1_4_getESDTTokenNonce, C.v1_4_getESDTTokenNonce)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getCallValueTokenName", v1_4_getCallValueTokenName, C.v1_4_getCallValueTokenName)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTValueByIndex", v1_4_getESDTValueByIndex, C.v1_4_getESDTValueByIndex)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTTokenNameByIndex", v1_4_getESDTTokenNameByIndex, C.v1_4_getESDTTokenNameByIndex)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTTokenTypeByIndex", v1_4_getESDTTokenTypeByIndex, C.v1_4_getESDTTokenTypeByIndex)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTTokenNonceByIndex", v1_4_getESDTTokenNonceByIndex, C.v1_4_getESDTTokenNonceByIndex)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getCallValueTokenNameByIndex", v1_4_getCallValueTokenNameByIndex, C.v1_4_getCallValueTokenNameByIndex)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getNumESDTTransfers", v1_4_getNumESDTTransfers, C.v1_4_getNumESDTTransfers)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getCurrentESDTNFTNonce", v1_4_getCurrentESDTNFTNonce, C.v1_4_getCurrentESDTNFTNonce)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("writeLog", v1_4_writeLog, C.v1_4_writeLog)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("writeEventLog", v1_4_writeEventLog, C.v1_4_writeEventLog)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("finish", v1_4_returnData, C.v1_4_returnData)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("signalError", v1_4_signalError, C.v1_4_signalError)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getBlockTimestamp", v1_4_getBlockTimestamp, C.v1_4_getBlockTimestamp)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getBlockNonce", v1_4_getBlockNonce, C.v1_4_getBlockNonce)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getBlockRound", v1_4_getBlockRound, C.v1_4_getBlockRound)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getBlockEpoch", v1_4_getBlockEpoch, C.v1_4_getBlockEpoch)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getBlockRandomSeed", v1_4_getBlockRandomSeed, C.v1_4_getBlockRandomSeed)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getStateRootHash", v1_4_getStateRootHash, C.v1_4_getStateRootHash)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getPrevBlockTimestamp", v1_4_getPrevBlockTimestamp, C.v1_4_getPrevBlockTimestamp)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getPrevBlockNonce", v1_4_getPrevBlockNonce, C.v1_4_getPrevBlockNonce)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getPrevBlockRound", v1_4_getPrevBlockRound, C.v1_4_getPrevBlockRound)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getPrevBlockEpoch", v1_4_getPrevBlockEpoch, C.v1_4_getPrevBlockEpoch)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getPrevBlockRandomSeed", v1_4_getPrevBlockRandomSeed, C.v1_4_getPrevBlockRandomSeed)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getOriginalTxHash", v1_4_getOriginalTxHash, C.v1_4_getOriginalTxHash)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getGasLeft", v1_4_getGasLeft, C.v1_4_getGasLeft)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("executeOnDestContext", v1_4_executeOnDestContext, C.v1_4_executeOnDestContext)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("executeOnDestContextByCaller", v1_4_executeOnDestContextByCaller, C.v1_4_executeOnDestContextByCaller)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("executeOnSameContext", v1_4_executeOnSameContext, C.v1_4_executeOnSameContext)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("delegateExecution", v1_4_delegateExecution, C.v1_4_delegateExecution)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("createContract", v1_4_createContract, C.v1_4_createContract)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("deployFromSourceContract", v1_4_deployFromSourceContract, C.v1_4_deployFromSourceContract)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("upgradeContract", v1_4_upgradeContract, C.v1_4_upgradeContract)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("upgradeFromSourceContract", v1_4_upgradeFromSourceContract, C.v1_4_upgradeFromSourceContract)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("executeReadOnly", v1_4_executeReadOnly, C.v1_4_executeReadOnly)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getNumReturnData", v1_4_getNumReturnData, C.v1_4_getNumReturnData)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getReturnDataSize", v1_4_getReturnDataSize, C.v1_4_getReturnDataSize)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getReturnData", v1_4_getReturnData, C.v1_4_getReturnData)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTBalance", v1_4_getESDTBalance, C.v1_4_getESDTBalance)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTTokenData", v1_4_getESDTTokenData, C.v1_4_getESDTTokenData)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTNFTNameLength", v1_4_getESDTNFTNameLength, C.v1_4_getESDTNFTNameLength)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTNFTAttributeLength", v1_4_getESDTNFTAttributeLength, C.v1_4_getESDTNFTAttributeLength)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getESDTNFTURILength", v1_4_getESDTNFTURILength, C.v1_4_getESDTNFTURILength)
	if err != nil {
		return nil, err
	}

	return imports, nil
}

//export v1_4_getGasLeft
func v1_4_getGasLeft(context unsafe.Pointer) int64 {
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetGasLeft
	metering.UseGas(gasToUse)

	return int64(metering.GasLeft())
}

//export v1_4_getSCAddress
func v1_4_getSCAddress(context unsafe.Pointer, resultOffset int32) {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetSCAddress
	metering.UseGas(gasToUse)

	owner := runtime.GetSCAddress()
	err := runtime.MemStore(resultOffset, owner)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}
}

//export v1_4_getOwnerAddress
func v1_4_getOwnerAddress(context unsafe.Pointer, resultOffset int32) {
	blockchain := arwen.GetBlockchainContext(context)
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetOwnerAddress
	metering.UseGas(gasToUse)

	owner, err := blockchain.GetOwnerAddress()
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	err = runtime.MemStore(resultOffset, owner)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}
}

//export v1_4_getShardOfAddress
func v1_4_getShardOfAddress(context unsafe.Pointer, addressOffset int32) int32 {
	blockchain := arwen.GetBlockchainContext(context)
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetShardOfAddress
	metering.UseGas(gasToUse)

	address, err := runtime.MemLoad(addressOffset, arwen.AddressLen)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 0
	}

	return int32(blockchain.GetShardOfAddress(address))
}

//export v1_4_isSmartContract
func v1_4_isSmartContract(context unsafe.Pointer, addressOffset int32) int32 {
	blockchain := arwen.GetBlockchainContext(context)
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.IsSmartContract
	metering.UseGas(gasToUse)

	address, err := runtime.MemLoad(addressOffset, arwen.AddressLen)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 0
	}

	isSmartContract := blockchain.IsSmartContract(address)

	return int32(arwen.BooleanToInt(isSmartContract))
}

//export v1_4_signalError
func v1_4_signalError(context unsafe.Pointer, messageOffset int32, messageLength int32) {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.SignalError
	metering.UseGas(gasToUse)

	message, err := runtime.MemLoad(messageOffset, messageLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}
	runtime.SignalUserError(string(message))
}

//export v1_4_getExternalBalance
func v1_4_getExternalBalance(context unsafe.Pointer, addressOffset int32, resultOffset int32) {
	blockchain := arwen.GetBlockchainContext(context)
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetExternalBalance
	metering.UseGas(gasToUse)

	address, err := runtime.MemLoad(addressOffset, arwen.AddressLen)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	balance := blockchain.GetBalance(address)

	err = runtime.MemStore(resultOffset, balance)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}
}

//export v1_4_blockHash
func v1_4_blockHash(context unsafe.Pointer, nonce int64, resultOffset int32) int32 {
	blockchain := arwen.GetBlockchainContext(context)
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetBlockHash
	metering.UseGas(gasToUse)

	hash := blockchain.BlockHash(nonce)
	err := runtime.MemStore(resultOffset, hash)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

func isBuiltInCall(data string, host arwen.VMHost) bool {
	argParser := parsers.NewCallArgsParser()
	functionName, _, _ := argParser.ParseData(data)
	return host.IsBuiltinFunctionName(functionName)
}

func getESDTDataFromBlockchainHook(
	context unsafe.Pointer,
	addressOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	nonce int64,
) (*esdt.ESDigitalToken, error) {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)
	blockchain := arwen.GetBlockchainContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetExternalBalance
	metering.UseGas(gasToUse)

	address, err := runtime.MemLoad(addressOffset, arwen.AddressLen)
	if err != nil {
		return nil, err
	}

	tokenID, err := runtime.MemLoad(tokenIDOffset, tokenIDLen)
	if err != nil {
		return nil, err
	}

	esdtToken, err := blockchain.GetESDTToken(address, tokenID, uint64(nonce))
	if err != nil {
		return nil, err
	}

	return esdtToken, nil
}

//export v1_4_getESDTBalance
func v1_4_getESDTBalance(
	context unsafe.Pointer,
	addressOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	nonce int64,
	resultOffset int32,
) int32 {
	runtime := arwen.GetRuntimeContext(context)
	esdtData, err := getESDTDataFromBlockchainHook(context, addressOffset, tokenIDOffset, tokenIDLen, nonce)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 0
	}
	err = runtime.MemStore(resultOffset, esdtData.Value.Bytes())
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 0
	}

	return int32(len(esdtData.Value.Bytes()))
}

//export v1_4_getESDTNFTNameLength
func v1_4_getESDTNFTNameLength(
	context unsafe.Pointer,
	addressOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	nonce int64,
) int32 {
	runtime := arwen.GetRuntimeContext(context)
	esdtData, err := getESDTDataFromBlockchainHook(context, addressOffset, tokenIDOffset, tokenIDLen, nonce)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 0
	}
	if esdtData == nil || esdtData.TokenMetaData == nil {
		return 0
	}

	return int32(len(esdtData.TokenMetaData.Name))
}

//export v1_4_getESDTNFTAttributeLength
func v1_4_getESDTNFTAttributeLength(
	context unsafe.Pointer,
	addressOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	nonce int64,
) int32 {
	runtime := arwen.GetRuntimeContext(context)
	esdtData, err := getESDTDataFromBlockchainHook(context, addressOffset, tokenIDOffset, tokenIDLen, nonce)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 0
	}
	if esdtData == nil || esdtData.TokenMetaData == nil {
		return 0
	}

	return int32(len(esdtData.TokenMetaData.Attributes))
}

//export v1_4_getESDTNFTURILength
func v1_4_getESDTNFTURILength(
	context unsafe.Pointer,
	addressOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	nonce int64,
) int32 {
	runtime := arwen.GetRuntimeContext(context)
	esdtData, err := getESDTDataFromBlockchainHook(context, addressOffset, tokenIDOffset, tokenIDLen, nonce)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 0
	}
	if esdtData == nil || esdtData.TokenMetaData == nil {
		return 0
	}
	if len(esdtData.TokenMetaData.URIs) == 0 {
		return 0
	}

	return int32(len(esdtData.TokenMetaData.URIs[0]))
}

//export v1_4_getESDTTokenData
func v1_4_getESDTTokenData(
	context unsafe.Pointer,
	addressOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	nonce int64,
	valueHandle int32,
	propertiesOffset int32,
	hashOffset int32,
	nameOffset int32,
	attributesOffset int32,
	creatorOffset int32,
	royaltiesHandle int32,
	urisOffset int32,
) int32 {
	managedType := arwen.GetManagedTypesContext(context)
	runtime := arwen.GetRuntimeContext(context)
	esdtData, err := getESDTDataFromBlockchainHook(context, addressOffset, tokenIDOffset, tokenIDLen, nonce)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 0
	}

	value := managedType.GetBigIntOrCreate(valueHandle)
	value.Set(esdtData.Value)

	err = runtime.MemStore(propertiesOffset, esdtData.Properties)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 0
	}

	if esdtData.TokenMetaData != nil {
		err = runtime.MemStore(hashOffset, esdtData.TokenMetaData.Hash)
		if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
			return 0
		}
		err = runtime.MemStore(nameOffset, esdtData.TokenMetaData.Name)
		if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
			return 0
		}
		err = runtime.MemStore(attributesOffset, esdtData.TokenMetaData.Attributes)
		if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
			return 0
		}
		err = runtime.MemStore(creatorOffset, esdtData.TokenMetaData.Creator)
		if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
			return 0
		}

		royalties := managedType.GetBigIntOrCreate(royaltiesHandle)
		royalties.SetUint64(uint64(esdtData.TokenMetaData.Royalties))

		if len(esdtData.TokenMetaData.URIs) > 0 {
			err = runtime.MemStore(urisOffset, esdtData.TokenMetaData.URIs[0])
			if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
				return 0
			}
		}
	}
	return int32(len(esdtData.Value.Bytes()))
}

//export v1_4_transferValue
func v1_4_transferValue(context unsafe.Pointer, destOffset int32, valueOffset int32, dataOffset int32, length int32) int32 {
	host := arwen.GetVMHost(context)
	runtime := host.Runtime()
	metering := host.Metering()
	output := host.Output()

	gasToUse := metering.GasSchedule().ElrondAPICost.TransferValue
	metering.UseGas(gasToUse)

	sender := runtime.GetSCAddress()
	dest, err := runtime.MemLoad(destOffset, arwen.AddressLen)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	valueBytes, err := runtime.MemLoad(valueOffset, arwen.BalanceLen)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.PersistPerByte, uint64(length))
	metering.UseGas(gasToUse)

	data, err := runtime.MemLoad(dataOffset, length)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	if isBuiltInCall(string(data), host) {
		return 1
	}

	err = output.Transfer(dest, sender, 0, 0, big.NewInt(0).SetBytes(valueBytes), data, vm.DirectCall)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

type indirectContractCallArguments struct {
	dest      []byte
	value     *big.Int
	function  []byte
	args      [][]byte
	actualLen int32
}

func extractIndirectContractCallArgumentsWithValue(
	host arwen.VMHost,
	destOffset int32,
	valueOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) (*indirectContractCallArguments, error) {
	return extractIndirectContractCallArguments(
		host,
		destOffset,
		valueOffset,
		true,
		functionOffset,
		functionLength,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)
}

func extractIndirectContractCallArgumentsWithoutValue(
	host arwen.VMHost,
	destOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) (*indirectContractCallArguments, error) {
	return extractIndirectContractCallArguments(
		host,
		destOffset,
		0,
		false,
		functionOffset,
		functionLength,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)
}

func extractIndirectContractCallArguments(
	host arwen.VMHost,
	destOffset int32,
	valueOffset int32,
	hasValueOffset bool,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) (*indirectContractCallArguments, error) {
	runtime := host.Runtime()
	metering := host.Metering()

	dest, err := runtime.MemLoad(destOffset, arwen.AddressLen)
	if err != nil {
		return nil, err
	}

	var value *big.Int

	if hasValueOffset {
		valueBytes, err := runtime.MemLoad(valueOffset, arwen.BalanceLen)
		if err != nil {
			return nil, err
		}
		value = big.NewInt(0).SetBytes(valueBytes)
	}

	function, err := runtime.MemLoad(functionOffset, functionLength)
	if err != nil {
		return nil, err
	}

	args, actualLen, err := getArgumentsFromMemory(
		host,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)

	gasToUse := math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(actualLen))
	metering.UseGas(gasToUse)

	return &indirectContractCallArguments{
		dest:      dest,
		value:     value,
		function:  function,
		args:      args,
		actualLen: actualLen,
	}, nil
}

//export v1_4_transferValueExecute
func v1_4_transferValueExecute(
	context unsafe.Pointer,
	destOffset int32,
	valueOffset int32,
	gasLimit int64,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := arwen.GetVMHost(context)
	return TransferValueExecuteWithHost(
		host,
		destOffset,
		valueOffset,
		gasLimit,
		functionOffset,
		functionLength,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)
}

// TransferValueExecuteWithHost - transferValueExecute with host instead of pointer context
func TransferValueExecuteWithHost(
	host arwen.VMHost,
	destOffset int32,
	valueOffset int32,
	gasLimit int64,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().ElrondAPICost.TransferValue
	metering.UseGas(gasToUse)

	callArgs, err := extractIndirectContractCallArgumentsWithValue(
		host, destOffset, valueOffset, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	return TransferValueExecuteWithTypedArgs(
		host,
		callArgs.dest,
		callArgs.value,
		gasLimit,
		callArgs.function,
		callArgs.args,
	)
}

// TransferValueExecuteWithTypedArgs - transferValueExecute with args already read from memory
func TransferValueExecuteWithTypedArgs(
	host arwen.VMHost,
	dest []byte,
	value *big.Int,
	gasLimit int64,
	function []byte,
	args [][]byte,
) int32 {
	runtime := host.Runtime()
	metering := host.Metering()
	output := host.Output()

	gasToUse := metering.GasSchedule().ElrondAPICost.TransferValue
	metering.UseGas(gasToUse)

	sender := runtime.GetSCAddress()

	var err error
	var contractCallInput *vmcommon.ContractCallInput

	if len(function) > 0 {
		contractCallInput, err = prepareIndirectContractCallInput(
			host,
			sender,
			value,
			gasLimit,
			dest,
			function,
			args,
			gasToUse,
			false,
		)
		if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
			return 1
		}
	}

	if contractCallInput != nil {
		if host.IsBuiltinFunctionName(contractCallInput.Function) {
			return 1
		}
	}

	if host.AreInSameShard(sender, dest) && contractCallInput != nil && host.Blockchain().IsSmartContract(dest) {
		logEEI.Trace("eGLD pre-transfer execution begin")
		_, _, err = host.ExecuteOnDestContext(contractCallInput)
		if err != nil {
			logEEI.Trace("eGLD pre-transfer execution failed", "error", err)
			return 1
		}

		return 0
	}

	data := makeCrossShardCallFromInput(contractCallInput)
	err = output.Transfer(dest, sender, uint64(gasLimit), 0, value, []byte(data), vm.DirectCall)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

func makeCrossShardCallFromInput(vmInput *vmcommon.ContractCallInput) string {
	if vmInput == nil {
		return ""
	}

	txData := vmInput.Function
	for _, arg := range vmInput.Arguments {
		txData += "@" + hex.EncodeToString(arg)
	}

	return txData
}

//export v1_4_transferESDTExecute
func v1_4_transferESDTExecute(
	context unsafe.Pointer,
	destOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	valueOffset int32,
	gasLimit int64,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	return v1_4_transferESDTNFTExecute(context, destOffset, tokenIDOffset, tokenIDLen, valueOffset, 0,
		gasLimit, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
}

//export v1_4_transferESDTNFTExecute
func v1_4_transferESDTNFTExecute(
	context unsafe.Pointer,
	destOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	valueOffset int32,
	nonce int64,
	gasLimit int64,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := arwen.GetVMHost(context)
	return TransferESDTNFTExecuteWithHost(
		host,
		destOffset,
		tokenIDOffset,
		tokenIDLen,
		valueOffset,
		nonce,
		gasLimit,
		functionOffset,
		functionLength,
		numArguments,
		argumentsLengthOffset,
		dataOffset)
}

//export v1_4_multiTransferESDTNFTExecute
func v1_4_multiTransferESDTNFTExecute(
	context unsafe.Pointer,
	destOffset int32,
	numTokenTransfers int32,
	tokenTransfersArgsLengthOffset int32,
	tokenTransferDataOffset int32,
	gasLimit int64,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := arwen.GetVMHost(context)
	runtime := host.Runtime()
	metering := host.Metering()

	if numTokenTransfers == 0 {
		_ = arwen.WithFaultAndHost(host, arwen.ErrFailedTransfer, runtime.ElrondAPIErrorShouldFailExecution())
		return 1
	}

	callArgs, err := extractIndirectContractCallArgumentsWithoutValue(
		host, destOffset, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	gasToUse := math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(callArgs.actualLen))
	metering.UseGas(gasToUse)

	transferArgs, actualLen, err := getArgumentsFromMemory(
		host,
		numTokenTransfers*parsers.ArgsPerTransfer,
		tokenTransfersArgsLengthOffset,
		tokenTransferDataOffset,
	)

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(actualLen))
	metering.UseGas(gasToUse)

	transfers := make([]*vmcommon.ESDTTransfer, numTokenTransfers)
	for i := int32(0); i < numTokenTransfers; i++ {
		tokenStartIndex := i * parsers.ArgsPerTransfer
		transfer := &vmcommon.ESDTTransfer{
			ESDTTokenName:  transferArgs[tokenStartIndex],
			ESDTTokenNonce: big.NewInt(0).SetBytes(transferArgs[tokenStartIndex+1]).Uint64(),
			ESDTValue:      big.NewInt(0).SetBytes(transferArgs[tokenStartIndex+2]),
			ESDTTokenType:  uint32(core.Fungible),
		}
		if transfer.ESDTTokenNonce > 0 {
			transfer.ESDTTokenType = uint32(core.NonFungible)
		}
		transfers[i] = transfer
	}

	return TransferESDTNFTExecuteWithTypedArgs(
		host,
		callArgs.dest,
		transfers,
		gasLimit,
		callArgs.function,
		callArgs.args,
	)
}

// TransferESDTNFTExecuteWithHost contains only memory reading of arguments
func TransferESDTNFTExecuteWithHost(
	host arwen.VMHost,
	destOffset int32,
	tokenIDOffset int32,
	tokenIDLen int32,
	valueOffset int32,
	nonce int64,
	gasLimit int64,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	runtime := host.Runtime()
	metering := host.Metering()

	tokenIdentifier, executeErr := runtime.MemLoad(tokenIDOffset, tokenIDLen)
	if arwen.WithFaultAndHost(host, executeErr, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	callArgs, err := extractIndirectContractCallArgumentsWithValue(
		host, destOffset, valueOffset, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	gasToUse := math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(callArgs.actualLen))
	metering.UseGas(gasToUse)

	transfer := &vmcommon.ESDTTransfer{
		ESDTValue:      callArgs.value,
		ESDTTokenName:  tokenIdentifier,
		ESDTTokenNonce: uint64(nonce),
		ESDTTokenType:  uint32(core.Fungible),
	}
	if nonce > 0 {
		transfer.ESDTTokenType = uint32(core.NonFungible)
	}
	return TransferESDTNFTExecuteWithTypedArgs(
		host,
		callArgs.dest,
		[]*vmcommon.ESDTTransfer{transfer},
		gasLimit,
		callArgs.function,
		callArgs.args,
	)
}

// TransferESDTNFTExecuteWithTypedArgs defines the actual transfer ESDT execute logic
func TransferESDTNFTExecuteWithTypedArgs(
	host arwen.VMHost,
	dest []byte,
	transfers []*vmcommon.ESDTTransfer,
	gasLimit int64,
	function []byte,
	data [][]byte,
) int32 {

	var executeErr error

	runtime := host.Runtime()
	metering := host.Metering()

	output := host.Output()

	gasToUse := metering.GasSchedule().ElrondAPICost.TransferValue * uint64(len(transfers))
	metering.UseGas(gasToUse)

	sender := runtime.GetSCAddress()

	var contractCallInput *vmcommon.ContractCallInput
	if len(function) > 0 {
		contractCallInput, executeErr = prepareIndirectContractCallInput(
			host,
			sender,
			big.NewInt(0),
			gasLimit,
			dest,
			function,
			data,
			gasToUse,
			false,
		)
		if arwen.WithFaultAndHost(host, executeErr, runtime.ElrondSyncExecAPIErrorShouldFailExecution()) {
			return 1
		}

		contractCallInput.ESDTTransfers = transfers
	}

	snapshotBeforeTransfer := host.Blockchain().GetSnapshot()

	gasLimitForExec, executeErr := output.TransferESDT(dest, sender, transfers, contractCallInput)
	if arwen.WithFaultAndHost(host, executeErr, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	if host.AreInSameShard(sender, dest) && contractCallInput != nil && host.Blockchain().IsSmartContract(dest) {
		contractCallInput.GasProvided = gasLimitForExec
		logEEI.Trace("ESDT post-transfer execution begin")
		_, _, executeErr = host.ExecuteOnDestContext(contractCallInput)
		if executeErr != nil {
			logEEI.Trace("ESDT post-transfer execution failed", "error", executeErr)
			host.Blockchain().RevertToSnapshot(snapshotBeforeTransfer)
			return 1
		}

		return 0
	}

	return 0
}

//export v1_4_createAsyncCall
func v1_4_createAsyncCall(context unsafe.Pointer,
	asyncContextIdentifier int32,
	identifierLength int32,
	destOffset int32,
	valueOffset int32,
	dataOffset int32,
	length int32,
	successOffset int32,
	successLength int32,
	errorOffset int32,
	errorLength int32,
	gas int64,
) {
	host := arwen.GetVMHost(context)
	runtime := host.Runtime()

	// TODO consume gas

	acIdentifier, err := runtime.MemLoad(asyncContextIdentifier, identifierLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	calledSCAddress, err := runtime.MemLoad(destOffset, arwen.AddressLen)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	value, err := runtime.MemLoad(valueOffset, arwen.BalanceLen)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	data, err := runtime.MemLoad(dataOffset, length)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	successFunc, err := runtime.MemLoad(successOffset, successLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	errorFunc, err := runtime.MemLoad(errorOffset, errorLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	err = runtime.AddAsyncContextCall(acIdentifier, &arwen.AsyncGeneratedCall{
		Destination:     calledSCAddress,
		Data:            data,
		ValueBytes:      value,
		SuccessCallback: string(successFunc),
		ErrorCallback:   string(errorFunc),
		ProvidedGas:     uint64(gas),
	})
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}
}

//export v1_4_setAsyncContextCallback
func v1_4_setAsyncContextCallback(context unsafe.Pointer,
	asyncContextIdentifier int32,
	identifierLength int32,
	callback int32,
	callbackLength int32,
) int32 {
	host := arwen.GetVMHost(context)
	runtime := host.Runtime()

	// TODO consume gas

	acIdentifier, err := runtime.MemLoad(asyncContextIdentifier, identifierLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	asyncContext, err := runtime.GetAsyncContext(acIdentifier)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	callbackFunc, err := runtime.MemLoad(callback, callbackLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	asyncContext.Callback = string(callbackFunc)

	return 0
}

//export v1_4_upgradeContract
func v1_4_upgradeContract(
	context unsafe.Pointer,
	destOffset int32,
	gasLimit int64,
	valueOffset int32,
	codeOffset int32,
	codeMetadataOffset int32,
	length int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) {
	host := arwen.GetVMHost(context)
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().ElrondAPICost.CreateContract
	metering.UseGas(gasToUse)

	value, err := runtime.MemLoad(valueOffset, arwen.BalanceLen)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	code, err := runtime.MemLoad(codeOffset, length)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	codeMetadata, err := runtime.MemLoad(codeMetadataOffset, arwen.CodeMetadataLen)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	data, actualLen, err := getArgumentsFromMemory(
		host,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(actualLen))
	metering.UseGas(gasToUse)

	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	calledSCAddress, err := runtime.MemLoad(destOffset, arwen.AddressLen)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	gasSchedule := metering.GasSchedule()
	gasToUse = math.MulUint64(gasSchedule.BaseOperationCost.DataCopyPerByte, uint64(length))
	metering.UseGas(gasToUse)

	upgradeContract(host, calledSCAddress, code, codeMetadata, value, data, gasLimit)
}

//export v1_4_upgradeFromSourceContract
func v1_4_upgradeFromSourceContract(
	context unsafe.Pointer,
	destOffset int32,
	gasLimit int64,
	valueOffset int32,
	sourceContractAddressOffset int32,
	codeMetadataOffset int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) {
	host := arwen.GetVMHost(context)
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().ElrondAPICost.CreateContract
	metering.UseGas(gasToUse)

	value, err := runtime.MemLoad(valueOffset, arwen.BalanceLen)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	sourceContractAddress, err := runtime.MemLoad(sourceContractAddressOffset, arwen.AddressLen)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	codeMetadata, err := runtime.MemLoad(codeMetadataOffset, arwen.CodeMetadataLen)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	data, actualLen, err := getArgumentsFromMemory(
		host,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(actualLen))
	metering.UseGas(gasToUse)

	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	calledSCAddress, err := runtime.MemLoad(destOffset, arwen.AddressLen)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	UpgradeFromSourceContractWithTypedArgs(
		host,
		sourceContractAddress,
		calledSCAddress,
		value,
		data,
		gasLimit,
		codeMetadata,
	)
}

// UpgradeFromSourceContractWithTypedArgs - upgradeFromSourceContract with args already read from memory
func UpgradeFromSourceContractWithTypedArgs(
	host arwen.VMHost,
	sourceContractAddress []byte,
	destContractAddress []byte,
	value []byte,
	data [][]byte,
	gasLimit int64,
	codeMetadata []byte,
) {
	runtime := host.Runtime()
	blockchain := host.Blockchain()

	code, err := blockchain.GetCode(sourceContractAddress)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	upgradeContract(host, destContractAddress, code, codeMetadata, value, data, gasLimit)
}

func upgradeContract(
	host arwen.VMHost,
	destContractAddress []byte,
	code []byte,
	codeMetadata []byte,
	value []byte,
	data [][]byte,
	gasLimit int64,
) {
	runtime := host.Runtime()
	metering := host.Metering()
	gasSchedule := metering.GasSchedule()
	minAsyncCallCost := math.AddUint64(
		math.MulUint64(2, gasSchedule.ElrondAPICost.AsyncCallStep),
		gasSchedule.ElrondAPICost.AsyncCallbackGasLock)
	if uint64(gasLimit) < minAsyncCallCost {
		runtime.SetRuntimeBreakpointValue(arwen.BreakpointOutOfGas)
		return
	}

	// Set up the async call as if it is not known whether the called SC
	// is in the same shard with the caller or not. This will be later resolved
	// by runtime.ExecuteAsyncCall().
	callData := arwen.UpgradeFunctionName + "@" + hex.EncodeToString(code) + "@" + hex.EncodeToString(codeMetadata)
	for _, arg := range data {
		callData += "@" + hex.EncodeToString(arg)
	}

	err := runtime.ExecuteAsyncCall(
		destContractAddress,
		[]byte(callData),
		value,
	)
	logEEI.Trace("upgradeContract", "error", err)
}

//export v1_4_asyncCall
func v1_4_asyncCall(context unsafe.Pointer, destOffset int32, valueOffset int32, dataOffset int32, length int32) {
	host := arwen.GetVMHost(context)
	runtime := host.Runtime()
	metering := host.Metering()

	gasSchedule := metering.GasSchedule()
	gasToUse := gasSchedule.ElrondAPICost.AsyncCallStep
	metering.UseGas(gasToUse)

	calledSCAddress, err := runtime.MemLoad(destOffset, arwen.AddressLen)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	value, err := runtime.MemLoad(valueOffset, arwen.BalanceLen)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	gasToUse = math.MulUint64(gasSchedule.BaseOperationCost.DataCopyPerByte, uint64(length))
	metering.UseGas(gasToUse)

	data, err := runtime.MemLoad(dataOffset, length)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	err = runtime.ExecuteAsyncCall(calledSCAddress, data, value)
	if errors.Is(err, arwen.ErrNotEnoughGas) {
		runtime.SetRuntimeBreakpointValue(arwen.BreakpointOutOfGas)
		return
	}
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}
}

//export v1_4_getArgumentLength
func v1_4_getArgumentLength(context unsafe.Pointer, id int32) int32 {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetArgument
	metering.UseGas(gasToUse)

	args := runtime.Arguments()
	if id < 0 || int32(len(args)) <= id {
		return -1
	}

	return int32(len(args[id]))
}

//export v1_4_getArgument
func v1_4_getArgument(context unsafe.Pointer, id int32, argOffset int32) int32 {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetArgument
	metering.UseGas(gasToUse)

	args := runtime.Arguments()
	if id < 0 || int32(len(args)) <= id {
		return -1
	}

	err := runtime.MemStore(argOffset, args[id])
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	return int32(len(args[id]))
}

//export v1_4_getFunction
func v1_4_getFunction(context unsafe.Pointer, functionOffset int32) int32 {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetFunction
	metering.UseGas(gasToUse)

	function := runtime.Function()
	err := runtime.MemStore(functionOffset, []byte(function))
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	return int32(len(function))
}

//export v1_4_getNumArguments
func v1_4_getNumArguments(context unsafe.Pointer) int32 {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetNumArguments
	metering.UseGas(gasToUse)

	args := runtime.Arguments()
	return int32(len(args))
}

//export v1_4_storageStore
func v1_4_storageStore(context unsafe.Pointer, keyOffset int32, keyLength int32, dataOffset int32, dataLength int32) int32 {
	runtime := arwen.GetRuntimeContext(context)
	storage := arwen.GetStorageContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.StorageStore
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	data, err := runtime.MemLoad(dataOffset, dataLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	storageStatus, err := storage.SetStorage(key, data)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	return int32(storageStatus)
}

//export v1_4_storageLoadLength
func v1_4_storageLoadLength(context unsafe.Pointer, keyOffset int32, keyLength int32) int32 {
	runtime := arwen.GetRuntimeContext(context)
	storage := arwen.GetStorageContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.StorageLoad
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	data := storage.GetStorageUnmetered(key)

	return int32(len(data))
}

//export v1_4_storageLoadFromAddress
func v1_4_storageLoadFromAddress(context unsafe.Pointer, addressOffset int32, keyOffset int32, keyLength int32, dataOffset int32) int32 {
	runtime := arwen.GetRuntimeContext(context)
	storage := arwen.GetStorageContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.StorageLoad
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	address, err := runtime.MemLoad(addressOffset, arwen.AddressLen)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	data := storage.GetStorageFromAddress(address, key)

	err = runtime.MemStore(dataOffset, data)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	return int32(len(data))
}

//export v1_4_storageLoad
func v1_4_storageLoad(context unsafe.Pointer, keyOffset int32, keyLength int32, dataOffset int32) int32 {
	runtime := arwen.GetRuntimeContext(context)
	storage := arwen.GetStorageContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.StorageLoad
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	data := storage.GetStorage(key)

	err = runtime.MemStore(dataOffset, data)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	return int32(len(data))
}

//export v1_4_setStorageLock
func v1_4_setStorageLock(context unsafe.Pointer, keyOffset int32, keyLength int32, lockTimestamp int64) int32 {
	runtime := arwen.GetRuntimeContext(context)
	storage := arwen.GetStorageContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.Int64StorageStore
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	timeLockKey := arwen.CustomStorageKey(arwen.TimeLockKeyPrefix, key)
	bigTimestamp := big.NewInt(0).SetInt64(lockTimestamp)
	storageStatus, err := storage.SetProtectedStorage(timeLockKey, bigTimestamp.Bytes())
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}
	return int32(storageStatus)
}

//export v1_4_getStorageLock
func v1_4_getStorageLock(context unsafe.Pointer, keyOffset int32, keyLength int32) int64 {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)
	storage := arwen.GetStorageContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.StorageLoad
	metering.UseGas(gasToUse)

	key, err := runtime.MemLoad(keyOffset, keyLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	timeLockKey := arwen.CustomStorageKey(arwen.TimeLockKeyPrefix, key)
	data := storage.GetStorage(timeLockKey)
	timeLock := big.NewInt(0).SetBytes(data).Int64()

	return timeLock
}

//export v1_4_isStorageLocked
func v1_4_isStorageLocked(context unsafe.Pointer, keyOffset int32, keyLength int32) int32 {
	timeLock := v1_4_getStorageLock(context, keyOffset, keyLength)
	if timeLock < 0 {
		return -1
	}

	currentTimestamp := v1_4_getBlockTimestamp(context)
	if timeLock <= currentTimestamp {
		return 0
	}

	return 1
}

//export v1_4_clearStorageLock
func v1_4_clearStorageLock(context unsafe.Pointer, keyOffset int32, keyLength int32) int32 {
	return v1_4_setStorageLock(context, keyOffset, keyLength, 0)
}

//export v1_4_getCaller
func v1_4_getCaller(context unsafe.Pointer, resultOffset int32) {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetCaller
	metering.UseGas(gasToUse)

	caller := runtime.GetVMInput().CallerAddr

	err := runtime.MemStore(resultOffset, caller)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}
}

//export v1_4_checkNoPayment
func v1_4_checkNoPayment(context unsafe.Pointer) {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetCallValue
	metering.UseGas(gasToUse)

	vmInput := runtime.GetVMInput()
	if vmInput.CallValue.Sign() > 0 {
		runtime := arwen.GetRuntimeContext(context)
		arwen.WithFault(arwen.ErrNonPayableFunctionEgld, context, runtime.ElrondAPIErrorShouldFailExecution())
		return
	}
	if len(vmInput.ESDTTransfers) > 0 {
		runtime := arwen.GetRuntimeContext(context)
		arwen.WithFault(arwen.ErrNonPayableFunctionEsdt, context, runtime.ElrondAPIErrorShouldFailExecution())
		return
	}
}

//export v1_4_callValue
func v1_4_callValue(context unsafe.Pointer, resultOffset int32) int32 {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetCallValue
	metering.UseGas(gasToUse)

	value := runtime.GetVMInput().CallValue.Bytes()
	value = arwen.PadBytesLeft(value, arwen.BalanceLen)

	err := runtime.MemStore(resultOffset, value)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	return int32(len(value))
}

//export v1_4_getESDTValue
func v1_4_getESDTValue(context unsafe.Pointer, resultOffset int32) int32 {
	_ = failIfMoreThanOneESDTTransfer(context)
	return v1_4_getESDTValueByIndex(context, resultOffset, 0)
}

//export v1_4_getESDTValueByIndex
func v1_4_getESDTValueByIndex(context unsafe.Pointer, resultOffset int32, index int32) int32 {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetCallValue
	metering.UseGas(gasToUse)

	var value []byte

	esdtTransfer := getESDTTransferFromInput(runtime.GetVMInput(), index)
	if esdtTransfer != nil && esdtTransfer.ESDTValue.Cmp(arwen.Zero) > 0 {
		value = esdtTransfer.ESDTValue.Bytes()
		value = arwen.PadBytesLeft(value, arwen.BalanceLen)
	}

	err := runtime.MemStore(resultOffset, value)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	return int32(len(value))
}

//export v1_4_getESDTTokenName
func v1_4_getESDTTokenName(context unsafe.Pointer, resultOffset int32) int32 {
	_ = failIfMoreThanOneESDTTransfer(context)
	return v1_4_getESDTTokenNameByIndex(context, resultOffset, 0)
}

//export v1_4_getESDTTokenNameByIndex
func v1_4_getESDTTokenNameByIndex(context unsafe.Pointer, resultOffset int32, index int32) int32 {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetCallValue
	metering.UseGas(gasToUse)

	esdtTransfer := getESDTTransferFromInput(runtime.GetVMInput(), index)
	var tokenName []byte
	if esdtTransfer != nil {
		tokenName = esdtTransfer.ESDTTokenName
	}

	err := runtime.MemStore(resultOffset, tokenName)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	return int32(len(tokenName))
}

//export v1_4_getESDTTokenNonce
func v1_4_getESDTTokenNonce(context unsafe.Pointer) int64 {
	_ = failIfMoreThanOneESDTTransfer(context)
	return v1_4_getESDTTokenNonceByIndex(context, 0)
}

//export v1_4_getESDTTokenNonceByIndex
func v1_4_getESDTTokenNonceByIndex(context unsafe.Pointer, index int32) int64 {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetCallValue
	metering.UseGas(gasToUse)

	esdtTransfer := getESDTTransferFromInput(runtime.GetVMInput(), index)
	nonce := uint64(0)
	if esdtTransfer != nil {
		nonce = esdtTransfer.ESDTTokenNonce
	}
	return int64(nonce)
}

//export v1_4_getCurrentESDTNFTNonce
func v1_4_getCurrentESDTNFTNonce(context unsafe.Pointer, addressOffset int32, tokenIDOffset int32, tokenIDLen int32) int64 {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)
	storage := arwen.GetStorageContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.StorageLoad
	metering.UseGas(gasToUse)

	destination, err := runtime.MemLoad(addressOffset, arwen.AddressLen)
	if err != nil {
		return 0
	}

	tokenID, err := runtime.MemLoad(tokenIDOffset, tokenIDLen)
	if err != nil {
		return 0
	}

	key := []byte(core.ElrondProtectedKeyPrefix + core.ESDTNFTLatestNonceIdentifier + string(tokenID))
	data := storage.GetStorageFromAddress(destination, key)

	nonce := big.NewInt(0).SetBytes(data).Uint64()
	return int64(nonce)
}

//export v1_4_getESDTTokenType
func v1_4_getESDTTokenType(context unsafe.Pointer) int32 {
	_ = failIfMoreThanOneESDTTransfer(context)
	return v1_4_getESDTTokenTypeByIndex(context, 0)
}

//export v1_4_getESDTTokenTypeByIndex
func v1_4_getESDTTokenTypeByIndex(context unsafe.Pointer, index int32) int32 {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetCallValue
	metering.UseGas(gasToUse)

	esdtTransfer := getESDTTransferFromInput(runtime.GetVMInput(), index)
	if esdtTransfer != nil {
		return int32(esdtTransfer.ESDTTokenType)
	}
	return 0
}

//export v1_4_getNumESDTTransfers
func v1_4_getNumESDTTransfers(context unsafe.Pointer) int32 {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetCallValue
	metering.UseGas(gasToUse)

	return int32(len(runtime.GetVMInput().ESDTTransfers))
}

//export v1_4_getCallValueTokenName
func v1_4_getCallValueTokenName(context unsafe.Pointer, callValueOffset int32, tokenNameOffset int32) int32 {
	_ = failIfMoreThanOneESDTTransfer(context)
	return v1_4_getCallValueTokenNameByIndex(context, callValueOffset, tokenNameOffset, 0)
}

//export v1_4_getCallValueTokenNameByIndex
func v1_4_getCallValueTokenNameByIndex(context unsafe.Pointer, callValueOffset int32, tokenNameOffset int32, index int32) int32 {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetCallValue
	metering.UseGas(gasToUse)

	callValue := runtime.GetVMInput().CallValue.Bytes()
	tokenName := make([]byte, 0)
	esdtTransfer := getESDTTransferFromInput(runtime.GetVMInput(), index)

	if esdtTransfer != nil {
		tokenName = make([]byte, 0, len(esdtTransfer.ESDTTokenName))
		copy(tokenName, esdtTransfer.ESDTTokenName)
		callValue = esdtTransfer.ESDTValue.Bytes()
	}
	callValue = arwen.PadBytesLeft(callValue, arwen.BalanceLen)

	err := runtime.MemStore(tokenNameOffset, tokenName)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	err = runtime.MemStore(callValueOffset, callValue)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return -1
	}

	return int32(len(tokenName))
}

//export v1_4_writeLog
func v1_4_writeLog(context unsafe.Pointer, dataPointer int32, dataLength int32, topicPtr int32, numTopics int32) {
	// note: deprecated
	runtime := arwen.GetRuntimeContext(context)
	output := arwen.GetOutputContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.Log
	gas := math.MulUint64(metering.GasSchedule().BaseOperationCost.PersistPerByte, uint64(numTopics*arwen.HashLen+dataLength))
	gasToUse = math.AddUint64(gasToUse, gas)
	metering.UseGas(gasToUse)

	log, err := runtime.MemLoad(dataPointer, dataLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	topics, err := arwen.GuardedMakeByteSlice2D(numTopics)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	for i := int32(0); i < numTopics; i++ {
		topics[i], err = runtime.MemLoad(topicPtr+i*arwen.HashLen, arwen.HashLen)
		if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
			return
		}
	}

	output.WriteLog(runtime.GetSCAddress(), topics, log)
}

//export v1_4_writeEventLog
func v1_4_writeEventLog(
	context unsafe.Pointer,
	numTopics int32,
	topicLengthsOffset int32,
	topicOffset int32,
	dataOffset int32,
	dataLength int32,
) {

	host := arwen.GetVMHost(context)
	runtime := arwen.GetRuntimeContext(context)
	output := arwen.GetOutputContext(context)
	metering := arwen.GetMeteringContext(context)

	topics, topicDataTotalLen, err := getArgumentsFromMemory(
		host,
		numTopics,
		topicLengthsOffset,
		topicOffset,
	)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	data, err := runtime.MemLoad(dataOffset, dataLength)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	gasToUse := metering.GasSchedule().ElrondAPICost.Log
	gasForData := math.MulUint64(
		metering.GasSchedule().BaseOperationCost.DataCopyPerByte,
		uint64(topicDataTotalLen+dataLength))
	gasToUse = math.AddUint64(gasToUse, gasForData)
	metering.UseGas(gasToUse)

	output.WriteLog(runtime.GetSCAddress(), topics, data)
}

//export v1_4_getBlockTimestamp
func v1_4_getBlockTimestamp(context unsafe.Pointer) int64 {
	blockchain := arwen.GetBlockchainContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetBlockTimeStamp
	metering.UseGas(gasToUse)

	return int64(blockchain.CurrentTimeStamp())
}

//export v1_4_getBlockNonce
func v1_4_getBlockNonce(context unsafe.Pointer) int64 {
	blockchain := arwen.GetBlockchainContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetBlockNonce
	metering.UseGas(gasToUse)

	return int64(blockchain.CurrentNonce())
}

//export v1_4_getBlockRound
func v1_4_getBlockRound(context unsafe.Pointer) int64 {
	blockchain := arwen.GetBlockchainContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetBlockRound
	metering.UseGas(gasToUse)

	return int64(blockchain.CurrentRound())
}

//export v1_4_getBlockEpoch
func v1_4_getBlockEpoch(context unsafe.Pointer) int64 {
	blockchain := arwen.GetBlockchainContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetBlockEpoch
	metering.UseGas(gasToUse)

	return int64(blockchain.CurrentEpoch())
}

//export v1_4_getBlockRandomSeed
func v1_4_getBlockRandomSeed(context unsafe.Pointer, pointer int32) {
	runtime := arwen.GetRuntimeContext(context)
	blockchain := arwen.GetBlockchainContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetBlockRandomSeed
	metering.UseGas(gasToUse)

	randomSeed := blockchain.CurrentRandomSeed()
	err := runtime.MemStore(pointer, randomSeed)
	arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution())
}

//export v1_4_getStateRootHash
func v1_4_getStateRootHash(context unsafe.Pointer, pointer int32) {
	runtime := arwen.GetRuntimeContext(context)
	blockchain := arwen.GetBlockchainContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetStateRootHash
	metering.UseGas(gasToUse)

	stateRootHash := blockchain.GetStateRootHash()
	err := runtime.MemStore(pointer, stateRootHash)
	arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution())
}

//export v1_4_getPrevBlockTimestamp
func v1_4_getPrevBlockTimestamp(context unsafe.Pointer) int64 {
	blockchain := arwen.GetBlockchainContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetBlockTimeStamp
	metering.UseGas(gasToUse)

	return int64(blockchain.LastTimeStamp())
}

//export v1_4_getPrevBlockNonce
func v1_4_getPrevBlockNonce(context unsafe.Pointer) int64 {
	blockchain := arwen.GetBlockchainContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetBlockNonce
	metering.UseGas(gasToUse)

	return int64(blockchain.LastNonce())
}

//export v1_4_getPrevBlockRound
func v1_4_getPrevBlockRound(context unsafe.Pointer) int64 {
	blockchain := arwen.GetBlockchainContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetBlockRound
	metering.UseGas(gasToUse)

	return int64(blockchain.LastRound())
}

//export v1_4_getPrevBlockEpoch
func v1_4_getPrevBlockEpoch(context unsafe.Pointer) int64 {
	blockchain := arwen.GetBlockchainContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetBlockEpoch
	metering.UseGas(gasToUse)

	return int64(blockchain.LastEpoch())
}

//export v1_4_getPrevBlockRandomSeed
func v1_4_getPrevBlockRandomSeed(context unsafe.Pointer, pointer int32) {
	runtime := arwen.GetRuntimeContext(context)
	blockchain := arwen.GetBlockchainContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetBlockRandomSeed
	metering.UseGas(gasToUse)

	randomSeed := blockchain.LastRandomSeed()
	err := runtime.MemStore(pointer, randomSeed)
	arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution())
}

//export v1_4_returnData
func v1_4_returnData(context unsafe.Pointer, pointer int32, length int32) {
	runtime := arwen.GetRuntimeContext(context)
	output := arwen.GetOutputContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.Finish
	gas := math.MulUint64(metering.GasSchedule().BaseOperationCost.PersistPerByte, uint64(length))
	gasToUse = math.AddUint64(gasToUse, gas)
	metering.UseGas(gasToUse)

	data, err := runtime.MemLoad(pointer, length)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return
	}

	output.Finish(data)
}

//export v1_4_executeOnSameContext
func v1_4_executeOnSameContext(
	context unsafe.Pointer,
	gasLimit int64,
	addressOffset int32,
	valueOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := arwen.GetVMHost(context)
	return ExecuteOnSameContextWithHost(
		host,
		gasLimit,
		addressOffset,
		valueOffset,
		functionOffset,
		functionLength,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)
}

// ExecuteOnSameContextWithHost - executeOnSameContext with host instead of pointer context
func ExecuteOnSameContextWithHost(
	host arwen.VMHost,
	gasLimit int64,
	addressOffset int32,
	valueOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	runtime := host.Runtime()

	callArgs, err := extractIndirectContractCallArgumentsWithValue(
		host, addressOffset, valueOffset, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	return ExecuteOnSameContextWithTypedArgs(
		host,
		gasLimit,
		callArgs.value,
		callArgs.function,
		callArgs.dest,
		callArgs.args,
	)
}

// ExecuteOnSameContextWithTypedArgs - executeOnSameContext with args already read from memory
func ExecuteOnSameContextWithTypedArgs(
	host arwen.VMHost,
	gasLimit int64,
	value *big.Int,
	function []byte,
	dest []byte,
	args [][]byte,
) int32 {
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().ElrondAPICost.ExecuteOnSameContext
	metering.UseGas(gasToUse)

	sender := runtime.GetSCAddress()
	contractCallInput, err := prepareIndirectContractCallInput(
		host,
		sender,
		value,
		gasLimit,
		dest,
		function,
		args,
		gasToUse,
		true,
	)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	if isBuiltInCall(contractCallInput.Function, host) {
		return 1
	}

	_, err = host.ExecuteOnSameContext(contractCallInput)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

//export v1_4_executeOnDestContext
func v1_4_executeOnDestContext(
	context unsafe.Pointer,
	gasLimit int64,
	addressOffset int32,
	valueOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := arwen.GetVMHost(context)
	return ExecuteOnDestContextWithHost(
		host,
		gasLimit,
		addressOffset,
		valueOffset,
		functionOffset,
		functionLength,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)
}

// ExecuteOnDestContextWithHost - executeOnDestContext with host instead of pointer context
func ExecuteOnDestContextWithHost(
	host arwen.VMHost,
	gasLimit int64,
	addressOffset int32,
	valueOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	runtime := host.Runtime()

	callArgs, err := extractIndirectContractCallArgumentsWithValue(
		host, addressOffset, valueOffset, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	return ExecuteOnDestContextWithTypedArgs(
		host,
		gasLimit,
		callArgs.value,
		callArgs.function,
		callArgs.dest,
		callArgs.args,
	)
}

// ExecuteOnDestContextWithTypedArgs - executeOnDestContext with args already read from memory
func ExecuteOnDestContextWithTypedArgs(
	host arwen.VMHost,
	gasLimit int64,
	value *big.Int,
	function []byte,
	dest []byte,
	args [][]byte,
) int32 {
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().ElrondAPICost.ExecuteOnDestContext
	metering.UseGas(gasToUse)

	sender := runtime.GetSCAddress()
	contractCallInput, err := prepareIndirectContractCallInput(
		host,
		sender,
		value,
		gasLimit,
		dest,
		function,
		args,
		gasToUse,
		true,
	)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	_, _, err = host.ExecuteOnDestContext(contractCallInput)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

//export v1_4_executeOnDestContextByCaller
func v1_4_executeOnDestContextByCaller(
	context unsafe.Pointer,
	gasLimit int64,
	addressOffset int32,
	valueOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := arwen.GetVMHost(context)
	return ExecuteOnDestContextByCallerWithHost(
		host,
		gasLimit,
		addressOffset,
		valueOffset,
		functionOffset,
		functionLength,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)
}

// ExecuteOnDestContextByCallerWithHost - executeOnDestContextByCaller with host instead of pointer context
func ExecuteOnDestContextByCallerWithHost(
	host arwen.VMHost,
	gasLimit int64,
	addressOffset int32,
	valueOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	runtime := host.Runtime()

	callArgs, err := extractIndirectContractCallArgumentsWithValue(
		host, addressOffset, valueOffset, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	return ExecuteOnDestContextByCallerWithTypedArgs(
		host,
		gasLimit,
		callArgs.value,
		callArgs.function,
		callArgs.dest,
		callArgs.args,
	)
}

// ExecuteOnDestContextByCallerWithTypedArgs - executeOnDestContextByCaller with args already read from memory
func ExecuteOnDestContextByCallerWithTypedArgs(
	host arwen.VMHost,
	gasLimit int64,
	value *big.Int,
	function []byte,
	dest []byte,
	args [][]byte,
) int32 {
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().ElrondAPICost.ExecuteOnDestContext
	metering.UseGas(gasToUse)

	send := runtime.GetVMInput().CallerAddr
	contractCallInput, err := prepareIndirectContractCallInput(
		host,
		send,
		value,
		gasLimit,
		dest,
		function,
		args,
		gasToUse,
		true,
	)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	if isBuiltInCall(contractCallInput.Function, host) {
		return 1
	}

	_, _, err = host.ExecuteOnDestContext(contractCallInput)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

//export v1_4_delegateExecution
func v1_4_delegateExecution(
	context unsafe.Pointer,
	gasLimit int64,
	addressOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := arwen.GetVMHost(context)
	return DelegateExecutionWithHost(
		host,
		gasLimit,
		addressOffset,
		functionOffset,
		functionLength,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)
}

// DelegateExecutionWithHost - delegateExecution with host instead of pointer context
func DelegateExecutionWithHost(
	host arwen.VMHost,
	gasLimit int64,
	addressOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	runtime := host.Runtime()

	callArgs, err := extractIndirectContractCallArgumentsWithoutValue(
		host, addressOffset, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	return DelegateExecutionWithTypedArgs(
		host,
		gasLimit,
		callArgs.function,
		callArgs.dest,
		callArgs.args,
	)
}

// DelegateExecutionWithTypedArgs - delegateExecution with args already read from memory
func DelegateExecutionWithTypedArgs(
	host arwen.VMHost,
	gasLimit int64,
	function []byte,
	dest []byte,
	args [][]byte,
) int32 {
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().ElrondAPICost.DelegateExecution
	metering.UseGas(gasToUse)

	sender := runtime.GetSCAddress()
	value := runtime.GetVMInput().CallValue
	contractCallInput, err := prepareIndirectContractCallInput(
		host,
		sender,
		value,
		gasLimit,
		dest,
		function,
		args,
		gasToUse,
		true,
	)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	if isBuiltInCall(contractCallInput.Function, host) {
		return 1
	}

	_, err = host.ExecuteOnSameContext(contractCallInput)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

//export v1_4_executeReadOnly
func v1_4_executeReadOnly(
	context unsafe.Pointer,
	gasLimit int64,
	addressOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := arwen.GetVMHost(context)
	return ExecuteReadOnlyWithHost(
		host,
		gasLimit,
		addressOffset,
		functionOffset,
		functionLength,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)
}

// ExecuteReadOnlyWithHost - executeReadOnly with host instead of pointer context
func ExecuteReadOnlyWithHost(
	host arwen.VMHost,
	gasLimit int64,
	addressOffset int32,
	functionOffset int32,
	functionLength int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	runtime := host.Runtime()

	callArgs, err := extractIndirectContractCallArgumentsWithoutValue(
		host, addressOffset, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	return ExecuteReadOnlyWithTypedArguments(
		host,
		gasLimit,
		callArgs.function,
		callArgs.dest,
		callArgs.args,
	)
}

// ExecuteReadOnlyWithTypedArguments - executeReadOnly with args already read from memory
func ExecuteReadOnlyWithTypedArguments(
	host arwen.VMHost,
	gasLimit int64,
	function []byte,
	dest []byte,
	args [][]byte,
) int32 {
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().ElrondAPICost.ExecuteReadOnly
	metering.UseGas(gasToUse)

	sender := runtime.GetSCAddress()
	value := runtime.GetVMInput().CallValue
	contractCallInput, err := prepareIndirectContractCallInput(
		host,
		sender,
		value,
		gasLimit,
		dest,
		function,
		args,
		gasToUse,
		true,
	)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	if isBuiltInCall(contractCallInput.Function, host) {
		return 1
	}

	runtime.SetReadOnly(true)
	_, err = host.ExecuteOnSameContext(contractCallInput)
	runtime.SetReadOnly(false)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

//export v1_4_createContract
func v1_4_createContract(
	context unsafe.Pointer,
	gasLimit int64,
	valueOffset int32,
	codeOffset int32,
	codeMetadataOffset int32,
	length int32,
	resultOffset int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := arwen.GetVMHost(context)
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().ElrondAPICost.CreateContract
	metering.UseGas(gasToUse)

	sender := runtime.GetSCAddress()
	value, err := runtime.MemLoad(valueOffset, arwen.BalanceLen)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	code, err := runtime.MemLoad(codeOffset, length)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	codeMetadata, err := runtime.MemLoad(codeMetadataOffset, arwen.CodeMetadataLen)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	data, actualLen, err := getArgumentsFromMemory(
		host,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(actualLen))
	metering.UseGas(gasToUse)

	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	valueAsInt := big.NewInt(0).SetBytes(value)
	newAddress, err := createContract(sender, data, valueAsInt, metering, gasLimit, code, codeMetadata, host, runtime)

	if err != nil {
		return 1
	}

	err = runtime.MemStore(resultOffset, newAddress)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

//export v1_4_deployFromSourceContract
func v1_4_deployFromSourceContract(
	context unsafe.Pointer,
	gasLimit int64,
	valueOffset int32,
	sourceContractAddressOffset int32,
	codeMetadataOffset int32,
	resultAddressOffset int32,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) int32 {
	host := arwen.GetVMHost(context)
	runtime := host.Runtime()
	metering := host.Metering()

	gasToUse := metering.GasSchedule().ElrondAPICost.CreateContract
	metering.UseGas(gasToUse)

	value, err := runtime.MemLoad(valueOffset, arwen.BalanceLen)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	sourceContractAddress, err := runtime.MemLoad(sourceContractAddressOffset, arwen.AddressLen)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	codeMetadata, err := runtime.MemLoad(codeMetadataOffset, arwen.CodeMetadataLen)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	data, actualLen, err := getArgumentsFromMemory(
		host,
		numArguments,
		argumentsLengthOffset,
		dataOffset,
	)

	gasToUse = math.MulUint64(metering.GasSchedule().BaseOperationCost.DataCopyPerByte, uint64(actualLen))
	metering.UseGas(gasToUse)

	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	newAddress, err := DeployFromSourceContractWithTypedArgs(
		host,
		sourceContractAddress,
		codeMetadata,
		big.NewInt(0).SetBytes(value),
		data,
		gasLimit,
	)

	if err != nil {
		return 1
	}

	err = runtime.MemStore(resultAddressOffset, newAddress)
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 1
	}

	return 0
}

// DeployFromSourceContractWithTypedArgs - deployFromSourceContract with args already read from memory
func DeployFromSourceContractWithTypedArgs(
	host arwen.VMHost,
	sourceContractAddress []byte,
	codeMetadata []byte,
	value *big.Int,
	data [][]byte,
	gasLimit int64,
) ([]byte, error) {
	runtime := host.Runtime()
	metering := host.Metering()
	sender := runtime.GetSCAddress()

	blockchain := host.Blockchain()
	code, err := blockchain.GetCode(sourceContractAddress)
	if arwen.WithFaultAndHost(host, err, runtime.ElrondAPIErrorShouldFailExecution()) {
		return nil, err
	}

	return createContract(sender, data, value, metering, gasLimit, code, codeMetadata, host, runtime)
}

func createContract(
	sender []byte,
	data [][]byte,
	value *big.Int,
	metering arwen.MeteringContext,
	gasLimit int64,
	code []byte,
	codeMetadata []byte,
	host arwen.VMHost,
	_ arwen.RuntimeContext,
) ([]byte, error) {
	contractCreate := &vmcommon.ContractCreateInput{
		VMInput: vmcommon.VMInput{
			CallerAddr:  sender,
			Arguments:   data,
			CallValue:   value,
			GasPrice:    0,
			GasProvided: metering.BoundGasLimit(gasLimit),
		},
		ContractCode:         code,
		ContractCodeMetadata: codeMetadata,
	}

	return host.CreateNewContract(contractCreate)
}

//export v1_4_getNumReturnData
func v1_4_getNumReturnData(context unsafe.Pointer) int32 {
	output := arwen.GetOutputContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetNumReturnData
	metering.UseGas(gasToUse)

	returnData := output.ReturnData()
	return int32(len(returnData))
}

//export v1_4_getReturnDataSize
func v1_4_getReturnDataSize(context unsafe.Pointer, resultID int32) int32 {
	output := arwen.GetOutputContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetReturnDataSize
	metering.UseGas(gasToUse)

	returnData := output.ReturnData()
	if resultID >= int32(len(returnData)) {
		return 0
	}

	return int32(len(returnData[resultID]))
}

//export v1_4_getReturnData
func v1_4_getReturnData(context unsafe.Pointer, resultID int32, dataOffset int32) int32 {
	runtime := arwen.GetRuntimeContext(context)
	output := arwen.GetOutputContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetReturnData
	metering.UseGas(gasToUse)

	returnData := output.ReturnData()
	if resultID >= int32(len(returnData)) {
		return 0
	}

	err := runtime.MemStore(dataOffset, returnData[resultID])
	if arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution()) {
		return 0
	}

	return int32(len(returnData[resultID]))
}

//export v1_4_getOriginalTxHash
func v1_4_getOriginalTxHash(context unsafe.Pointer, dataOffset int32) {
	runtime := arwen.GetRuntimeContext(context)
	metering := arwen.GetMeteringContext(context)

	gasToUse := metering.GasSchedule().ElrondAPICost.GetBlockHash
	metering.UseGas(gasToUse)

	err := runtime.MemStore(dataOffset, runtime.GetOriginalTxHash())
	_ = arwen.WithFault(err, context, runtime.ElrondAPIErrorShouldFailExecution())
}

func prepareIndirectContractCallInput(
	host arwen.VMHost,
	sender []byte,
	value *big.Int,
	gasLimit int64,
	destination []byte,
	function []byte,
	data [][]byte,
	_ uint64,
	syncExecutionRequired bool,
) (*vmcommon.ContractCallInput, error) {
	runtime := host.Runtime()
	metering := host.Metering()

	if syncExecutionRequired && !host.AreInSameShard(runtime.GetSCAddress(), destination) {
		return nil, arwen.ErrSyncExecutionNotInSameShard
	}

	contractCallInput := &vmcommon.ContractCallInput{
		VMInput: vmcommon.VMInput{
			CallerAddr:  sender,
			Arguments:   data,
			CallValue:   value,
			GasPrice:    0,
			GasProvided: metering.BoundGasLimit(gasLimit),
		},
		RecipientAddr: destination,
		Function:      string(function),
	}

	return contractCallInput, nil
}

func getArgumentsFromMemory(
	host arwen.VMHost,
	numArguments int32,
	argumentsLengthOffset int32,
	dataOffset int32,
) ([][]byte, int32, error) {
	runtime := host.Runtime()

	if numArguments < 0 {
		return nil, 0, fmt.Errorf("negative numArguments (%d)", numArguments)
	}

	argumentsLengthData, err := runtime.MemLoad(argumentsLengthOffset, numArguments*4)
	if err != nil {
		return nil, 0, err
	}

	argumentLengths := createInt32Array(argumentsLengthData, numArguments)
	data, err := runtime.MemLoadMultiple(dataOffset, argumentLengths)
	if err != nil {
		return nil, 0, err
	}

	totalArgumentBytes := int32(0)
	for _, length := range argumentLengths {
		totalArgumentBytes += length
	}

	return data, totalArgumentBytes, nil
}

func createInt32Array(rawData []byte, numIntegers int32) []int32 {
	integers := make([]int32, numIntegers)
	index := 0
	for cursor := 0; cursor < len(rawData); cursor += 4 {
		rawInt := rawData[cursor : cursor+4]
		actualInt := binary.LittleEndian.Uint32(rawInt)
		integers[index] = int32(actualInt)
		index++
	}
	return integers
}
