package config

import (
	"slices"
	"testing"

	"github.com/gauthierbraillon/ai-os-installer/internal/install"
)

func TestAC101_WebDevUseCaseProducesRelevantPackages(t *testing.T) {
	useCases := []string{"web development", "web dev", "I want to do web development"}

	for _, useCase := range useCases {
		t.Run(useCase, func(t *testing.T) {
			cfg := Generate(useCase)

			allPackages := append(cfg.Packages.Base, cfg.Packages.Dev...)
			assertContains(t, allPackages, "nodejs", "web dev config must include nodejs")
			assertContains(t, allPackages, "docker", "web dev config must include docker")
			assertContains(t, allPackages, "git", "web dev config must include git")
		})
	}
}

func TestAC102_GamingUseCaseProducesRelevantPackages(t *testing.T) {
	useCases := []string{"gaming", "I want to play games", "game development"}

	for _, useCase := range useCases {
		t.Run(useCase, func(t *testing.T) {
			cfg := Generate(useCase)

			allPackages := append(cfg.Packages.Base, cfg.Packages.Dev...)
			assertContains(t, allPackages, "steam", "gaming config must include steam")
		})
	}
}

func TestAC103_DataScienceUseCaseProducesRelevantPackages(t *testing.T) {
	useCases := []string{"data science", "machine learning", "AI research"}

	for _, useCase := range useCases {
		t.Run(useCase, func(t *testing.T) {
			cfg := Generate(useCase)

			allPackages := append(cfg.Packages.Base, cfg.Packages.Dev...)
			assertContains(t, allPackages, "python3", "data science config must include python3")
			assertContains(t, allPackages, "jupyter", "data science config must include jupyter")
		})
	}
}

func TestAC104_GeneratedConfigAlwaysHasBasePackages(t *testing.T) {
	cfg := Generate("anything at all")

	assertContains(t, cfg.Packages.Base, "git", "base packages must always include git")
	assertContains(t, cfg.Packages.Base, "curl", "base packages must always include curl")
	assertContains(t, cfg.Packages.Base, "vim", "base packages must always include vim")
}

func TestAC105_GeneratedConfigVersionIsSet(t *testing.T) {
	cfg := Generate("web development")

	if cfg.Version != "1.0" {
		t.Errorf("expected version 1.0, got %s", cfg.Version)
	}
}

func TestAC106_UnknownUseCaseReturnsBaseConfig(t *testing.T) {
	cfg := Generate("something totally unrecognised xyz")

	if len(cfg.Packages.Base) == 0 {
		t.Error("unknown use case must still return base packages")
	}
	if cfg.Version == "" {
		t.Error("unknown use case must still return a versioned config")
	}
}

func assertContains(t *testing.T, packages []string, pkg, msg string) {
	t.Helper()
	if !slices.Contains(packages, pkg) {
		t.Errorf("%s: package list %v does not contain %q", msg, packages, pkg)
	}
}

func TestAC101_GeneratedConfigPassesInstallValidation(t *testing.T) {
	cfg := Generate("web development")
	cfg.User = install.UserConfig{Name: "john", FullName: "John Doe", Shell: "bash"}
	cfg.System = install.SystemConfig{Hostname: "dev-machine", Timezone: "UTC", Locale: "en_US.UTF-8"}

	installer := install.NewUbuntuInstaller()
	if err := installer.Validate(cfg); err != nil {
		t.Errorf("generated config failed validation: %v", err)
	}
}
