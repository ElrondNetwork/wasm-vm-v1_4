package nodepart

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/ElrondNetwork/arwen-wasm-vm/ipc/common"
	vmcommon "github.com/ElrondNetwork/elrond-vm-common"
)

var _ vmcommon.VMExecutionHandler = (*ArwenDriver)(nil)

// ArwenDriver is
type ArwenDriver struct {
	blockchainHook vmcommon.BlockchainHook
	vmType         []byte
	blockGasLimit  uint64
	gasSchedule    map[string]map[string]uint64

	arwenInputRead   *os.File
	arwenInputWrite  *os.File
	arwenOutputRead  *os.File
	arwenOutputWrite *os.File
	command          *exec.Cmd
	part             *NodePart
}

// NewArwenDriver creates
func NewArwenDriver(
	blockchainHook vmcommon.BlockchainHook,
	vmType []byte,
	blockGasLimit uint64,
	gasSchedule map[string]map[string]uint64,
) (*ArwenDriver, error) {
	driver := &ArwenDriver{
		blockchainHook: blockchainHook,
		vmType:         vmType,
		blockGasLimit:  blockGasLimit,
		gasSchedule:    gasSchedule,
	}

	err := driver.startArwen()
	return driver, err
}

func (driver *ArwenDriver) startArwen() error {
	common.LogDebug("ArwenDriver.startArwen()")

	driver.resetPipeStreams()

	arwenPath, err := driver.getArwenPath()
	if err != nil {
		return err
	}

	driver.command = exec.Command(arwenPath)
	driver.command.Stdout = os.Stdout
	driver.command.Stderr = os.Stderr
	// TODO: pass vmType, blockGasLimit and gasSchedule when starting Arwen

	driver.command.ExtraFiles = []*os.File{driver.arwenInputRead, driver.arwenOutputWrite}
	err = driver.command.Start()
	if err != nil {
		return err
	}

	driver.part, err = NewNodePart(driver.arwenOutputRead, driver.arwenInputWrite, driver.blockchainHook)
	if err != nil {
		return err
	}

	return nil
}

func (driver *ArwenDriver) getArwenPath() (string, error) {
	arwenPath := os.Getenv("ARWEN_PATH")
	common.LogDebug("ARWEN_PATH environment variable: %s", arwenPath)

	if fileExists(arwenPath) {
		return arwenPath, nil
	}

	arwenPath = path.Join(".", "arwen")
	if fileExists(arwenPath) {
		return arwenPath, nil
	}

	return "", common.ErrArwenNotFound
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func (driver *ArwenDriver) resetPipeStreams() error {
	closeFile(driver.arwenInputRead)
	closeFile(driver.arwenInputWrite)
	closeFile(driver.arwenOutputRead)
	closeFile(driver.arwenOutputWrite)

	var err error

	driver.arwenInputRead, driver.arwenInputWrite, err = os.Pipe()
	if err != nil {
		return err
	}

	driver.arwenOutputRead, driver.arwenOutputWrite, err = os.Pipe()
	if err != nil {
		return err
	}

	return nil
}

func closeFile(file *os.File) {
	if file != nil {
		err := file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot close file.\n")
		}
	}
}

func (driver *ArwenDriver) forceRestartArwen() error {
	if !driver.command.ProcessState.Exited() {
		err := driver.command.Process.Kill()
		if err != nil {
			return err
		}
	}

	err := driver.startArwen()
	return err
}

func (driver *ArwenDriver) restartArwenIfNecessary() error {
	state := driver.command.ProcessState
	stopped := state != nil && state.Exited()
	if !stopped {
		return nil
	}

	err := driver.startArwen()
	return err
}

// RunSmartContractCreate creates
func (driver *ArwenDriver) RunSmartContractCreate(input *vmcommon.ContractCreateInput) (*vmcommon.VMOutput, error) {
	driver.restartArwenIfNecessary()

	request := &common.ContractRequest{
		Action:      "Deploy",
		CreateInput: input,
	}

	response, err := driver.part.StartLoop(request)
	if err != nil {
		driver.stopArwen()
		return nil, common.WrapCriticalError(err)
	}

	return response.VMOutput, response.GetError()
}

// RunSmartContractCall calls
func (driver *ArwenDriver) RunSmartContractCall(input *vmcommon.ContractCallInput) (*vmcommon.VMOutput, error) {
	driver.restartArwenIfNecessary()

	request := &common.ContractRequest{
		Action:    "Call",
		CallInput: input,
	}

	response, err := driver.part.StartLoop(request)
	if err != nil {
		driver.stopArwen()
		return nil, common.WrapCriticalError(err)
	}

	return response.VMOutput, response.GetError()
}

// TODO: func OnRoundEnded -> triggers Arwen restart.

// Close stops Arwen
func (driver *ArwenDriver) Close() error {
	err := driver.stopArwen()
	if err != nil {
		return err
	}

	return nil
}

func (driver *ArwenDriver) stopArwen() error {
	err := driver.command.Process.Kill()
	if err != nil {
		common.LogError("stopArwen error=%s", err)
	}

	return err
}

// TODO: Add test for arwen crash. Run Tx, force crash, Run Tx again.
