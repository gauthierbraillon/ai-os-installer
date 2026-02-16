# AI OS Installer - Project Guide

## Product Vision

AI OS Installer makes Linux installation as easy as talking to an expert friend. Instead of navigating complex menus and documentation, users describe what they want and the AI guides them through every choice.

### MVP Scope (Weeks 1-3)

**Target:** Ubuntu 24.04 LTS only
**Environment:** VM testing (QEMU/VirtualBox)
**Distribution:** Docker container
**AI:** Cloud API (Claude/OpenAI)
**Goal:** Conversational installer that works end-to-end in a VM

### Core User Experience

```
User starts installer
  ↓
AI: "Hi! Let's set up your Linux system. What will you mainly use this for?"
  ↓
User: "Web development and some gaming"
  ↓
AI: "Great! For web dev, I'll set up Node, Docker, and VS Code.
     For gaming, I'll install Steam and configure drivers.
     Want a simple setup or detailed customization?"
  ↓
User: "Simple"
  ↓
AI: "Perfect. Installing now..."
  [Progress updates]
  ↓
Done! Ubuntu ready with your preferences
```

### Out of Scope (Post-MVP)

- Other distros (Arch, Fedora)
- Bootable USB creation
- Offline mode with local LLM
- GUI interface
- Multi-language support

## Testing Philosophy

### ATDD (Acceptance Test Driven Development)

**ATDD IS TDD** - same RED → GREEN → REFACTOR cycle, different test focus.

**The ONLY difference:**
- **ATDD tests** describe WHAT the system should do (acceptance criteria, observable behavior)
- **Classic TDD tests** often describe HOW the system works (class structure, object interactions)

**ATDD workflow (same as TDD):**
1. **RED**: Write acceptance test first (test IS the requirement)
2. **GREEN**: Implement minimum code to pass
3. **REFACTOR**: Clean up code and tests
4. **DEPLOY**: Ship it

**No separate requirements document. No waterfall. Just TDD with acceptance-focused tests.**

### Test Types (Development Focus)

**PRIMARY: Sociable Unit Tests (Fast Feedback During Development)**
- **Tests ARE the acceptance criteria and testable requirements** (no separate requirements doc)
- Test full flow through OUR code/modules
- Mock external components only (Claude API, QEMU, file system)
- No mocks of our own code - use real collaborators
- Fast execution (< 2 seconds)
- Run constantly during RED → GREEN → REFACTOR cycle
- Test names describe requirements (e.g., "AC101: User can describe use case and get relevant packages")
- Examples: conversation flow with config generation, input validation with AI parsing, config generation with package selection

**SECONDARY: Integration Tests (Deployment Pipeline)**
- Test integration with real external systems
- Claude API tests (real API calls with rate limiting)
- Docker container tests (actual container build and run)
- QEMU tests (real VM creation and install)
- Slower execution (30-60 seconds)
- Run in deployment pipeline before deploy
- Verify our code works with actual dependencies

**TERTIARY: E2E Smoke Tests (Production Verification + Rollback Trigger)**
- Verify production deployment works end-to-end
- Run against actual Docker image from Docker Hub
- Test full conversation → install flow in VM
- Slowest execution (5-15 minutes)
- Run AFTER deployment as smoke tests
- **CRITICAL: If E2E smoke tests fail → ROLLBACK deployment immediately**
- Examples: Full Ubuntu install test, config validation test, post-install verification

### Test Organization

```
tests/
├── unit/                           # Sociable unit tests (acceptance criteria)
│   ├── conversation_test.py        # Conversation manager tests
│   ├── ai_client_test.py          # AI client tests (mocked API)
│   ├── config_generator_test.py   # Config generation tests
│   └── validator_test.py          # Input validation tests
├── integration/                    # Real external systems
│   ├── docker_test.py             # Docker container tests
│   ├── claude_api_test.py         # Real Claude API calls
│   └── qemu_test.py               # Real VM tests
├── e2e/                           # Production smoke tests
│   └── full_install_test.py      # Full conversation → install flow
└── validation/                     # Static checks
    ├── validate_structure.py      # File structure validation
    └── validate_security.py       # Security checks
```

