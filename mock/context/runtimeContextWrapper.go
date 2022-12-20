package mock

import (
	vmcommon "github.com/ElrondNetwork/elrond-vm-common"
	"github.com/ElrondNetwork/wasm-vm-v1_4/arwen"
	"github.com/ElrondNetwork/wasm-vm-v1_4/wasmer"
)

// making sure we implement all functions of RuntimeContext
var _ arwen.RuntimeContext = (*RuntimeContextWrapper)(nil)

// RuntimeContextWrapper a wrapper over a RuntimeContext that delegates to if if function is not redefined
type RuntimeContextWrapper struct {
	runtimeContext arwen.RuntimeContext

	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	InitStateFromContractCallInputFunc func(input *vmcommon.ContractCallInput)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	SetCustomCallFunctionFunc func(callFunction string)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	GetVMInputFunc func() *vmcommon.ContractCallInput
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	SetVMInputFunc func(vmInput *vmcommon.ContractCallInput)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	GetSCAddressFunc func() []byte
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	SetCodeAddressFunc func(scAddress []byte)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	GetSCCodeFunc func() ([]byte, error)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	GetSCCodeSizeFunc func() uint64
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	GetVMTypeFunc func() []byte
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	FunctionFunc func() string
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	ArgumentsFunc func() [][]byte
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	GetCurrentTxHashFunc func() []byte
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	GetOriginalTxHashFunc func() []byte
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	ExtractCodeUpgradeFromArgsFunc func() ([]byte, []byte, error)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	SignalUserErrorFunc func(message string)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	FailExecutionFunc func(err error)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	MustVerifyNextContractCodeFunc func()
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	SetRuntimeBreakpointValueFunc func(value arwen.BreakpointValue)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	GetRuntimeBreakpointValueFunc func() arwen.BreakpointValue
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	GetAsyncCallInfoFunc func() *arwen.AsyncCallInfo
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	SetAsyncCallInfoFunc func(asyncCallInfo *arwen.AsyncCallInfo)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	AddAsyncContextCallFunc func(contextIdentifier []byte, asyncCall *arwen.AsyncGeneratedCall) error
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	GetAsyncContextInfoFunc func() *arwen.AsyncContextInfo
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	GetAsyncContextFunc func(contextIdentifier []byte) (*arwen.AsyncContext, error)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	RunningInstancesCountFunc func() uint64
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	IsFunctionImportedFunc func(name string) bool
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	ReadOnlyFunc func() bool
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	SetReadOnlyFunc func(readOnly bool)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	StartWasmerInstanceFunc func(contract []byte, gasLimit uint64, newCode bool) error
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	ClearWarmInstanceCacheFunc func()
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	SetMaxInstanceCountFunc func(maxInstances uint64)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	VerifyContractCodeFunc func() error
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	GetInstanceFunc func() wasmer.InstanceHandler
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	GetInstanceExportsFunc func() wasmer.ExportsMap
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	GetInitFunctionFunc func() wasmer.ExportedFunctionCallback
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	GetFunctionToCallFunc func() (wasmer.ExportedFunctionCallback, error)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	GetPointsUsedFunc func() uint64
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	SetPointsUsedFunc func(gasPoints uint64)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	MemStoreFunc func(offset int32, data []byte) error
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	MemLoadFunc func(offset int32, length int32) ([]byte, error)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	MemLoadMultipleFunc func(offset int32, lengths []int32) ([][]byte, error)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	ElrondAPIErrorShouldFailExecutionFunc func() bool
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	ElrondSyncExecAPIErrorShouldFailExecutionFunc func() bool
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	CryptoAPIErrorShouldFailExecutionFunc func() bool
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	BigIntAPIErrorShouldFailExecutionFunc func() bool
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	BigFloatAPIErrorShouldFailExecutionFunc func() bool
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	ManagedBufferAPIErrorShouldFailExecutionFunc func() bool
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	ExecuteAsyncCallFunc func(address []byte, data []byte, value []byte) error
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	ReplaceInstanceBuilderFunc func(builder arwen.InstanceBuilder)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	AddErrorFunc func(err error, otherInfo ...string)
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	GetAllErrorsFunc func() error
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	InitStateFunc func()
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	PushStateFunc func()
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	PopSetActiveStateFunc func()
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	PopDiscardFunc func()
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	ClearStateStackFunc func()
	// function that will be called by the corresponding RuntimeContext function implementation (by default this will call the same wrapped context function)
	CleanInstanceFunc func()
}

