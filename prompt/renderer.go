package prompt

import (
	"os"

	"raja.aiml/ai.explorer/logger"
)

// Renderer defines rendering methods for different output targets.
type Renderer interface {
	RenderToFile(templatePath, configPath, outputPath string, userQuery ...string)
	RenderToStdout(templatePath, configPath string, userQuery ...string)
}

// DefaultRenderer is the standard implementation of Renderer.
var DefaultRenderer Renderer = &Builder{
	ReadFile:  os.ReadFile,
	WriteFile: os.WriteFile,
	Logger:    logger.New(),
}
