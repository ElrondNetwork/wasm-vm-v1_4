package hostCore

import (
	"testing"

	"github.com/multiversx/mx-chain-vm-v1_4-go/vmhost"
	"github.com/multiversx/mx-chain-vm-v1_4-go/vmhost/mock"

	"github.com/multiversx/mx-chain-scenario-go/worldmock"
	"github.com/multiversx/mx-chain-vm-common-go/builtInFunctions"
	"github.com/multiversx/mx-chain-vm-common-go/parsers"
	"github.com/stretchr/testify/require"
)

func TestNewVMHost(t *testing.T) {
	blockchainHook := mock.NewMockWorldVM14()
	bfc := builtInFunctions.NewBuiltInFunctionContainer()
	epochNotifier := &mock.EpochNotifierStub{}
	epochsHandler := &mock.EnableEpochsHandlerStub{}
	vmType := []byte("vmType")
	esdtTransferParser, err := parsers.NewESDTTransferParser(worldmock.WorldMarshalizer)
	require.Nil(t, err)

	makeHostParameters := func() *vmhost.VMHostParameters {
		return &vmhost.VMHostParameters{
			VMType:               vmType,
			ESDTTransferParser:   esdtTransferParser,
			BuiltInFuncContainer: bfc,
			EpochNotifier:        epochNotifier,
			EnableEpochsHandler:  epochsHandler,
			Hasher:               worldmock.DefaultHasher,
		}
	}

	t.Run("NilBlockchainHook", func(t *testing.T) {
		host, err := NewVMHost(nil, makeHostParameters())
		require.Nil(t, host)
		require.ErrorIs(t, err, vmhost.ErrNilBlockChainHook)
	})
	t.Run("NilHostParameters", func(t *testing.T) {
		host, err := NewVMHost(blockchainHook, nil)
		require.Nil(t, host)
		require.ErrorIs(t, err, vmhost.ErrNilHostParameters)
	})
	t.Run("NilESDTTransferParser", func(t *testing.T) {
		hostParameters := makeHostParameters()
		hostParameters.ESDTTransferParser = nil
		host, err := NewVMHost(blockchainHook, hostParameters)
		require.Nil(t, host)
		require.ErrorIs(t, err, vmhost.ErrNilESDTTransferParser)
	})
	t.Run("NilBuiltInFunctionsContainer", func(t *testing.T) {
		hostParameters := makeHostParameters()
		hostParameters.BuiltInFuncContainer = nil
		host, err := NewVMHost(blockchainHook, hostParameters)
		require.Nil(t, host)
		require.ErrorIs(t, err, vmhost.ErrNilBuiltInFunctionsContainer)
	})
	t.Run("NilEpochNotifier", func(t *testing.T) {
		hostParameters := makeHostParameters()
		hostParameters.EpochNotifier = nil
		host, err := NewVMHost(blockchainHook, hostParameters)
		require.Nil(t, host)
		require.ErrorIs(t, err, vmhost.ErrNilEpochNotifier)
	})
	t.Run("NilEnableEpochsHandler", func(t *testing.T) {
		hostParameters := makeHostParameters()
		hostParameters.EnableEpochsHandler = nil
		host, err := NewVMHost(blockchainHook, hostParameters)
		require.Nil(t, host)
		require.ErrorIs(t, err, vmhost.ErrNilEnableEpochsHandler)
	})
	t.Run("NilHasher", func(t *testing.T) {
		hostParameters := makeHostParameters()
		hostParameters.Hasher = nil
		host, err := NewVMHost(blockchainHook, hostParameters)
		require.Nil(t, host)
		require.ErrorIs(t, err, vmhost.ErrNilHasher)
	})
	t.Run("NilVMType", func(t *testing.T) {
		hostParameters := makeHostParameters()
		hostParameters.VMType = nil
		host, err := NewVMHost(blockchainHook, hostParameters)
		require.Nil(t, host)
		require.ErrorIs(t, err, vmhost.ErrNilVMType)
	})
}
