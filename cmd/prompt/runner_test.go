package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// --- Mock Implementation ---

type mockRenderer struct {
	ToFileCalled   bool
	ToStdoutCalled bool
	GotTmpl        string
	GotCfg         string
	GotOut         string
	GotQuery       string
}

func (m *mockRenderer) RenderToFile(tmpl, cfg, out string, userQuery ...string) {
	m.ToFileCalled = true
	m.GotTmpl = tmpl
	m.GotCfg = cfg
	m.GotOut = out
	if len(userQuery) > 0 {
		m.GotQuery = userQuery[0]
	}
}

func (m *mockRenderer) RenderToStdout(tmpl, cfg string, userQuery ...string) {
	m.ToStdoutCalled = true
	m.GotTmpl = tmpl
	m.GotCfg = cfg
	if len(userQuery) > 0 {
		m.GotQuery = userQuery[0]
	}
}

// --- Tests ---

func TestPromptRunner_Run_ExplicitPaths(t *testing.T) {
	var buf bytes.Buffer
	renderer := &mockRenderer{}

	runner := PromptRunner{
		Out:            &buf,
		Renderer:       renderer,
		PromptCategory: "topics",
		Topic:          "ignored",
		Template:       "template.yaml",
		Config:         "config.yaml",
		Output:         "output.txt",
		Preview:        false,
		UserQuery:      "explain Git in simple terms",
	}

	runner.Run()

	assert.True(t, renderer.ToFileCalled)
	assert.Equal(t, "template.yaml", renderer.GotTmpl)
	assert.Equal(t, "config.yaml", renderer.GotCfg)
	assert.Equal(t, "output.txt", renderer.GotOut)
	assert.Equal(t, "explain Git in simple terms", renderer.GotQuery)
	assert.Contains(t, buf.String(), "Prompt saved to: output.txt")
}

func TestPromptRunner_Run_WithPreview(t *testing.T) {
	var buf bytes.Buffer
	renderer := &mockRenderer{}

	runner := PromptRunner{
		Out:            &buf,
		Renderer:       renderer,
		PromptCategory: "topics",
		Topic:          "demo",
		Preview:        true,
		UserQuery:      "what is prompt engineering?",
	}

	runner.Run()

	assert.True(t, renderer.ToStdoutCalled)
	assert.Equal(t, "resources/topics/template.yaml", renderer.GotTmpl)
	assert.Equal(t, "resources/topics/demo/config.yaml", renderer.GotCfg)
	assert.Equal(t, "what is prompt engineering?", renderer.GotQuery)
	assert.NotContains(t, buf.String(), "Prompt saved to:")
}

func TestPromptRunner_Run_UsesDerivedPaths(t *testing.T) {
	var buf bytes.Buffer
	renderer := &mockRenderer{}

	runner := PromptRunner{
		Out:            &buf,
		Renderer:       renderer,
		PromptCategory: "topics",
		Topic:          "demo",
		UserQuery:      "what are embeddings?",
	}

	runner.Run()

	assert.True(t, renderer.ToFileCalled)
	assert.Equal(t, "resources/topics/template.yaml", renderer.GotTmpl)
	assert.Equal(t, "resources/topics/demo/config.yaml", renderer.GotCfg)
	assert.Equal(t, "resources/topics/demo/prompt.txt", renderer.GotOut)
	assert.Equal(t, "what are embeddings?", renderer.GotQuery)
	assert.Contains(t, buf.String(), "Prompt saved to: resources/topics/demo/prompt.txt")
}

func TestPromptRunner_Run_DefaultsWhenMissing(t *testing.T) {
	var buf bytes.Buffer
	renderer := &mockRenderer{}

	runner := PromptRunner{
		Out:       &buf,
		Renderer:  renderer,
		UserQuery: "summarize vector databases",
	}

	runner.Run()

	assert.True(t, renderer.ToFileCalled)
	assert.Equal(t, "resources/topics/template.yaml", renderer.GotTmpl)
	assert.Equal(t, "resources/topics/git/config.yaml", renderer.GotCfg)
	assert.Equal(t, "resources/topics/git/prompt.txt", renderer.GotOut)
	assert.Equal(t, "summarize vector databases", renderer.GotQuery)
	assert.Contains(t, buf.String(), "Prompt saved to: resources/topics/git/prompt.txt")
}
