# AI OS Installer

> **Conversational AI-guided Linux installer**

[![CI/CD](https://github.com/gauthierbraillon/ai-os-installer/actions/workflows/ci-cd.yml/badge.svg)](https://github.com/gauthierbraillon/ai-os-installer/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Vision

Installing Linux shouldn't require reading documentation or navigating complex menus. Describe what you want, and the AI guides you through every choice.

## Example

```
$ docker run -it ai-os-installer

AI: Hi! I'll help you install Ubuntu. What will you mainly use it for?

You: Web development

AI: Great! I'll set up Node, Docker, and VS Code.
    Want a simple setup or detailed customization?

You: Simple

AI: Perfect. Installing Ubuntu now...
```

## Status

⚠️ **MVP in development** - Not ready for production use yet.

### Roadmap

- [x] Architecture design
- [x] Development workflow
- [ ] CD pipeline setup
- [ ] Conversation engine
- [ ] Ubuntu installer integration
- [ ] First working demo

## Installation

Once MVP is ready:

```bash
docker run -it --privileged \
  -e ANTHROPIC_API_KEY=your_key_here \
  ghcr.io/gauthierbraillon/ai-os-installer:latest
```

## Architecture

```
User Terminal
     ↓
Conversation Manager
     ↓
AI Engine (Claude API)
     ↓
Config Generator
     ↓
Install Executor
```

**Tech Stack:** Go, Claude API, QEMU, Docker, Ubuntu 24.04 LTS

See [ARCHITECTURE.md](./ARCHITECTURE.md) for details.

## Testing

Following ATDD (Acceptance Test-Driven Development):

```bash
go test -v ./...
```

See [CLAUDE.md](./CLAUDE.md) for development workflow.

## Contributing

Contributions welcome! This project follows:
- ATDD/TDD workflow
- Continuous delivery
- Trunk-based development

See [CLAUDE.md](./CLAUDE.md) for guidelines.

## Documentation

- [CLAUDE.md](./CLAUDE.md) - Development workflow
- [ARCHITECTURE.md](./ARCHITECTURE.md) - Technical architecture

## Issues & Support

Found a bug or have a question? [Open an issue](https://github.com/gauthierbraillon/ai-os-installer/issues)

For security vulnerabilities, please open a private security advisory.

## License

MIT

## Author

Gauthier Braillon - [@gauthierbraillon](https://github.com/gauthierbraillon)

---

⭐ Star this repo if you think conversational OS installation is the future!
