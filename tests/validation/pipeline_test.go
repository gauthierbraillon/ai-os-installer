package validation

import (
	"bufio"
	"strings"
	"testing"
)

func TestSC005_PipelineHasIntegrationJob(t *testing.T) {
	f := openFile(t, "../../.github/workflows/ci-cd.yml")

	found := false
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "integration:" {
			found = true
			break
		}
	}
	if !found {
		t.Error("SC005: pipeline has no 'integration:' job — integration tests must run before docker push")
	}
}

func TestSC006_DockerJobDoesNotPushLatestDirectly(t *testing.T) {
	f := openFile(t, "../../.github/workflows/ci-cd.yml")

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "type=raw,value=latest" {
			t.Error("SC006: docker job pushes 'latest' directly — latest must only be promoted after E2E passes")
			return
		}
	}
}

func TestSC007_PipelineHasPromoteJob(t *testing.T) {
	f := openFile(t, "../../.github/workflows/ci-cd.yml")

	found := false
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "promote:" {
			found = true
			break
		}
	}
	if !found {
		t.Error("SC007: pipeline has no 'promote:' job — latest tag must only be pushed after E2E smoke tests pass")
	}
}
