package resources

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustReadFile_Success(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "example.txt")
	expected := "Hello, AI Explorer!"

	err := os.WriteFile(filePath, []byte(expected), 0644)
	assert.NoError(t, err)

	content := MustReadFile(filePath)
	assert.Equal(t, expected, content)
}

func TestMustReadFile_PanicOnMissingFile(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic when file does not exist")
		}
	}()

	_ = MustReadFile("non-existent-file.txt")
}
