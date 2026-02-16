# AI OS Installer - Architecture

## Vision

Conversational AI-guided Linux installer that makes OS installation as easy as talking to an expert friend.

## MVP Scope (Week 1-3)

**In scope:**
- Ubuntu 24.04 LTS installation only
- Conversational setup via terminal
- Testing in VMs (QEMU/VirtualBox)
- Cloud API (Claude/OpenAI) for conversations
- Docker-based distribution
- Automated CI/CD pipeline

**Out of scope (future):**
- Other distros (Arch, Fedora)
- Bootable USB creation
- Offline mode with local LLM
- GUI interface
- Multi-language support

## Architecture Decisions

### 1. Technology Stack

**Language:** Go 1.21+
- Reason: Single binary, fast startup (<100ms), small footprint (~20MB), excellent for CLI tools
- User benefit: 7x smaller download, 20x faster startup vs Python

**AI Provider:** Claude API (Anthropic)
- Reason: Best conversational quality, API-first design
- Fallback: OpenAI GPT-4 for users without Claude access

**OS Target:** Ubuntu 24.04 LTS
- Reason: Most popular distro, excellent tooling, stable

**Distribution:** Docker container + standalone binary
- Reason: Easy to distribute, consistent environment, simple CI/CD
- Users can also download single binary without Docker

**Testing:** Go testing package + testify + Docker + QEMU
- Reason: Standard library, table-driven tests, no external test framework needed

**CI/CD:** GitHub Actions → Docker Hub + GitHub Releases
- Reason: Free for public repos, standard tooling

### 2. System Components

```
┌─────────────────────────────────────────────────────┐
│                    User Terminal                     │
└─────────────────────┬───────────────────────────────┘
                      │
                      ▼
┌─────────────────────────────────────────────────────┐
│              Conversation Manager                    │
│  - Collects user requirements                        │
│  - Manages conversation state                        │
│  - Validates responses                               │
└─────────────────────┬───────────────────────────────┘
                      │
                      ▼
┌─────────────────────────────────────────────────────┐
│                  AI Engine                           │
│  - Claude API client                                 │
│  - Prompt engineering                                │
│  - Response parsing                                  │
└─────────────────────┬───────────────────────────────┘
                      │
                      ▼
┌─────────────────────────────────────────────────────┐
│              Config Generator                        │
│  - Converts conversation → install config            │
│  - Validates config                                  │
│  - Generates preseed/cloud-init                      │
└─────────────────────┬───────────────────────────────┘
                      │
                      ▼
┌─────────────────────────────────────────────────────┐
│              Install Executor                        │
│  - Runs installation in VM                           │
│  - Progress reporting                                │
│  - Error handling                                    │
└─────────────────────────────────────────────────────┘
```

### 3. Core Modules

**`internal/conversation/`**
- `manager.go` - Main conversation loop
- `state.go` - Conversation state management
- `validator.go` - User input validation

**`internal/ai/`**
- `client.go` - AI API client (Claude/OpenAI)
- `prompts.go` - Prompt templates
- `parser.go` - Response parsing

**`internal/config/`**
- `generator.go` - Convert conversation → install config
- `validator.go` - Config validation
- `templates/` - Preseed/cloud-init templates

**`internal/install/`**
- `executor.go` - VM installation orchestrator
- `vm.go` - VM lifecycle (create, start, stop)
- `progress.go` - Installation progress tracking

**`pkg/logger/`**
- `logger.go` - Structured logging

**`pkg/errors/`**
- `errors.go` - Custom error types

**`cmd/ai-os-installer/`**
- `main.go` - CLI entry point
- `root.go` - Root command (cobra)

**Tests** (colocated with code)
- `*_test.go` - Unit tests
- `*_integration_test.go` - Integration tests
- `testdata/` - Test fixtures

### 4. Data Flow

**Phase 1: Conversation (User → AI → Config)**
```
User: "I want a web dev machine"
  ↓
AI: "What languages? (Python, Node, Go...)"
  ↓
User: "Python and Node"
  ↓
AI: "Desktop environment? (GNOME, KDE, minimal)"
  ↓
User: "GNOME"
  ↓
Config: {
  "use_case": "web_dev",
  "languages": ["python", "nodejs"],
  "desktop": "gnome",
  "packages": ["python3", "nodejs", "npm", "git", "vscode"]
}
```

**Phase 2: Installation (Config → Ubuntu)**
```
Config
  ↓
Generate preseed file
  ↓
Create VM (QEMU)
  ↓
Mount Ubuntu ISO
  ↓
Run automated install
  ↓
Post-install scripts
  ↓
Verify installation
  ↓
Success!
```

### 5. Configuration Format

**Install Config (JSON):**
```json
{
  "version": "1.0",
  "user": {
    "name": "john",
    "full_name": "John Doe",
    "password_hash": "...",
    "shell": "bash"
  },
  "system": {
    "hostname": "dev-machine",
    "timezone": "Europe/Brussels",
    "locale": "en_US.UTF-8"
  },
  "packages": {
    "base": ["git", "curl", "vim"],
    "dev": ["python3", "nodejs", "docker"],
    "desktop": ["gnome-shell", "firefox"]
  },
  "services": {
    "enable": ["docker", "ssh"],
    "disable": []
  }
}
```

### 6. Testing Strategy

