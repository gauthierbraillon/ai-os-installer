package tests

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/gauthierbraillon/ai-os-installer/internal/install"
)

func LoadTestConfig(t *testing.T, filename string) install.Config {
	t.Helper()

	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("failed to read test config: %v", err)
	}

	var config install.Config
	if err := json.Unmarshal(data, &config); err != nil {
		t.Fatalf("failed to parse test config: %v", err)
	}

	return config
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func AssertError(t *testing.T, err error, msg string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error: %s", msg)
	}
}
