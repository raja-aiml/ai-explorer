package resources

import "fmt"

// PathResolver defines how paths are derived for prompts.
type PathResolver struct {
	PromptCategory string
}

// Derive returns template, config, and output file paths.
func (r PathResolver) Derive(topic string) (template, config, output string) {
	base := fmt.Sprintf("resources/%s/%s", r.PromptCategory, topic)

	if r.PromptCategory == "topics" {
		// Template is shared across topics
		template = "resources/topics/template.yaml"
	} else {
		template = fmt.Sprintf("%s/template.yaml", base)
	}

	config = fmt.Sprintf("%s/config.yaml", base)
	output = fmt.Sprintf("%s/prompt.txt", base)

	return
}
