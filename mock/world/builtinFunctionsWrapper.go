package worldmock

import (
	"bytes"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/marshal"
	vmcommon "github.com/multiversx/mx-chain-vm-common-go"
	"github.com/multiversx/mx-chain-vm-common-go/builtInFunctions"
	"github.com/multiversx/mx-chain-vm-v1_4-go/config"
	"github.com/multiversx/mx-chain-vm-v1_4-go/vmhost/mock"
)

// WorldMarshalizer is the global marshalizer to be used by the components of
// the BuiltinFunctionsWrapper.
var WorldMarshalizer = &marshal.GogoProtoMarshalizer{}

// BuiltinFunctionsWrapper manages and initializes a BuiltInFunctionContainer
// along with its dependencies
type BuiltinFunctionsWrapper struct {
	Container       vmcommon.BuiltInFunctionContainer
	MapDNSAddresses map[string]struct{}
	World           *MockWorld
	Marshalizer     vmcommon.Marshalizer
}

// NewBuiltinFunctionsWrapper creates a new BuiltinFunctionsWrapper with
// default dependencies.
func NewBuiltinFunctionsWrapper(
	world *MockWorld,
	gasMap config.GasScheduleMap,
) (*BuiltinFunctionsWrapper, error) {

	dnsMap := makeDNSAddresses(numDNSAddresses)

	argsBuiltIn := builtInFunctions.ArgsCreateBuiltInFunctionContainer{
		GasMap:                           gasMap,
		MapDNSAddresses:                  dnsMap,
		Marshalizer:                      WorldMarshalizer,
		Accounts:                         world.AccountsAdapter,
		GuardedAccountHandler:            world.GuardedAccountHandler,
		ShardCoordinator:                 world,
		MaxNumOfAddressesForTransferRole: 100,
		EnableEpochsHandler: &mock.EnableEpochsHandlerStub{
			IsStorageAPICostOptimizationFlagEnabledField:         true,
			IsMultiESDTTransferFixOnCallBackFlagEnabledField:     true,
			IsFixOOGReturnCodeFlagEnabledField:                   true,
			IsRemoveNonUpdatedStorageFlagEnabledField:            true,
			IsCreateNFTThroughExecByCallerFlagEnabledField:       true,
			IsManagedCryptoAPIsFlagEnabledField:                  true,
			IsFailExecutionOnEveryAPIErrorFlagEnabledField:       true,
			IsRefactorContextFlagEnabledField:                    true,
			IsCheckCorrectTokenIDForTransferRoleFlagEnabledField: true,
			IsDisableExecByCallerFlagEnabledField:                true,
			IsESDTTransferRoleFlagEnabledField:                   true,
			IsGlobalMintBurnFlagEnabledField:                     true,
			IsTransferToMetaFlagEnabledField:                     true,
			IsCheckFrozenCollectionFlagEnabledField:              true,
			IsFixAsyncCallbackCheckFlagEnabledField:              true,
			IsESDTNFTImprovementV1FlagEnabledField:               true,
			IsSaveToSystemAccountFlagEnabledField:                true,
			IsValueLengthCheckFlagEnabledField:                   true,
			IsSCDeployFlagEnabledField:                           true,
			IsRepairCallbackFlagEnabledField:                     true,
			IsAheadOfTimeGasUsageFlagEnabledField:                true,
			IsCheckFunctionArgumentFlagEnabledField:              true,
			IsCheckExecuteOnReadOnlyFlagEnabledField:             true,
			IsFixOldTokenLiquidityEnabledField:                   true,
		},
	}

	builtinFuncFactory, err := builtInFunctions.NewBuiltInFunctionsCreator(argsBuiltIn)
	if err != nil {
		return nil, err
	}

	err = builtinFuncFactory.CreateBuiltInFunctionContainer()
	if err != nil {
		return nil, err
	}

	err = builtinFuncFactory.SetPayableHandler(world)
	if err != nil {
		return nil, err
	}

	builtinFuncsWrapper := &BuiltinFunctionsWrapper{
		Container:       builtinFuncFactory.BuiltInFunctionContainer(),
		MapDNSAddresses: argsBuiltIn.MapDNSAddresses,
		World:           world,
	}

	return builtinFuncsWrapper, nil
}

// ProcessBuiltInFunction delegates the execution of a real builtin function to
// the inner BuiltInFunctionContainer.
func (bf *BuiltinFunctionsWrapper) ProcessBuiltInFunction(input *vmcommon.ContractCallInput) (*vmcommon.VMOutput, error) {
	caller := bf.getAccountSharded(input.CallerAddr)
	recipient := bf.getAccountSharded(input.RecipientAddr)

	function, err := bf.Container.Get(input.Function)
	if err != nil {
		return nil, err
	}

	vmOutput, err := function.ProcessBuiltinFunction(caller, recipient, input)
	if err != nil {
		return nil, err
	}

	if !check.IfNil(caller) {
		err = bf.World.AccountsAdapter.SaveAccount(caller)
		if err != nil {
			return nil, err
		}
	}

	if !check.IfNil(recipient) && !bytes.Equal(input.CallerAddr, input.RecipientAddr) {
		err = bf.World.AccountsAdapter.SaveAccount(recipient)
		if err != nil {
			return nil, err
		}
	}

	return vmOutput, nil
}

// GetBuiltinFunctionNames returns the list of defined builtin-in functions.
func (bf *BuiltinFunctionsWrapper) GetBuiltinFunctionNames() vmcommon.FunctionNames {
	return bf.Container.Keys()
}

// TODO change AccountMap to support this instead
func (bf *BuiltinFunctionsWrapper) getAccountSharded(address []byte) vmcommon.UserAccountHandler {
	accountShard := bf.World.ComputeId(address)
	if accountShard != bf.World.SelfId() {
		return nil
	}
	return bf.World.AcctMap.GetAccount(address)
}
