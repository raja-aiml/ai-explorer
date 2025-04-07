package prompt

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// --- Tests: Public rendering interface ---

func Test_Builder_RenderToFile_WritesExpectedOutput(t *testing.T) {
	dir := t.TempDir()
	tmpl := writeTempFile(t, dir, "template.tmpl", `Hello {{ name }}`)
	cfg := writeTempFile(t, dir, "config.yaml", `name: world`)
	out := filepath.Join(dir, "output.txt")

	var wrote []byte
	builder := &Builder{
		ReadFile:  os.ReadFile,
		WriteFile: func(_ string, data []byte, _ os.FileMode) error { wrote = data; return nil },
		Logger:    &fakeLogger{},
	}

	builder.RenderToFile(tmpl, cfg, out)
	assert.Contains(t, string(wrote), "Hello world")
}

func Test_Builder_RenderToStdout_PrintsExpectedOutput(t *testing.T) {
	dir := t.TempDir()
	tmpl := writeTempFile(t, dir, "template.tmpl", `Hi {{ name }}`)
	cfg := writeTempFile(t, dir, "config.yaml", `name: stdout`)

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	builder := &Builder{
		ReadFile:  os.ReadFile,
		WriteFile: nil,
		Logger:    &fakeLogger{},
	}
	builder.RenderToStdout(tmpl, cfg)

	_ = w.Close()
	os.Stdout = oldStdout

	out, _ := io.ReadAll(r)
	assert.Contains(t, string(out), "Hi stdout")
}
