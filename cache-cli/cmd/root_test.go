package cmd

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/semaphoreci/toolbox/cache-cli/pkg/storage"
	assert "github.com/stretchr/testify/assert"
)

type TestBackend struct {
	envVars      map[string]string
	runInWindows bool
}

var testBackends = map[string]TestBackend{
	"sftp": {
		runInWindows: false,
		envVars: map[string]string{
			"NEETO_CI_CACHE_BACKEND":          "sftp",
			"NEETO_CI_CACHE_URL":              "sftp-server:22",
			"NEETO_CI_CACHE_USERNAME":         "tester",
			"NEETO_CI_CACHE_PRIVATE_KEY_PATH": "/root/.ssh/neeto_ci_cache_key",
		},
	},
}

func runTestForAllBackends(t *testing.T, test func(string, storage.Storage)) {
	for backendType, testBackend := range testBackends {
		if runtime.GOOS == "windows" && !testBackend.runInWindows {
			continue
		}

		for envVarName, envVarValue := range testBackend.envVars {
			os.Setenv(envVarName, envVarValue)
		}

		storage, err := storage.InitStorage()
		assert.Nil(t, err)
		test(backendType, storage)
	}
}

func readOutputFromFile(t *testing.T) string {
	path := filepath.Join(os.TempDir(), "cache_log")

	defer os.Truncate(path, 0)

	output, err := ioutil.ReadFile(path)
	assert.NoError(t, err)

	return string(output)
}

func openLogfileForTests(t *testing.T) io.Writer {
	filePath := filepath.Join(os.TempDir(), "cache_log")
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	assert.NoError(t, err)
	return io.MultiWriter(f, os.Stdout)
}