// NewRuntimeContextWrapper builds a new runtimeContextWrapper that by default will delagate all calls to the provided RuntimeContext
func NewRuntimeContextWrapper(inputRuntimeContext *arwen.RuntimeContext) *RuntimeContextWrapper {

	runtimeWrapper := &RuntimeContextWrapper{runtimeContext: *inputRuntimeContext}

	/*
		default implementations delegate to wrapped context
	*/

	runtimeWrapper.InitStateFromContractCallInputFunc = func(input *vmcommon.ContractCallInput) {
		runtimeWrapper.runtimeContext.InitStateFromContractCallInput(input)
	}

	runtimeWrapper.SetCustomCallFunctionFunc = func(callFunction string) {
		runtimeWrapper.runtimeContext.SetCustomCallFunction(callFunction)
	}

	runtimeWrapper.GetVMInputFunc = func() *vmcommon.ContractCallInput {
		return runtimeWrapper.runtimeContext.GetVMInput()
	}

	runtimeWrapper.SetVMInputFunc = func(vmInput *vmcommon.ContractCallInput) {
		runtimeWrapper.runtimeContext.SetVMInput(vmInput)
	}

	runtimeWrapper.GetSCAddressFunc = func() []byte {
		return runtimeWrapper.runtimeContext.GetContextAddress()
	}

	runtimeWrapper.SetCodeAddressFunc = func(scAddress []byte) {
		runtimeWrapper.runtimeContext.SetCodeAddress(scAddress)
	}

	runtimeWrapper.GetSCCodeFunc = func() ([]byte, error) {
		return runtimeWrapper.runtimeContext.GetSCCode()
	}

	runtimeWrapper.GetSCCodeSizeFunc = func() uint64 {
		return runtimeWrapper.runtimeContext.GetSCCodeSize()
	}

	runtimeWrapper.GetVMTypeFunc = func() []byte {
		return runtimeWrapper.runtimeContext.GetVMType()
	}

	runtimeWrapper.FunctionFunc = func() string {
		return runtimeWrapper.runtimeContext.Function()
	}

	runtimeWrapper.ArgumentsFunc = func() [][]byte {
		return runtimeWrapper.runtimeContext.Arguments()
	}

	runtimeWrapper.GetCurrentTxHashFunc = func() []byte {
		return runtimeWrapper.runtimeContext.GetCurrentTxHash()
	}

	runtimeWrapper.GetOriginalTxHashFunc = func() []byte {
		return runtimeWrapper.runtimeContext.GetOriginalTxHash()
	}

	runtimeWrapper.ExtractCodeUpgradeFromArgsFunc = func() ([]byte, []byte, error) {
		return runtimeWrapper.runtimeContext.ExtractCodeUpgradeFromArgs()
	}

	runtimeWrapper.SignalUserErrorFunc = func(message string) {
		runtimeWrapper.runtimeContext.SignalUserError(message)
	}

	runtimeWrapper.FailExecutionFunc = func(err error) {
		runtimeWrapper.runtimeContext.FailExecution(err)
	}

	runtimeWrapper.MustVerifyNextContractCodeFunc = func() {
		runtimeWrapper.runtimeContext.MustVerifyNextContractCode()
	}

	runtimeWrapper.SetRuntimeBreakpointValueFunc = func(value arwen.BreakpointValue) {
		runtimeWrapper.runtimeContext.SetRuntimeBreakpointValue(value)
	}

	runtimeWrapper.GetRuntimeBreakpointValueFunc = func() arwen.BreakpointValue {
		return runtimeWrapper.runtimeContext.GetRuntimeBreakpointValue()
	}

	runtimeWrapper.GetAsyncCallInfoFunc = func() *arwen.AsyncCallInfo {
		return runtimeWrapper.runtimeContext.GetAsyncCallInfo()
	}

	runtimeWrapper.SetAsyncCallInfoFunc = func(asyncCallInfo *arwen.AsyncCallInfo) {
		runtimeWrapper.runtimeContext.SetAsyncCallInfo(asyncCallInfo)
	}

	runtimeWrapper.AddAsyncContextCallFunc = func(contextIdentifier []byte, asyncCall *arwen.AsyncGeneratedCall) error {
		return runtimeWrapper.runtimeContext.AddAsyncContextCall(contextIdentifier, asyncCall)
	}

	runtimeWrapper.GetAsyncContextInfoFunc = func() *arwen.AsyncContextInfo {
		return runtimeWrapper.runtimeContext.GetAsyncContextInfo()
	}

	runtimeWrapper.GetAsyncContextFunc = func(contextIdentifier []byte) (*arwen.AsyncContext, error) {
		return runtimeWrapper.runtimeContext.GetAsyncContext(contextIdentifier)
	}

	runtimeWrapper.RunningInstancesCountFunc = func() uint64 {
		return runtimeWrapper.runtimeContext.RunningInstancesCount()
	}

	runtimeWrapper.IsFunctionImportedFunc = func(name string) bool {
		return runtimeWrapper.runtimeContext.IsFunctionImported(name)
	}

	runtimeWrapper.ReadOnlyFunc = func() bool {
		return runtimeWrapper.runtimeContext.ReadOnly()
	}

	runtimeWrapper.SetReadOnlyFunc = func(readOnly bool) {
		runtimeWrapper.runtimeContext.SetReadOnly(readOnly)
	}

	runtimeWrapper.StartWasmerInstanceFunc = func(contract []byte, gasLimit uint64, newCode bool) error {
		return runtimeWrapper.runtimeContext.StartWasmerInstance(contract, gasLimit, newCode)
	}

	runtimeWrapper.ClearWarmInstanceCacheFunc = func() {
		runtimeWrapper.runtimeContext.ClearWarmInstanceCache()
	}

	runtimeWrapper.SetMaxInstanceCountFunc = func(maxInstances uint64) {
		runtimeWrapper.runtimeContext.SetMaxInstanceCount(maxInstances)
	}

	runtimeWrapper.VerifyContractCodeFunc = func() error {
		return runtimeWrapper.runtimeContext.VerifyContractCode()
	}

	runtimeWrapper.GetInstanceFunc = func() wasmer.InstanceHandler {
		return runtimeWrapper.runtimeContext.GetInstance()
	}

	runtimeWrapper.GetInstanceExportsFunc = func() wasmer.ExportsMap {
		return runtimeWrapper.runtimeContext.GetInstanceExports()
	}

	runtimeWrapper.GetInitFunctionFunc = func() wasmer.ExportedFunctionCallback {
		return runtimeWrapper.runtimeContext.GetInitFunction()
	}

	runtimeWrapper.GetFunctionToCallFunc = func() (wasmer.ExportedFunctionCallback, error) {
		return runtimeWrapper.runtimeContext.GetFunctionToCall()
	}

	runtimeWrapper.GetPointsUsedFunc = func() uint64 {
		return runtimeWrapper.runtimeContext.GetPointsUsed()
	}

	runtimeWrapper.SetPointsUsedFunc = func(gasPoints uint64) {
		runtimeWrapper.runtimeContext.SetPointsUsed(gasPoints)
	}

	runtimeWrapper.MemStoreFunc = func(offset int32, data []byte) error {
		return runtimeWrapper.runtimeContext.MemStore(offset, data)
	}

	runtimeWrapper.MemLoadFunc = func(offset int32, length int32) ([]byte, error) {
		return runtimeWrapper.runtimeContext.MemLoad(offset, length)
	}

	runtimeWrapper.MemLoadMultipleFunc = func(offset int32, lengths []int32) ([][]byte, error) {
		return runtimeWrapper.runtimeContext.MemLoadMultiple(offset, lengths)
	}

	runtimeWrapper.ElrondAPIErrorShouldFailExecutionFunc = func() bool {
		return runtimeWrapper.runtimeContext.ElrondAPIErrorShouldFailExecution()
	}

	runtimeWrapper.ElrondSyncExecAPIErrorShouldFailExecutionFunc = func() bool {
		return runtimeWrapper.runtimeContext.ElrondSyncExecAPIErrorShouldFailExecution()
	}

	runtimeWrapper.CryptoAPIErrorShouldFailExecutionFunc = func() bool {
		return runtimeWrapper.runtimeContext.CryptoAPIErrorShouldFailExecution()
	}

	runtimeWrapper.BigIntAPIErrorShouldFailExecutionFunc = func() bool {
		return runtimeWrapper.runtimeContext.BigIntAPIErrorShouldFailExecution()
	}

	runtimeWrapper.BigFloatAPIErrorShouldFailExecutionFunc = func() bool {
		return runtimeWrapper.runtimeContext.BigFloatAPIErrorShouldFailExecution()
	}

	runtimeWrapper.ManagedBufferAPIErrorShouldFailExecutionFunc = func() bool {
		return runtimeWrapper.runtimeContext.ManagedBufferAPIErrorShouldFailExecution()
	}

	runtimeWrapper.ExecuteAsyncCallFunc = func(address []byte, data []byte, value []byte) error {
		return runtimeWrapper.runtimeContext.ExecuteAsyncCall(address, data, value)
	}

	runtimeWrapper.ReplaceInstanceBuilderFunc = func(builder arwen.InstanceBuilder) {
		runtimeWrapper.runtimeContext.ReplaceInstanceBuilder(builder)
	}

	runtimeWrapper.AddErrorFunc = func(err error, otherInfo ...string) {
		runtimeWrapper.runtimeContext.AddError(err, otherInfo...)
	}

	runtimeWrapper.GetAllErrorsFunc = func() error {
		return runtimeWrapper.runtimeContext.GetAllErrors()
	}

	runtimeWrapper.InitStateFunc = func() {
		runtimeWrapper.runtimeContext.InitState()
	}

	runtimeWrapper.PushStateFunc = func() {
		runtimeWrapper.runtimeContext.PushState()
	}

	runtimeWrapper.PopSetActiveStateFunc = func() {
		runtimeWrapper.runtimeContext.PopSetActiveState()
	}

	runtimeWrapper.PopDiscardFunc = func() {
		runtimeWrapper.runtimeContext.PopDiscard()
	}

	runtimeWrapper.ClearStateStackFunc = func() {
		runtimeWrapper.runtimeContext.ClearStateStack()
	}

	runtimeWrapper.CleanInstanceFunc = func() {
		runtimeWrapper.runtimeContext.CleanInstance()
	}

	return runtimeWrapper
}

