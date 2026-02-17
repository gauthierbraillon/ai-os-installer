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

⚠️ **MVP in development** — not ready for production use yet.

## Installation

```bash
docker run -it --privileged \
  -e ANTHROPIC_API_KEY=<your-key> \
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

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md) for the development workflow and contribution guide.

## License

MIT
