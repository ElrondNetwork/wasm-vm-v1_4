package scenarioexec

import (
	"errors"

	"github.com/multiversx/mx-chain-core-go/core/check"
	vmi "github.com/multiversx/mx-chain-vm-common-go"
	mc "github.com/multiversx/mx-chain-vm-v1_4-go/scenarios/controller"
	fr "github.com/multiversx/mx-chain-vm-v1_4-go/scenarios/fileresolver"
	mj "github.com/multiversx/mx-chain-vm-v1_4-go/scenarios/model"
	"github.com/multiversx/mx-chain-vm-v1_4-go/vmhost"
)

// Reset clears state/world.
// Is called in RunAllJSONScenariosInDirectory, but not in RunSingleJSONScenario.
func (ae *VMTestExecutor) Reset() {
	if !check.IfNil(ae.vmHost) {
		ae.vmHost.Reset()
	}
	ae.World.Clear()
}

// Close will simply close the VM
func (ae *VMTestExecutor) Close() {
	if !check.IfNil(ae.vmHost) {
		ae.vmHost.Reset()
	}
}

// ExecuteScenario executes an individual test.
func (ae *VMTestExecutor) ExecuteScenario(scenario *mj.Scenario, fileResolver fr.FileResolver) error {
	ae.fileResolver = fileResolver
	ae.checkGas = scenario.CheckGas
	resetGasTracesIfNewTest(ae, scenario)

	err := ae.InitVM(scenario.GasSchedule)
	if err != nil {
		return err
	}

	txIndex := 0
	for _, generalStep := range scenario.Steps {
		setGasTraceInMetering(ae, true)
		err := ae.ExecuteStep(generalStep)
		if err != nil {
			return err
		}
		setGasTraceInMetering(ae, false)
		txIndex++
	}

	return nil
}

// ExecuteStep executes an individual step from a scenario.
func (ae *VMTestExecutor) ExecuteStep(generalStep mj.Step) error {
	err := error(nil)

	switch step := generalStep.(type) {
	case *mj.ExternalStepsStep:
		err = ae.ExecuteExternalStep(step)
		length := len(ae.scenarioTraceGas)
		ae.scenarioTraceGas = ae.scenarioTraceGas[:length-1]
		return err
	case *mj.SetStateStep:
		err = ae.ExecuteSetStateStep(step)
	case *mj.CheckStateStep:
		err = ae.ExecuteCheckStateStep(step)
	case *mj.TxStep:
		_, err = ae.ExecuteTxStep(step)
	case *mj.DumpStateStep:
		err = ae.DumpWorld()
	}

	logGasTrace(ae)

	return err
}

// ExecuteExternalStep executes an external step referenced by the scenario.
func (ae *VMTestExecutor) ExecuteExternalStep(step *mj.ExternalStepsStep) error {
	log.Trace("ExternalStepsStep", "path", step.Path)
	if len(step.Comment) > 0 {
		log.Trace("ExternalStepsStep", "comment", step.Comment)
	}

	fileResolverBackup := ae.fileResolver
	clonedFileResolver := ae.fileResolver.Clone()
	externalStepsRunner := mc.NewScenarioRunner(ae, clonedFileResolver)

	extAbsPth := ae.fileResolver.ResolveAbsolutePath(step.Path)
	setExternalStepGasTracing(ae, step)

	err := externalStepsRunner.RunSingleJSONScenario(extAbsPth, mc.DefaultRunScenarioOptions())
	if err != nil {
		return err
	}

	ae.fileResolver = fileResolverBackup

	return nil
}

// ExecuteSetStateStep executes a SetStateStep.
func (ae *VMTestExecutor) ExecuteSetStateStep(step *mj.SetStateStep) error {
	if len(step.Comment) > 0 {
		log.Trace("SetStateStep", "comment", step.Comment)
	}

	for _, scenAccount := range step.Accounts {
		if scenAccount.Update {
			err := ae.UpdateAccount(scenAccount)
			if err != nil {
				log.Debug("could not update account", err)
				return err
			}
		} else {
			err := ae.PutNewAccount(scenAccount)
			if err != nil {
				log.Debug("could not put new account", err)
				return err
			}
		}
	}

	// replace block info
	ae.World.PreviousBlockInfo = convertBlockInfo(step.PreviousBlockInfo, ae.World.PreviousBlockInfo)
	ae.World.CurrentBlockInfo = convertBlockInfo(step.CurrentBlockInfo, ae.World.CurrentBlockInfo)
	ae.World.Blockhashes = step.BlockHashes.ToValues()

	// append NewAddressMocks
	err := validateNewAddressMocks(step.NewAddressMocks)
	if err != nil {
		return err
	}
	addressMocksToAdd := convertNewAddressMocks(step.NewAddressMocks)
	ae.World.NewAddressMocks = append(ae.World.NewAddressMocks, addressMocksToAdd...)

	return nil
}

// ExecuteTxStep executes a TxStep.
func (ae *VMTestExecutor) ExecuteTxStep(step *mj.TxStep) (*vmi.VMOutput, error) {
	log.Trace("ExecuteTxStep", "id", step.TxIdent)
	if len(step.Comment) > 0 {
		log.Trace("ExecuteTxStep", "comment", step.Comment)
	}

	if step.DisplayLogs {
		vmhost.SetLoggingForTests()
	}

	output, err := ae.executeTx(step.TxIdent, step.Tx)
	if err != nil {
		return nil, err
	}

	if step.DisplayLogs {
		vmhost.DisableLoggingForTests()
	}

	// check results
	if step.ExpectedResult != nil {
		err = ae.checkTxResults(step.TxIdent, step.ExpectedResult, ae.checkGas, output)
		if err != nil {
			return nil, err
		}
	}

	return output, nil
}

// PutNewAccount Puts a new account in world account map. Overwrites.
func (ae *VMTestExecutor) PutNewAccount(scenAccount *mj.Account) error {
	worldAccount, err := convertAccount(scenAccount, ae.World)
	if err != nil {
		return err
	}
	err = validateSetStateAccount(scenAccount, worldAccount)
	if err != nil {
		return err
	}

	ae.World.AcctMap.PutAccount(worldAccount)
	return nil
}

// UpdateAccount Updates an account in world account map.
func (ae *VMTestExecutor) UpdateAccount(scenAccount *mj.Account) error {
	worldAccount, err := convertAccount(scenAccount, ae.World)
	if err != nil {
		return err
	}
	err = validateSetStateAccount(scenAccount, worldAccount)
	if err != nil {
		return err
	}

	existingAccount := ae.World.AcctMap.GetAccount(scenAccount.Address.Value)
	if existingAccount == nil {
		return errors.New("account not found. could not update")
	}

	for k, v := range worldAccount.Storage {
		existingAccount.Storage[k] = v
	}
	if !scenAccount.Nonce.Unspecified {
		existingAccount.Nonce = worldAccount.Nonce
	}
	if !scenAccount.Balance.Unspecified {
		existingAccount.Balance = worldAccount.Balance
	}
	if !scenAccount.Username.Unspecified {
		existingAccount.Username = worldAccount.Username
	}
	if !scenAccount.Owner.Unspecified {
		existingAccount.OwnerAddress = worldAccount.OwnerAddress
	}
	if !scenAccount.Code.Unspecified {
		existingAccount.Code = worldAccount.Code
	}
	if !scenAccount.Shard.Unspecified {
		existingAccount.ShardID = worldAccount.ShardID
	}
	existingAccount.AsyncCallData = worldAccount.AsyncCallData

	ae.World.AcctMap.PutAccount(existingAccount)
	return nil
}