// GetWrappedRuntimeContext gets the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) GetWrappedRuntimeContext() arwen.RuntimeContext {
	return contextWrapper.runtimeContext
}

// InitStateFromContractCallInput calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) InitStateFromContractCallInput(input *vmcommon.ContractCallInput) {
	contextWrapper.InitStateFromContractCallInputFunc(input)
}

// SetCustomCallFunction calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) SetCustomCallFunction(callFunction string) {
	contextWrapper.SetCustomCallFunctionFunc(callFunction)
}

// GetVMInput calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) GetVMInput() *vmcommon.ContractCallInput {
	return contextWrapper.GetVMInputFunc()
}

// SetVMInput calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) SetVMInput(vmInput *vmcommon.ContractCallInput) {
	contextWrapper.SetVMInputFunc(vmInput)
}

// GetContextAddress calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) GetContextAddress() []byte {
	return contextWrapper.GetSCAddressFunc()
}

// SetCodeAddress calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) SetCodeAddress(scAddress []byte) {
	contextWrapper.SetCodeAddressFunc(scAddress)
}

// GetSCCode calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) GetSCCode() ([]byte, error) {
	return contextWrapper.GetSCCodeFunc()
}

// GetSCCodeSize calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) GetSCCodeSize() uint64 {
	return contextWrapper.GetSCCodeSizeFunc()
}

