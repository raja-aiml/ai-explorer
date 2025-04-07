package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func newRootForTest() *cobra.Command {
	root := &cobra.Command{Use: "ai-explorer"}
	root.AddCommand(completionCmd)
	return root
}

func TestCompletionCommand_Metadata(t *testing.T) {
	assert.Equal(t, "completion [bash|zsh|fish|powershell]", completionCmd.Use)
	assert.Contains(t, completionCmd.Short, "Generate shell completion")
	assert.ElementsMatch(t, []string{"bash", "zsh", "fish", "powershell"}, completionCmd.ValidArgs)
}

func TestCompletionCommand_ValidArgs(t *testing.T) {
	shells := []string{"bash", "zsh", "fish", "powershell"}

	for _, shell := range shells {
		t.Run("Shell="+shell, func(t *testing.T) {
			root := newRootForTest()
			root.SetArgs([]string{"completion", shell})

			r, w, _ := os.Pipe()
			stdout := os.Stdout
			os.Stdout = w

			err := root.Execute()
			_ = w.Close()
			os.Stdout = stdout

			var out bytes.Buffer
			_, _ = out.ReadFrom(r)

			assert.NoError(t, err)
			assert.Contains(t, out.String(), "completion", "Expected output for shell: %s", shell)
		})
	}
}

func TestCompletionCommand_InvalidArgs(t *testing.T) {
	root := newRootForTest()
	root.SetArgs([]string{"completion", "invalidshell"})

	var stderr bytes.Buffer
	root.SetErr(&stderr)

	err := root.Execute()
	assert.Error(t, err)
	assert.Contains(t, stderr.String(), "invalid argument")
}
