//go:build e2e
// +build e2e

package e2e

import (
	"testing"

	"github.com/gauthierbraillon/ai-os-installer/internal/install"
)

func TestUbuntuInstall_E2E(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping E2E test in short mode")
	}

	t.Run("AC001: User can install Ubuntu with valid config", func(t *testing.T) {
		installer := install.NewUbuntuInstaller()

		config := install.Config{
			Version: "1.0",
			User: install.UserConfig{
				Name:     "testuser",
				FullName: "Test User",
				Shell:    "bash",
			},
			System: install.SystemConfig{
				Hostname: "test-machine",
				Timezone: "UTC",
				Locale:   "en_US.UTF-8",
			},
			Packages: install.PackageConfig{
				Base: []string{"git", "curl"},
			},
		}

		err := installer.Validate(config)
		if err != nil {
			t.Fatalf("validation failed: %v", err)
		}

		t.Skip("not implemented")
	})
}