// GetVMType calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) GetVMType() []byte {
	return contextWrapper.GetVMTypeFunc()
}

// Function calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) Function() string {
	return contextWrapper.FunctionFunc()
}

// Arguments calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) Arguments() [][]byte {
	return contextWrapper.ArgumentsFunc()
}

// GetCurrentTxHash calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) GetCurrentTxHash() []byte {
	return contextWrapper.GetCurrentTxHashFunc()
}

// GetOriginalTxHash calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) GetOriginalTxHash() []byte {
	return contextWrapper.GetOriginalTxHashFunc()
}

// ExtractCodeUpgradeFromArgs calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) ExtractCodeUpgradeFromArgs() ([]byte, []byte, error) {
	return contextWrapper.ExtractCodeUpgradeFromArgsFunc()
}

// SignalUserError calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) SignalUserError(message string) {
	contextWrapper.SignalUserErrorFunc(message)
}

// FailExecution calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) FailExecution(err error) {
	contextWrapper.FailExecutionFunc(err)
}

// MustVerifyNextContractCode calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) MustVerifyNextContractCode() {
	contextWrapper.MustVerifyNextContractCodeFunc()
}

// SetRuntimeBreakpointValue calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) SetRuntimeBreakpointValue(value arwen.BreakpointValue) {
	contextWrapper.SetRuntimeBreakpointValueFunc(value)
}

