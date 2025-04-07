package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"raja.aiml/ai.explorer/cmd/llm"
	cmd "raja.aiml/ai.explorer/cmd/prompt"
)

var rootCmd = newRootCmd()
var exit = os.Exit // overridable for testing

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:           "ai-explorer",
		Short:         "Prompt generation + LLM interaction CLI",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
}

func Execute() {
	_ = godotenv.Load()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(rootCmd.ErrOrStderr(), "Error: %v\n", err)
		exit(1)
	}
}

func init() {
	rootCmd.AddCommand(completionCmd)
	rootCmd.AddCommand(cmd.GetPromptCommand())
	rootCmd.AddCommand(llm.GetLLMCommand())
}
