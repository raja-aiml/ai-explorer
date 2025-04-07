package prompt

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/flosch/pongo2/v6"
	"github.com/stretchr/testify/assert"
)

// --- Test Helpers ---

type fakeLogger struct {
	FatalMsg string
	PrintLog []string
}

func (f *fakeLogger) Printf(format string, v ...any) {
	f.PrintLog = append(f.PrintLog, fmt.Sprintf(format, v...))
}

func (f *fakeLogger) Fatalf(format string, v ...any) {
	f.FatalMsg = fmt.Sprintf(format, v...)
	panic("fakeLogger.Fatalf called: " + f.FatalMsg)
}

func writeTempFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	full := filepath.Join(dir, name)
	assert.NoError(t, os.WriteFile(full, []byte(content), 0644))
	return full
}

func assertPanics(t *testing.T, fn func(), msg string) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Error(msg)
		}
	}()
	fn()
}

// --- Tests: Internal helpers ---

func Test_Builder_mustParseTemplate_Success(t *testing.T) {
	logger := &fakeLogger{}
	builder := &Builder{
		ReadFile: func(string) ([]byte, error) { return []byte("Hi {{ name }}"), nil },
		Logger:   logger,
	}

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("unexpected panic: %v", r)
		}
	}()

	tpl := builder.mustParseTemplate("ok.tmpl")
	out, err := tpl.Execute(pongo2.Context{"name": "Go"})
	assert.NoError(t, err)
	assert.Equal(t, "Hi Go", out)
}

func Test_Builder_mustParseTemplate_ReadError(t *testing.T) {
	builder := &Builder{
		ReadFile: func(string) ([]byte, error) { return nil, errors.New("read error") },
		Logger:   &fakeLogger{},
	}

	assertPanics(t, func() {
		builder.mustParseTemplate("bad.txt")
	}, "expected panic on template read error")
}

func Test_Builder_mustParseTemplate_ParseError(t *testing.T) {
	builder := &Builder{
		ReadFile: func(string) ([]byte, error) { return []byte("{{ broken"), nil },
		Logger:   &fakeLogger{},
	}

	assertPanics(t, func() {
		builder.mustParseTemplate("bad.txt")
	}, "expected panic on template parse error")
}

func Test_Builder_mustParseConfig_Success(t *testing.T) {
	builder := &Builder{
		ReadFile: func(string) ([]byte, error) { return []byte("key: val"), nil },
		Logger:   &fakeLogger{},
	}

	ctx := builder.mustParseConfig("ok.yaml")
	assert.Equal(t, "val", ctx["key"])
}

func Test_Builder_mustParseConfig_ReadError(t *testing.T) {
	builder := &Builder{
		ReadFile: func(string) ([]byte, error) { return nil, errors.New("read fail") },
		Logger:   &fakeLogger{},
	}

	assertPanics(t, func() {
		builder.mustParseConfig("bad.yaml")
	}, "expected panic on config read error")
}

func Test_Builder_mustParseConfig_ParseError(t *testing.T) {
	builder := &Builder{
		ReadFile: func(string) ([]byte, error) { return []byte("invalid: ["), nil },
		Logger:   &fakeLogger{},
	}

	assertPanics(t, func() {
		builder.mustParseConfig("broken.yaml")
	}, "expected panic on YAML parse error")
}

func Test_Builder_renderAndWrite_Success(t *testing.T) {
	var captured []byte
	tpl, _ := pongo2.FromString(`{{ foo }}!`)

	builder := &Builder{
		WriteFile: func(_ string, data []byte, _ os.FileMode) error {
			captured = data
			return nil
		},
		Logger: &fakeLogger{},
	}

	builder.renderAndWrite(tpl, pongo2.Context{"foo": "Yo"}, "out.txt")
	assert.Equal(t, "Yo!", string(captured))
}

func Test_Builder_renderAndWrite_RenderError(t *testing.T) {
	tpl, _ := pongo2.FromString(`{{ 'fail'|nonexistent }}`)

	builder := &Builder{
		WriteFile: os.WriteFile,
		Logger:    &fakeLogger{},
	}

	assertPanics(t, func() {
		builder.renderAndWrite(tpl, pongo2.Context{}, "ignored.txt")
	}, "expected panic on render failure")
}

func Test_Builder_renderAndWrite_WriteError(t *testing.T) {
	tpl, _ := pongo2.FromString(`OK`)

	builder := &Builder{
		WriteFile: func(string, []byte, os.FileMode) error {
			return errors.New("disk full")
		},
		Logger: &fakeLogger{},
	}

	assertPanics(t, func() {
		builder.renderAndWrite(tpl, pongo2.Context{}, "fail.txt")
	}, "expected panic on write error")
}
