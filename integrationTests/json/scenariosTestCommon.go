package vmjsonintegrationtest

import (
	"os"
	"path"
	"path/filepath"
	"testing"

	logger "github.com/multiversx/mx-chain-logger-go"
	am "github.com/multiversx/mx-chain-vm-v1_4-go/scenarioexec"
	mc "github.com/multiversx/mx-chain-vm-v1_4-go/scenarios/controller"
	"github.com/stretchr/testify/require"
)

func init() {
	_ = logger.SetLogLevel("*:INFO")
}

func getTestRoot() string {
	exePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	vmTestRoot := filepath.Join(exePath, "../../test")
	return vmTestRoot
}

func runAllTestsInFolder(t *testing.T, folder string) {
	runTestsInFolder(t, folder, []string{})
}

func runTestsInFolder(t *testing.T, folder string, exclusions []string) {
	executor, err := am.NewVMTestExecutor()
	require.Nil(t, err)
	defer executor.Close()

	runner := mc.NewScenarioRunner(
		executor,
		mc.NewDefaultFileResolver(),
	)

	err = runner.RunAllJSONScenariosInDirectory(
		getTestRoot(),
		folder,
		".scen.json",
		exclusions,
		mc.DefaultRunScenarioOptions())

	if err != nil {
		t.Error(err)
	}
}

func runSingleTestReturnError(folder string, filename string) error {
	executor, err := am.NewVMTestExecutor()
	if err != nil {
		return err
	}
	defer executor.Close()

	runner := mc.NewScenarioRunner(
		executor,
		mc.NewDefaultFileResolver(),
	)

	fullPath := path.Join(getTestRoot(), folder)
	fullPath = path.Join(fullPath, filename)

	return runner.RunSingleJSONScenario(
		fullPath,
		mc.DefaultRunScenarioOptions())
}
