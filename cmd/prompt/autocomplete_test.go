package cmd

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestPromptAutoComplete(t *testing.T) {
	cmd := &cobra.Command{}
	completions, directive := promptAutoComplete(cmd, []string{}, "")

	expected := []string{"topics", "demo", "classification"}
	assert.Equal(t, expected, completions)
	assert.Equal(t, cobra.ShellCompDirectiveNoFileComp, directive)
}
