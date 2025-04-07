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


-------

To Do 

Thanks for the detailed dump—your project is looking excellent! It’s already modular, cleanly structured, and tightly tested with both unit and E2E coverage.

Since there wasn’t a direct question in your last message, here’s a checklist of refinement and polish suggestions as you’re prepping for a stable release:

⸻

✅ 1. CLI Experience (UX) Polish
	•	✅ You already have completion for bash, zsh, fish, and powershell.
	•	📌 Add --help examples for the llm command similar to how prompt uses resources/help/examples.md.
	•	📌 Add descriptions or defaults to printed CLI flag values via cmd/llm/llm.go (using cmd.Flags().VisitAll() + SetUsageTemplate()).

⸻

✅ 2. Project Cleanup / Consistency
	•	🔍 You have a resources/paths.go and a separate paths/manager.go. Consider merging these into a unified paths package to avoid ambiguity.
	•	💡 Suggest: move PathResolver into paths (renaming resources/paths.go to paths/resolver.go) to align directory and purpose.

⸻

✅ 3. Observability or Debug Mode
	•	💬 Add a global --verbose or --debug flag to:
	•	Print intermediate paths used (template, config, output)
	•	Print which model/provider is being used
	•	📁 Suggest using your logger interface for this—replace log.Fatalf in LLMRunner with logger.Fatalf to decouple from hardcoded logging.

⸻

✅ 4. Output Paths for llm and prompt
	•	🧼 Paths like resources/classification/router/prompt.txt and llm.txt are useful but maybe too hardcoded.
	•	📌 Let the CLI auto-pick a timestamped name like:

output = fmt.Sprintf("resources/classification/router/%s.llm.txt", time.Now().Format("20060102-150405"))

…if the --output flag is not provided.

⸻

✅ 5. Release Tasks

Consider adding the following to your Taskfile.yaml:

release:
  desc: "Build binary and tag a release"
  cmds:
    - task: build
    - git tag v1.0.0
    - git push origin v1.0.0

And for completions on install:

eval "$(.build/ai-explorer completion zsh)"



⸻

✅ 6. Possible Nice-to-Have
	•	🧠 Add a --dry-run to llm that only prints what prompt would be sent (useful for debugging prompt logic).
	•	📦 Add a config.yaml flag to llm to allow specifying all config from YAML (reusing llm/config.ConfigLoader).

⸻

Let me know if you want a code patch for anything above (like swapping log.Fatalf in LLMRunner for your logger, or merging the paths packages cleanly). Also happy to help build a real v1.0.0 changelog and release checklist.