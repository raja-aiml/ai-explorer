package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// -------- Tests --------

func TestPromptCommand_Metadata(t *testing.T) {
	cmd := GetPromptCommand()
	assert.Equal(t, "prompt", cmd.Use)
	assert.Contains(t, cmd.Short, "Generate prompt")
}

func TestPromptCommand_FlagsRegistered(t *testing.T) {
	cmd := GetPromptCommand()
	flags := cmd.Flags()

	assert.NotNil(t, flags.Lookup("category"))
	assert.NotNil(t, flags.Lookup("topic"))
	assert.NotNil(t, flags.Lookup("template"))
	assert.NotNil(t, flags.Lookup("config"))
	assert.NotNil(t, flags.Lookup("output"))
	assert.NotNil(t, flags.Lookup("preview"))
}
