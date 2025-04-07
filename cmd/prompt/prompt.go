package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"raja.aiml/ai.explorer/prompt"
)

// CLI flags
var (
	promptCategory     string
	topic              string
	promptTemplatePath string
	promptConfigPath   string
	promptOutputPath   string
	preview            bool
	userQuery          string
)

const (
	defaultPromptCategory = "topics"
	defaultTopic          = "git"
)

// Cobra CLI command for `prompt`
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Generate prompt from a category (folder), topic, and config YAML",
	Run: func(cmd *cobra.Command, args []string) {
		runner := &PromptRunner{
			Out:            cmd.OutOrStdout(),
			Renderer:       prompt.DefaultRenderer,
			PromptCategory: promptCategory,
			Topic:          topic,
			Template:       promptTemplatePath,
			Config:         promptConfigPath,
			Output:         promptOutputPath,
			Preview:        preview,
			UserQuery:      userQuery,
		}
		runner.Run()
	},
	ValidArgsFunction: promptAutoComplete,
}

// promptAutoComplete suggests categories or topics for CLI completions.
func promptAutoComplete(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// Suggest common categories or hardcoded ones for now
	return []string{"topics", "demo", "classification"}, cobra.ShellCompDirectiveNoFileComp
}

// GetPromptCommand exposes the `prompt` Cobra command.
func GetPromptCommand() *cobra.Command {
	return promptCmd
}

func init() {
	// Load usage examples from file if available
	if examples, err := os.ReadFile("resources/help/examples.md"); err == nil {
		promptCmd.Example = string(examples)
	}

	// Register CLI flags
	promptCmd.Flags().StringVar(&promptCategory, "category", "", "Base folder for prompt templates (default: topics)")
	promptCmd.Flags().StringVar(&topic, "topic", "", "Topic name (used to infer default paths)")
	promptCmd.Flags().StringVarP(&promptTemplatePath, "template", "t", "", "Path to template YAML")
	promptCmd.Flags().StringVarP(&promptConfigPath, "config", "c", "", "Path to config YAML")
	promptCmd.Flags().StringVarP(&promptOutputPath, "output", "o", "", "Path to output file")
	promptCmd.Flags().BoolVar(&preview, "preview", false, "Print output to stdout instead of writing to file")
	promptCmd.Flags().StringVarP(&userQuery, "query", "q", "", "User query to inject into template context")
}
