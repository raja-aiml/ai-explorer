package llm

import "time"

// Default file paths
const (
	DefaultConfigPath  = "resources/demo/hello/config.yaml" // Default config file
	DefaultProvider    = "ollama"
	DefaultModel       = "phi4"
	DefaultTemperature = 0.8
	DefaultTimeout     = 2 * time.Minute
	DefaultPromptPath  = "resources/demo/hello/prompt.txt" // Ensure a valid default prompt path
)

// CLI flags
var (
	providerName string
	modelName    string
	temperature  float64
	promptPath   string
	timeout      time.Duration
	outputPath   string
)