// GetRuntimeBreakpointValue calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) GetRuntimeBreakpointValue() arwen.BreakpointValue {
	return contextWrapper.GetRuntimeBreakpointValueFunc()
}

// GetAsyncCallInfo calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) GetAsyncCallInfo() *arwen.AsyncCallInfo {
	return contextWrapper.GetAsyncCallInfoFunc()
}

// SetAsyncCallInfo calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) SetAsyncCallInfo(asyncCallInfo *arwen.AsyncCallInfo) {
	contextWrapper.SetAsyncCallInfoFunc(asyncCallInfo)
}

// AddAsyncContextCall calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) AddAsyncContextCall(contextIdentifier []byte, asyncCall *arwen.AsyncGeneratedCall) error {
	return contextWrapper.AddAsyncContextCallFunc(contextIdentifier, asyncCall)
}

// GetAsyncContextInfo calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) GetAsyncContextInfo() *arwen.AsyncContextInfo {
	return contextWrapper.GetAsyncContextInfoFunc()
}

// GetAsyncContext calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) GetAsyncContext(contextIdentifier []byte) (*arwen.AsyncContext, error) {
	return contextWrapper.GetAsyncContextFunc(contextIdentifier)
}

// RunningInstancesCount calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) RunningInstancesCount() uint64 {
	return contextWrapper.RunningInstancesCountFunc()
}

// IsFunctionImported calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) IsFunctionImported(name string) bool {
	return contextWrapper.IsFunctionImportedFunc(name)
}

// ReadOnly calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) ReadOnly() bool {
	return contextWrapper.ReadOnlyFunc()
}

// SetReadOnly calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) SetReadOnly(readOnly bool) {
	contextWrapper.SetReadOnlyFunc(readOnly)
}

// StartWasmerInstance calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) StartWasmerInstance(contract []byte, gasLimit uint64, newCode bool) error {
	return contextWrapper.StartWasmerInstanceFunc(contract, gasLimit, newCode)
}

// ClearWarmInstanceCache calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) ClearWarmInstanceCache() {
	contextWrapper.ClearWarmInstanceCacheFunc()
}

// SetMaxInstanceCount calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) SetMaxInstanceCount(maxInstances uint64) {
	contextWrapper.SetMaxInstanceCountFunc(maxInstances)
}

// VerifyContractCode calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) VerifyContractCode() error {
	return contextWrapper.VerifyContractCodeFunc()
}

// GetInstance calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) GetInstance() wasmer.InstanceHandler {
	return contextWrapper.GetInstanceFunc()
}

