package config

import (
	"strings"

	"github.com/gauthierbraillon/ai-os-installer/internal/install"
)

var basePackages = []string{"git", "curl", "vim"}

var useCasePackages = map[string][]string{
	"web":     {"nodejs", "docker", "npm"},
	"gaming":  {"steam", "vulkan-utils"},
	"game":    {"steam", "vulkan-utils"},
	"data":    {"python3", "jupyter", "python3-pip"},
	"machine": {"python3", "jupyter", "python3-pip"},
	"ai":      {"python3", "jupyter", "python3-pip"},
	"ml":      {"python3", "jupyter", "python3-pip"},
}

func Generate(useCase string) install.Config {
	lower := strings.ToLower(useCase)
	devPackages := devPackagesFor(lower)

	return install.Config{
		Version: "1.0",
		Packages: install.PackageConfig{
			Base: basePackages,
			Dev:  devPackages,
		},
	}
}

func devPackagesFor(useCase string) []string {
	seen := make(map[string]bool)
	var packages []string

	for keyword, pkgs := range useCasePackages {
		if strings.Contains(useCase, keyword) {
			for _, pkg := range pkgs {
				if !seen[pkg] {
					seen[pkg] = true
					packages = append(packages, pkg)
				}
			}
		}
	}

	return packages
}
