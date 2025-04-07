package logger_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"raja.aiml/ai.explorer/logger"
)

func TestNewLogger_Printf(t *testing.T) {
	log := logger.New()

	assert.NotPanics(t, func() {
		log.Printf("Hello %s", "world") // Should print safely
	})
}

func TestNewLogger_Fatalf_Exits(t *testing.T) {
	if os.Getenv("LOGGER_FATAL_TEST") == "1" {
		log := logger.New()
		log.Fatalf("Fatal error test")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestNewLogger_Fatalf_Exits")
	cmd.Env = append(os.Environ(), "LOGGER_FATAL_TEST=1")

	err := cmd.Run()
	if exitErr, ok := err.(*exec.ExitError); ok {
		assert.Equal(t, 1, exitErr.ExitCode(), "expected exit code 1 from Fatalf")
	} else {
		t.Fatalf("expected exit error, got: %v", err)
	}
}

func Test_defaultLogger(t *testing.T) {
	t.Run("Printf does not panic", func(t *testing.T) {
		logger := logger.New()
		assert.NotPanics(t, func() {
			logger.Printf("Printed %d", 123)
		})
	})

	t.Run("Fatalf causes os.Exit", func(t *testing.T) {
		if os.Getenv("FATALF_TEST") == "1" {
			logger := logger.New()
			logger.Fatalf("should exit")
			return
		}

		cmd := exec.Command(os.Args[0], "-test.run=Test_defaultLogger")
		cmd.Env = append(os.Environ(), "FATALF_TEST=1")

		err := cmd.Run()
		if exitErr, ok := err.(*exec.ExitError); ok {
			assert.Equal(t, 1, exitErr.ExitCode(), "expected exit code 1 from Fatalf")
		} else {
			t.Fatalf("expected exit error, got: %v", err)
		}
	})
}