// GetInstanceExports calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) GetInstanceExports() wasmer.ExportsMap {
	return contextWrapper.GetInstanceExportsFunc()
}

// GetInitFunction calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) GetInitFunction() wasmer.ExportedFunctionCallback {
	return contextWrapper.GetInitFunctionFunc()
}

// GetFunctionToCall calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) GetFunctionToCall() (wasmer.ExportedFunctionCallback, error) {
	return contextWrapper.GetFunctionToCallFunc()
}

// GetPointsUsed calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) GetPointsUsed() uint64 {
	return contextWrapper.GetPointsUsedFunc()
}

// SetPointsUsed calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) SetPointsUsed(gasPoints uint64) {
	contextWrapper.SetPointsUsedFunc(gasPoints)
}

// MemStore calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) MemStore(offset int32, data []byte) error {
	return contextWrapper.MemStoreFunc(offset, data)
}

// MemLoad calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) MemLoad(offset int32, length int32) ([]byte, error) {
	return contextWrapper.MemLoadFunc(offset, length)
}

// MemLoadMultiple calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) MemLoadMultiple(offset int32, lengths []int32) ([][]byte, error) {
	return contextWrapper.MemLoadMultipleFunc(offset, lengths)
}

// ElrondAPIErrorShouldFailExecution calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) ElrondAPIErrorShouldFailExecution() bool {
	return contextWrapper.ElrondAPIErrorShouldFailExecutionFunc()
}

// ElrondSyncExecAPIErrorShouldFailExecution calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) ElrondSyncExecAPIErrorShouldFailExecution() bool {
	return contextWrapper.ElrondSyncExecAPIErrorShouldFailExecutionFunc()
}

// CryptoAPIErrorShouldFailExecution calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) CryptoAPIErrorShouldFailExecution() bool {
	return contextWrapper.CryptoAPIErrorShouldFailExecutionFunc()
}

// BigIntAPIErrorShouldFailExecution calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) BigIntAPIErrorShouldFailExecution() bool {
	return contextWrapper.BigIntAPIErrorShouldFailExecutionFunc()
}

// BigIntAPIErrorShouldFailExecution calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) BigFloatAPIErrorShouldFailExecution() bool {
	return contextWrapper.BigFloatAPIErrorShouldFailExecutionFunc()
}

// ManagedBufferAPIErrorShouldFailExecution calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) ManagedBufferAPIErrorShouldFailExecution() bool {
	return contextWrapper.ManagedBufferAPIErrorShouldFailExecution()
}

// ExecuteAsyncCall calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) ExecuteAsyncCall(address []byte, data []byte, value []byte) error {
	return contextWrapper.ExecuteAsyncCallFunc(address, data, value)
}

// ReplaceInstanceBuilder calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) ReplaceInstanceBuilder(builder arwen.InstanceBuilder) {
	contextWrapper.ReplaceInstanceBuilderFunc(builder)
}

// AddError calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) AddError(err error, otherInfo ...string) {
	contextWrapper.AddErrorFunc(err, otherInfo...)
}

// GetAllErrors calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) GetAllErrors() error {
	return contextWrapper.GetAllErrorsFunc()
}

// InitState calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) InitState() {
	contextWrapper.InitStateFunc()
}

// PushState calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) PushState() {
	contextWrapper.PushStateFunc()
}

// PopSetActiveState calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) PopSetActiveState() {
	contextWrapper.PopSetActiveStateFunc()
}

// PopDiscard calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) PopDiscard() {
	contextWrapper.PopDiscardFunc()
}

// ClearStateStack calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) ClearStateStack() {
	contextWrapper.ClearStateStackFunc()
}

// CleanInstance calls corresponding xxxFunc function, that by default in turn calls the original method of the wrapped RuntimeContext
func (contextWrapper *RuntimeContextWrapper) CleanInstance() {
	contextWrapper.CleanInstanceFunc()
}

// NumRunningInstances -
func (contextWrapper *RuntimeContextWrapper) NumRunningInstances() (int, int) {
	return 0, 0
}
