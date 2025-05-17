package llm

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"raja.aiml/ai.explorer/llm"

	llmConfig "raja.aiml/ai.explorer/llm/config"
)

// Cobra command for `llm`
var llmCmd = &cobra.Command{
	Use:   "llm",
	Short: "Send a raw prompt to LLM",
	Run: func(cmd *cobra.Command, args []string) {
		runner := &LLMRunner{
			Out:          os.Stdout,
			PromptPath:   promptPath,
			OutputPath:   outputPath,
			GetPrompt:    getPrompt,
			RunLLM:       runLLMInteraction,
			SaveResponse: saveResponse,
		}
		runner.Run()
	},
}

// Expose cobra command
func GetLLMCommand() *cobra.Command {
	return llmCmd
}

func init() {
	llmCmd.Flags().StringVarP(&providerName, "provider", "l", DefaultProvider, "LLM provider")
	llmCmd.Flags().StringVarP(&modelName, "model", "m", DefaultModel, "LLM model")
	llmCmd.Flags().Float64VarP(&temperature, "temperature", "t", DefaultTemperature, "Temperature")
	llmCmd.Flags().StringVarP(&promptPath, "prompt", "p", DefaultPromptPath, "Prompt file")
	llmCmd.Flags().DurationVarP(&timeout, "timeout", "d", DefaultTimeout, "Timeout duration")
	llmCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output file path to save LLM response")
	// Optional: override Ollama server URL if not using OLLAMA_HOST env var
	llmCmd.Flags().StringVar(&serverURL, "server-url", "", "Ollama server URL (e.g. http://localhost:11434)")
}

// runLLMInteraction initializes the LLM client and returns the response for the given prompt.
func runLLMInteraction(prompt string) (string, error) {
	cfg := llmConfig.Config{
		Provider: providerName,
		Model: llmConfig.ModelConfig{
			Name:        modelName,
			Temperature: temperature,
		},
		Client: llmConfig.ClientConfig{
			Timeout:        timeout,
			VerboseLogging: true,
		},
	}

	// If using Ollama, ensure a host is configured via env or flag
	if providerName == "ollama" && os.Getenv("OLLAMA_HOST") == "" && serverURL == "" {
		return "", fmt.Errorf("ollama selected but neither OLLAMA_HOST nor --server-url provided")
	}
	// Override OLLAMA_HOST env var if server-url flag is set
	if serverURL != "" {
		os.Setenv("OLLAMA_HOST", serverURL)
	}
	client, err := llm.NewDefaultClient(cfg)
	if err != nil {
		return "", fmt.Errorf("failed to create LLM client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return client.Chat(ctx, prompt)
}
