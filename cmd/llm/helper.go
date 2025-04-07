package llm

import (
	"fmt"
	"os"
)

// getPrompt reads the prompt from a file and returns its content as a string.
func getPrompt(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error reading prompt file '%s': %w", path, err)
	}
	return string(content), nil
}

// saveResponse writes the LLM response to the specified file.
func saveResponse(response, path string) error {
	return os.WriteFile(path, []byte(response), 0644)
}
