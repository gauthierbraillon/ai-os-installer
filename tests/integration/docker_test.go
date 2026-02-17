//go:build integration

package integration

import (
	"os/exec"
	"strings"
	"testing"
)

func TestIT001_DockerImageBuildsAndRuns(t *testing.T) {
	build := exec.Command("docker", "build",
		"-f", "../../docker/Dockerfile",
		"-t", "ai-os-installer:integration-test",
		"../..")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("IT001: docker build failed: %v\n%s", err, out)
	}

	run := exec.Command("docker", "run", "--rm", "ai-os-installer:integration-test")
	out, err := run.CombinedOutput()
	if err != nil {
		t.Fatalf("IT001: container run failed: %v\n%s", err, out)
	}

	if !strings.Contains(string(out), "AI OS Installer") {
		t.Errorf("IT001: expected installer output, got: %s", out)
	}
}
