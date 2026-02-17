# Contributing

## Development setup

```bash
git clone https://github.com/gauthierbraillon/ai-os-installer.git
cd ai-os-installer
export ANTHROPIC_API_KEY=sk-...
go test ./...
```

Requirements: Go 1.22+, Docker (for integration tests), QEMU (for E2E tests).

## Workflow: RED → GREEN → REFACTOR → DEPLOY

Every change — bug fix, feature, refactor — follows this cycle. No exceptions.

**1. RED** — write a failing test that describes the required behavior

```bash
go test ./tests/unit/... -run TestYourNewTest  # must fail
```

**2. GREEN** — implement the minimum code to make it pass

```bash
go test ./tests/unit/... -run TestYourNewTest  # must pass
```

**3. REFACTOR** — clean up duplication, naming, structure; run all tests

```bash
go test ./...
```

**4. DEPLOY** — commit in small atomic chunks, push to main

```bash
git add <specific files>
git commit -m "feat: describe the change"
git push
```

See [CLAUDE.md](CLAUDE.md) for the full development guide.

## Test types

| Type | Location | When to run |
|------|----------|-------------|
| Unit (sociable) | `tests/unit/` | Every RED/GREEN/REFACTOR loop |
| Integration | `tests/integration/` | Before push (requires API key) |
| E2E smoke | `tests/e2e/` | After deploy — failure triggers rollback |
| Validation | `tests/validation/` | Static checks (pipeline config, security) |

**Mock only external systems** (Claude API, QEMU, filesystem). Never mock our own code.

## Commit messages

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add gaming use case to config generator
fix: handle empty user input in conversation manager
refactor: extract package resolver into separate function
test: add validation for missing hostname config
docs: update architecture diagram
```

One logical change per commit. Keep commits small enough to revert independently.

## Pull requests

PRs are welcome. Keep them small and focused. Each PR should:

- Have a clear problem statement
- Include tests (RED was written first)
- Pass the full CI pipeline before requesting review

## Reporting issues

Open an issue on GitHub describing:
- What you expected to happen
- What actually happened
- Steps to reproduce
- Environment (OS, Go version, Docker version)