### What to Mock

- ✅ External APIs (Claude API, OpenAI)
- ✅ File system operations (during unit tests)
- ✅ VM operations (during unit tests)
- ✅ Time/dates (for deterministic tests)
- ❌ Our own code
- ❌ Simple utilities

### Test Pyramid Rules

**Test at lowest level possible. No duplicate coverage.**
- Unit test? → Don't test again at integration level
- Integration test? → Don't test again at E2E level
- Fast feedback: unit (<2s) > integration (~60s) > E2E (~15min)

**Static checks don't validate requirements.**
- Linting checks code style, not behavior
- Security scans find vulnerabilities, not functional bugs
- Real tests verify the system does what users need

## Multi-Perspective Analysis

For new features or significant changes, analyze from multiple angles before implementing:
- **Requirements**: What problem does this solve? What are the acceptance criteria?
- **User Experience**: How will users interact with this? Is it intuitive for Linux beginners?
- **Technical Design**: What's the simplest implementation approach?
- **Security**: What are the risks? How do we mitigate them?

Present a synthesis, then implement using the TDD workflow.

## Development Workflow

### Continuous Delivery Principles

Following [Minimum CD](https://minimumcd.org/) practices:

1. **Small batches**: Deploy after each feature/fix
2. **Fast feedback**: All tests run in < 10 minutes
3. **Green to deploy**: If tests pass, code is deployable
4. **No manual gates**: Automation decides quality
5. **Trunk-based**: Direct commits to main (or short-lived branches)

### Deployment Pipeline

**On every push to main:**

```bash
# Automated via GitHub Actions
1. Lint (ruff, mypy)                    # 10s
2. Security scan (bandit)               # 10s
3. Unit tests                           # 30s
4. Integration tests                    # 2min
5. Build Docker image                   # 1min
6. Push to Docker Hub                   # 30s
7. Tag release (semantic)               # 5s
8. E2E smoke tests (sample)             # 5min

Total: ~9 minutes
```

**Gates (must all pass):**
- Pre-deploy (1-6): If any fail, deployment aborts
- Post-deploy smoke tests (8): If fail, ROLLBACK immediately

### Writing Code

**🚨 CRITICAL: ALWAYS Follow the TDD Cycle (RED → GREEN → REFACTOR → DEPLOY) 🚨**

**This workflow is MANDATORY for ALL code changes - no exceptions:**
- ✅ Bug fixes → Follow workflow
- ✅ New features → Follow workflow
- ✅ Refactoring → Follow workflow
- ✅ Configuration changes → Follow workflow
- ✅ Documentation updates → May skip if purely textual, but consider validation tests

**If you skip the workflow, you're doing it wrong.** The workflow ensures quality, prevents bugs, and maintains continuous delivery discipline.

**TDD Cycle (RED → GREEN → REFACTOR → DEPLOY):**

1. **RED phase** - Write tests that ARE the acceptance criteria

   - **Step 1: Understand the problem** - What observable behavior needs to change?
     - For simple bugs: understand the incorrect behavior
     - For complex changes: break down into smaller testable behaviors
     - Think in terms of "Given-When-Then" scenarios

   - **Step 2: Write sociable unit tests** - Tests ARE the acceptance criteria and requirements
     - Test names describe the requirement (e.g., "AC101: User can describe use case and get relevant packages")
     - Test full flow through our code with mocked externals
     - Mock external systems only (Claude API, QEMU, file system)
     - Keep tests fast (< 2 seconds)
     - Tests ARE the specification - no separate requirements document needed

   - **Step 3: Run test to verify RED** - Confirm test fails for the right reason

2. **GREEN phase** - Make test pass (fast feedback)

   - Implement minimum code to pass the test
   - Run ONLY the failing test (fast feedback loop)
   - Iterate quickly until test passes
   - Don't refactor yet - just make it work

3. **REFACTOR phase** - Clean up code AND tests

   - **CRITICAL: ACTUALLY REFACTOR THE CODE** - Don't just run tests and move on
   - **Look for code smells**:
     - Duplication (repeated logic, magic strings, duplicate conditions)
     - Poor naming (unclear variable/function names, inconsistent terminology)
     - Long functions (extract helper functions for clarity)
     - Hidden dependencies (calculate values once, not repeatedly)
     - Magic values (extract to named constants)
   - **Refactor the implementation**: Extract functions, remove duplication, improve naming
   - **Refactor tests if needed**: Improve test readability, reduce duplication
   - **Run ALL tests** (unit + integration) to ensure refactoring didn't break anything
   - Tests stay sociable - no new mocks of our own code
   - **Example refactoring**:
     - Before: `if use_case == "web_dev": packages = ["nodejs", "docker", ...]` appears twice
     - After: `get_packages_for_use_case(use_case: str) -> List[str]`

4. **DEPLOY phase** - Full pipeline

   - **CRITICAL: Commit in small chunks** - After each logical change (not at the end of the day)
   - Each commit should be a single, atomic change that could be reverted independently
   - Use conventional commit messages: `feat:`, `fix:`, `refactor:`, `test:`, `docs:`
   - **Commit examples:**
     - ✅ `feat: add Installer interface`
     - ✅ `feat: implement UbuntuInstaller.Validate`
     - ✅ `test: add validation tests for UbuntuInstaller`
     - ❌ Don't: `feat: add entire installer module` (too big)
   - Push to main triggers GitHub Actions automatically
   - **Smoke tests pass**: Deployment successful
   - **Smoke tests fail**: ROLLBACK immediately
   - Update MEMORY.md with lessons learned

**Why small commits matter:**
- Easier to review
- Easier to debug (git bisect)
- Easier to revert if needed
- Shows progress continuously
- Enables true continuous delivery

**Why This Workflow is NON-NEGOTIABLE:**
- **Prevents bugs**: Tests catch issues before they reach production
- **Documents behavior**: Tests serve as executable specifications
- **Enables refactoring**: Safe to improve code when tests verify behavior
- **Maintains quality**: Every commit is tested and deployable
- **Fast feedback**: Catch issues in seconds, not hours or days
- **Continuous delivery**: Small, safe, frequent deployments

**Code Style:**
- No comments (code should be self-explanatory)
- Functional programming preferred (pure functions, immutable data)
- Interfaces for testability (dependency injection)
- No unnecessary abstractions
- Follow standard Go conventions (gofmt, golangci-lint)
- Table-driven tests where appropriate

**Documentation Philosophy:**
- Keep docs minimal - every line needs maintenance
- Avoid hard data (sizes, timings) that can become outdated
- Don't duplicate information across files
- CLAUDE.md is single source of truth for development workflow
- README is for users, not developers
- Comments in code only when logic is truly unclear (rare)

## Architecture

### System Components

```
User Terminal
     ↓
Conversation Manager (collects requirements, manages state)
     ↓
AI Engine (Claude API, prompt engineering, response parsing)
     ↓
Config Generator (conversation → install config)
     ↓
Install Executor (runs installation in VM)
```

### File Structure

```
src/
├── conversation/
│   ├── manager.py              # Main conversation loop
│   ├── state.py                # Conversation state management
│   └── validator.py            # User input validation
├── ai/
│   ├── client.py               # AI API client (Claude/OpenAI)
│   ├── prompts.py              # Prompt templates
│   └── parser.py               # Response parsing
├── config/
│   ├── generator.py            # Conversation → install config
│   ├── validator.py            # Config validation
│   └── templates/              # Preseed/cloud-init templates
├── install/
│   ├── executor.py             # VM installation orchestrator
│   ├── vm_manager.py           # VM lifecycle
│   └── progress.py             # Installation progress tracking
└── utils/
    ├── logger.py               # Structured logging
    ├── errors.py               # Custom exceptions
    └── validation.py           # Common validators

tests/
├── unit/                       # Fast sociable unit tests
├── integration/                # External system integration
├── e2e/                        # Production smoke tests
└── validation/                 # Static checks

docker/
├── Dockerfile                  # Production image
└── docker-compose.yml          # Local development

.github/
└── workflows/
    └── ci-cd.yml               # GitHub Actions pipeline
```

### Key Concepts

**Conversation State**
- Track user responses
- Build requirements incrementally
- Allow backtracking ("go back")
- Validate responses before proceeding

**Config Format (JSON)**
```json
{
  "version": "1.0",
  "user": {
    "name": "john",
    "full_name": "John Doe",
    "shell": "bash"
  },
  "system": {
    "hostname": "dev-machine",
    "timezone": "Europe/Brussels"
  },
  "packages": {
    "base": ["git", "curl", "vim"],
    "dev": ["python3", "nodejs", "docker"],
    "desktop": ["gnome-shell"]
  }
}
```

**VM Installation Flow**
1. Generate preseed/cloud-init config from conversation
2. Create VM with QEMU
3. Mount Ubuntu ISO
4. Run automated install
5. Post-install scripts (packages, config)
6. Verify installation
7. Export VM or provide access

## Common Tasks

**Add new use case (e.g., "gaming" setup):**
1. Add test in `tests/unit/config_generator_test.py` (RED)
2. Update `src/config/generator.py` to handle new use case (GREEN)
3. Refactor if duplicated logic (REFACTOR)
4. Commit and push (DEPLOY)

**Improve conversation flow:**
1. Add test in `tests/unit/conversation_test.py` (RED)
2. Update `src/conversation/manager.py` (GREEN)
3. Refactor conversation logic (REFACTOR)
4. Commit and push (DEPLOY)

**Fix production bug:**
1. **Reproduce** - Verify bug exists (run E2E test or manual test)
2. **Test** - Write test that catches bug (should fail - RED)
3. **Fix** - Minimum code to pass test (GREEN)
4. **Refactor** - Clean up if needed (REFACTOR)
5. **Deploy** - Push and verify fix in production (DEPLOY)
6. **Learn** - Update MEMORY.md with lesson

**Add support for new distro:**
1. Design abstraction for distro-specific logic
2. Write tests for new distro (RED)
3. Implement distro module (GREEN)
4. Refactor to reduce duplication (REFACTOR)
5. Update docs and deploy (DEPLOY)

## Deployment Targets

**Development:** Local Docker container
- Run: `docker run -it --privileged -e ANTHROPIC_API_KEY=... ai-os-installer:dev`
- Fast iteration, no push needed

**Production:** Docker Hub
- Image: `ghcr.io/gauthierbraillon/ai-os-installer:latest`
- Public, anyone can pull and run
- Automated builds on every push to main

**Usage:**
```bash
# Pull and run
docker run -it --privileged \
  -e ANTHROPIC_API_KEY=sk-... \
  ghcr.io/gauthierbraillon/ai-os-installer:latest

# With custom config
docker run -it --privileged \
  -v $(pwd)/config.json:/app/config.json \
  ghcr.io/gauthierbraillon/ai-os-installer:latest --config /app/config.json
```

## Security Reminders

**Never commit:**
- API keys (ANTHROPIC_API_KEY, OPENAI_API_KEY)
- Test credentials
- Private configs

**Always:**
- Use environment variables for secrets
- Validate all user inputs (prevent injection)
- Sanitize error messages (no sensitive data in logs)
- Run security scans (bandit, safety)
- Keep dependencies updated
- Run VMs with minimal privileges

**API Key Handling:**
- Store in environment variables
- GitHub Actions: Use secrets
- Never log API keys
- Rotate keys regularly

## References

- [Minimum Continuous Delivery](https://minimumcd.org/)
- [Ubuntu Preseed](https://help.ubuntu.com/lts/installation-guide/amd64/apb.html)
- [Cloud-Init](https://cloudinit.readthedocs.io/)
- [QEMU](https://www.qemu.org/docs/master/)
- [Claude API](https://docs.anthropic.com/)
- [Docker Best Practices](https://docs.docker.com/develop/dev-best-practices/)
