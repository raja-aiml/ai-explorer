# AI Explorer

> Prompt generation + LLM interaction CLI tool powered by Go, Cobra, and Pongo2.

---

## ✨ Features

- 🧠 **Prompt Generator**: Generate prompts using YAML config + template.
- 🤖 **LLM CLI**: Send raw prompts to LLMs (e.g., OpenAI, Ollama).
- 🛠️ **Template System**: Uses `pongo2` (Jinja-like) templating.
- 📁 **Category-based routing**: Organize prompts under `resources/<category>/<topic>`.
- ✅ **100% Unit + E2E Tests**: Ginkgo-powered full coverage.
- 📜 **Shell Completion**: Autogenerate completions for bash, zsh, fish, PowerShell.

---

## 🚀 Quickstart

```bash
# Build
task build

# Generate a prompt (auto path resolution)
ai-explorer prompt --category=topics --topic=git

# Preview prompt
ai-explorer prompt --category=topics --topic=git --preview

# Send prompt to LLM
ai-explorer llm --provider=ollama --model=phi4 --prompt=resources/topics/git/prompt.txt

# Generate zsh completion script
task completion


