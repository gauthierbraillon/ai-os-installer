package validation

import (
	"bufio"
	"os"
	"regexp"
	"strings"
	"testing"
)

var mutableTagPattern = regexp.MustCompile(`uses:\s+\S+@(v\d[\w.]*)$`)
var dockerFromPattern = regexp.MustCompile(`^FROM\s+(\S+)`)
var digestPattern = regexp.MustCompile(`@sha256:[a-f0-9]{64}`)

func openFile(t *testing.T, path string) *os.File {
	t.Helper()
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("cannot open %s: %v", path, err)
	}
	t.Cleanup(func() {
		if err := f.Close(); err != nil {
			t.Errorf("failed to close %s: %v", path, err)
		}
	})
	return f
}

func TestSC001_GitHubActionsArePinnedToCommitSHAs(t *testing.T) {
	f := openFile(t, "../../.github/workflows/ci-cd.yml")

	lineNum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if m := mutableTagPattern.FindStringSubmatch(line); m != nil {
			t.Errorf("SC001: line %d uses mutable tag %q — pin to a commit SHA instead: %s", lineNum, m[1], line)
		}
	}
}

func TestSC002_DockerfileBaseImagesArePinnedToDigests(t *testing.T) {
	f := openFile(t, "../../docker/Dockerfile")

	lineNum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		m := dockerFromPattern.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		image := m[1]
		if image == "scratch" || strings.HasPrefix(image, "--") {
			continue
		}
		if !digestPattern.MatchString(image) {
			t.Errorf("SC002: line %d FROM image not pinned to digest: %s", lineNum, image)
		}
	}
}

func TestSC003_DockerignoreExists(t *testing.T) {
	if _, err := os.Stat("../../.dockerignore"); os.IsNotExist(err) {
		t.Error("SC003: .dockerignore does not exist — sensitive files may leak into Docker build context")
	}
}

func TestSC004_GolangciLintVersionIsPinned(t *testing.T) {
	f := openFile(t, "../../.github/workflows/ci-cd.yml")

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "version:") && strings.Contains(line, "latest") {
			t.Errorf("SC004: golangci-lint version is unpinned (using 'latest'): %s", line)
		}
	}
}
