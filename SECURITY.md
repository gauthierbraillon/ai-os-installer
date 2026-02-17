# Security Policy

Supported version: `main` (latest).

## Reporting a vulnerability

Open a [GitHub Security Advisory](https://github.com/gauthierbraillon/ai-os-installer/security/advisories/new) to report privately. Do not open a public issue.

## Security practices

**Secrets** — API keys via environment variables only; never logged or hardcoded.

**Dependencies** — GitHub Actions pinned to commit SHAs; Docker base images pinned to digest hashes; supply chain validation runs on every push (`tests/validation/`).

**Runtime** — user inputs validated before any system call; VMs run with minimal privileges.

**CI/CD** — `golangci-lint`, `gofmt`, SBOM and provenance attestation on every Docker image push.
