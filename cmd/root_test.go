package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestExecuteSuccess(t *testing.T) {
	// Save and restore global rootCmd
	origRoot := rootCmd
	defer func() { rootCmd = origRoot }()

	rootCmd = &cobra.Command{
		Use: "ai-explorer",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(), "CLI ran successfully")
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)

	Execute()
	assert.Contains(t, buf.String(), "CLI ran successfully")
}

func TestExecuteFailure(t *testing.T) {
	// Save and restore globals
	origRoot := rootCmd
	origExit := exit
	defer func() {
		rootCmd = origRoot
		exit = origExit
	}()

	// Mock os.Exit
	var code int
	exit = func(c int) {
		code = c
	}

	rootCmd = &cobra.Command{
		Use: "ai-explorer",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("something went wrong")
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)

	Execute()
	assert.Contains(t, buf.String(), "Error: something went wrong")
	assert.Equal(t, 1, code)
}
