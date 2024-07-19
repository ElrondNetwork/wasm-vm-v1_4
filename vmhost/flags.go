package vmhost

import "github.com/multiversx/mx-chain-core-go/core"

const (
	// MultiESDTTransferFixOnCallBackFlag defines the flag that activates the multi esdt transfer fix on callback
	MultiESDTTransferFixOnCallBackFlag core.EnableEpochFlag = "MultiESDTTransferFixOnCallBackFlag"
	// RemoveNonUpdatedStorageFlag defines the flag that activates the remove non updated storage fix
	RemoveNonUpdatedStorageFlag core.EnableEpochFlag = "RemoveNonUpdatedStorageFlag"
	// CreateNFTThroughExecByCallerFlag defines the flag that activates the create nft through exec by caller fix
	CreateNFTThroughExecByCallerFlag core.EnableEpochFlag = "CreateNFTThroughExecByCallerFlag"
	// StorageAPICostOptimizationFlag defines the flag that activates the storage api cost optimization
	StorageAPICostOptimizationFlag core.EnableEpochFlag = "StorageAPICostOptimizationFlag"
	// CheckExecuteOnReadOnlyFlag defines the flag that activates the check execute on read only
	CheckExecuteOnReadOnlyFlag core.EnableEpochFlag = "CheckExecuteOnReadOnlyFlag"
	// FailExecutionOnEveryAPIErrorFlag defines the flag that activates the fail execution on every api error
	FailExecutionOnEveryAPIErrorFlag core.EnableEpochFlag = "FailExecutionOnEveryAPIErrorFlag"
	// ManagedCryptoAPIsFlag defines the flag that activates the manage crypto apis
	ManagedCryptoAPIsFlag core.EnableEpochFlag = "ManagedCryptoAPIsFlag"
	// DisableExecByCallerFlag defines the flag that activates disable exec by caller
	DisableExecByCallerFlag core.EnableEpochFlag = "DisableExecByCallerFlag"
	// RefactorContextFlag defines the flag that activates the refactor context
	RefactorContextFlag core.EnableEpochFlag = "RefactorContextFlag"
	// RuntimeMemStoreLimitFlag defines the flag that activates the runtime mem store limit
	RuntimeMemStoreLimitFlag core.EnableEpochFlag = "RuntimeMemStoreLimitFlag"
	// RuntimeCodeSizeFixFlag defines the flag that activates the runtime code size fix
	RuntimeCodeSizeFixFlag core.EnableEpochFlag = "RuntimeCodeSizeFixFlag"
	// FixOOGReturnCodeFlag defines the flag that activates the fix oog return code
	FixOOGReturnCodeFlag core.EnableEpochFlag = "FixOOGReturnCodeFlag"
	// ValidationOnGobDecodeFlag defines the flag that allows the GobDecode validation added on go1.21
	ValidationOnGobDecodeFlag core.EnableEpochFlag = "ValidationOnGobDecodeFlag"
	// all new flags must be added to allFlags slice from hostCore/host
)