**Unit Tests (Fast - <1s):**
- Pure functions (validators, parsers, config generators)
- No external dependencies
- Mock AI responses

**Integration Tests (Medium - 10-30s):**
- Docker container builds
- AI API calls (with rate limiting)
- Config generation → validation

**Acceptance Tests (Slow - 5-15min):**
- Full conversation → install flow
- VM creation + Ubuntu install
- Post-install verification
- Run in CI but cacheable

**Test Pyramid:**
```
        /\      10 Acceptance Tests (ATDD)
       /  \
      /    \    50 Integration Tests
     /      \
    /________\  200 Unit Tests
```

### 7. CI/CD Pipeline

**On every push to main:**

```yaml
1. Lint (ruff, mypy)          # 10s
2. Security scan (bandit)     # 10s
3. Unit tests                 # 30s
4. Integration tests          # 2min
5. Build Docker image         # 1min
6. Push to Docker Hub         # 30s
7. Acceptance tests (sample)  # 5min
8. Tag release                # 5s

Total: ~9 minutes
```

**On pull requests:**
- Steps 1-4 only (fast feedback)
- Full pipeline on approval

**On tags (v*.*.*):**
- Full pipeline
- Create GitHub release
- Update documentation

### 8. Docker Image

**Base:** `alpine:latest` (or `scratch` for minimal)

**Includes:**
- Single Go binary (statically compiled)
- QEMU/KVM for VM testing
- Ubuntu installation tools
- Pre-cached Ubuntu ISO (optional)

**Size:** ~20MB (vs ~150MB for Python)

**Usage:**
```bash
# Run installer
docker run -it --privileged \
  -e ANTHROPIC_API_KEY=sk-... \
  ghcr.io/gauthierbraillon/ai-os-installer:latest

# Or with local config
docker run -it --privileged \
  -v $(pwd)/config.json:/app/config.json \
  ghcr.io/gauthierbraillon/ai-os-installer:latest --config /app/config.json

# Or download standalone binary (no Docker)
curl -L https://github.com/gauthierbraillon/ai-os-installer/releases/latest/download/ai-os-installer-linux-amd64 -o ai-os-installer
chmod +x ai-os-installer
./ai-os-installer
```

### 9. Security Considerations

**API Keys:**
- Never commit API keys
- Use environment variables
- GitHub secrets for CI/CD

**VM Isolation:**
- Run VMs with minimal privileges
- Network isolation (no external access during install)
- Cleanup after tests

**Input Validation:**
- Sanitize all user inputs
- Validate config before execution
- Prevent injection attacks

**Secrets in Config:**
- Hash passwords before storing
- Never log sensitive data
- Clear secrets from memory after use

### 10. Development Workflow

**ATDD Cycle:**
```
1. Write acceptance test (RED)
   - Describes user story
   - Fails initially

2. Write integration tests (RED)
   - Components working together
   - Fails initially

3. Write unit tests (RED)
   - Pure functions
   - Fails initially

4. Implement feature (GREEN)
   - Make all tests pass
   - Minimum code needed

5. Refactor (REFACTOR)
   - Clean up code
   - Keep tests green

6. Push to main (DEPLOY)
   - CI/CD runs full pipeline
   - Auto-deploys to Docker Hub
```

### 11. Metrics & Observability

**Log everything:**
- Structured JSON logs
- Conversation turns
- API calls (redact keys)
- Installation steps
- Errors with context

**Metrics to track:**
- Conversation length (avg turns)
- Installation success rate
- Time to install
- API call count/cost
- Error types and frequency

**Export:**
- stdout for Docker logs
- Optional: Datadog/Grafana integration

### 12. Future Architecture

**Post-MVP additions:**

**Offline Mode:**
- Embed TinyLlama/Phi-3
- Fallback when no internet
- Lower quality but functional

**Bootable USB:**
- Generate ISO with installer
- Boot from USB → conversation → install
- No Docker needed

**Multi-Distro:**
- Abstract install executor
- Distro-specific modules
- Shared conversation layer

**Plugin System:**
- Community-contributed configs
- Pre-built profiles (gaming, ML, etc.)
- Marketplace (future)

## Technical Constraints

**Must have:**
- Python 3.11+ (for match/case, typing improvements)
- Docker with --privileged (for VM testing)
- 4GB RAM minimum (for VM)
- Internet connection (for AI API)

**Nice to have:**
- KVM acceleration (faster VMs)
- SSD (faster ISO handling)
- GitHub account (for contributing)

## Open Questions

1. Should we support headless install (no terminal, config-only)?
2. How to handle installation failures gracefully?
3. Should we cache AI responses for common questions?
4. Rate limiting strategy for AI API calls?
5. How to version install configs for backwards compatibility?

## References

- Ubuntu Preseed: https://help.ubuntu.com/lts/installation-guide/amd64/apb.html
- Cloud-Init: https://cloudinit.readthedocs.io/
- QEMU: https://www.qemu.org/docs/master/
- Claude API: https://docs.anthropic.com/
- Minimum CD: https://minimumcd.org/

## Decision Log

**2026-02-16:** Initial architecture
- Chose Python for familiarity
- Chose Docker for easy distribution
- Chose Ubuntu as first target (most users)
- Chose Claude API for best conversation quality
- Chose ATDD for quality and documentation
