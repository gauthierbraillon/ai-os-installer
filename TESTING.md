# Testing Guide

## Test Organization

Tests follow ATDD (Acceptance Test-Driven Development) principles. Tests ARE the requirements.

### Test Types

**Unit Tests (colocated with code)**
```bash
go test ./internal/...
```
- Fast (< 2s)
- Mock external systems only (Claude API, file system, QEMU)
- No mocks of our own code (sociable tests)
- Example: `internal/install/installer_test.go`

**E2E Tests (tests/e2e/)**
```bash
go test -tags=e2e ./tests/e2e/
```
- Slow (~15min)
- Full conversation → install flow
- Real VM creation and Ubuntu installation
- Run after deployment as smoke tests

**Integration Tests (use -tags=integration)**
```bash
go test -tags=integration ./...
```
- Medium (~1min)
- Real external systems (Docker, Claude API)
- Rate limited for API calls

## Test Helpers

**Load test fixtures:**
```go
config := tests.LoadTestConfig(t, "testdata/sample_config.json")
```

**Assertions:**
```go
tests.AssertNoError(t, err)
tests.AssertError(t, err, "expected validation error")
```

## Writing Tests (ATDD Workflow)

### RED Phase - Write Acceptance Test First

Test names describe acceptance criteria:

```go
func TestUbuntuInstaller_Install(t *testing.T) {
    t.Run("AC101: User can install Ubuntu with web dev packages", func(t *testing.T) {
        // Given: A valid config for web development
        config := install.Config{
            Packages: install.PackageConfig{
                Dev: []string{"nodejs", "docker", "code"},
            },
        }

        // When: Install is executed
        installer := install.NewUbuntuInstaller()
        err := installer.Install(config)

        // Then: Installation succeeds
        AssertNoError(t, err)
        // And: Packages are installed
        // (verification logic here)
    })
}
```

### GREEN Phase - Make Test Pass

Implement minimum code to pass the test.

### REFACTOR Phase - Clean Up

- Remove duplication
- Improve naming
- Extract helpers
- Keep tests passing

## Running Tests

```bash
# All unit tests (fast)
make test

# With coverage
go test -v -cover ./...

# Only E2E tests
go test -tags=e2e ./tests/e2e/

# Skip slow tests
go test -short ./...
```

## Test Data

Place test fixtures in `tests/testdata/`:
- `sample_config.json` - Valid install config
- Add more as needed

## CI/CD Integration

Tests run automatically on every push:
1. Unit tests (must pass to merge)
2. Lint (must pass to merge)
3. Build (must succeed to merge)
4. E2E smoke tests (run after deploy, trigger rollback if fail)
