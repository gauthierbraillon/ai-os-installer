# Security Policy

## Supported versions

| Version | Supported |
|---------|-----------|
| latest (main) | Yes |

## Reporting a vulnerability

Open a [GitHub Security Advisory](https://github.com/gauthierbraillon/ai-os-installer/security/advisories/new) to report a vulnerability privately.

Do not open a public issue for security vulnerabilities.

We will respond within 5 business days and coordinate a fix and disclosure timeline with you.

## Security practices

**Secrets**
- API keys are passed via environment variables only, never hardcoded
- No secrets are logged or included in error messages
- GitHub Actions secrets are used for CI credentials

**Dependencies**
- GitHub Actions are pinned to commit SHAs (not mutable version tags)
- Docker base images are pinned to digest hashes
- Supply chain validation tests run on every push (`tests/validation/`)

**Runtime**
- User inputs are validated before being passed to any system command
- VMs run with minimal privileges
- The installer container requires `--privileged` only for VM management

**CI/CD**
- `golangci-lint` runs `errcheck`, `staticcheck`, `govet`, and `unused` on every push
- `gofmt` formatting is enforced in CI
- SBOM and provenance attestation are generated for every Docker image push
