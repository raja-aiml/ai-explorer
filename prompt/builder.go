package prompt

import (
	"fmt"
	"os"

	"github.com/flosch/pongo2/v6"
	"gopkg.in/yaml.v3"
	"raja.aiml/ai.explorer/logger"
	"raja.aiml/ai.explorer/paths"
)

// Builder implements the Renderer interface using pongo2 and YAML config.
type Builder struct {
	ReadFile  func(path string) ([]byte, error)
	WriteFile func(path string, data []byte, perm os.FileMode) error
	Logger    logger.Logger
}

// Ensure Builder satisfies Renderer interface.
var _ Renderer = (*Builder)(nil)

// --- Public API ---

func (b *Builder) RenderToFile(templatePath, configPath, outputPath string, userQuery ...string) {
	tpl := b.mustParseTemplate(templatePath)
	ctx := b.mustParseConfig(configPath, userQuery...)
	b.renderAndWrite(tpl, ctx, outputPath)
}

func (b *Builder) RenderToStdout(templatePath, configPath string, userQuery ...string) {
	tpl := b.mustParseTemplate(templatePath)
	ctx := b.mustParseConfig(configPath, userQuery...)

	out, err := tpl.Execute(ctx)
	if err != nil {
		b.Logger.Fatalf("template rendering failed: %v", err)
	}
	fmt.Println(out)
}

// --- Internal helpers ---

func (b *Builder) mustParseTemplate(path string) *pongo2.Template {
	data, err := b.ReadFile(path)
	if err != nil {
		b.Logger.Fatalf("failed to read template file: %v", err)
	}

	tpl, err := pongo2.FromString(string(data))
	if err != nil {
		b.Logger.Fatalf("failed to parse template: %v", err)
	}
	return tpl
}

func (b *Builder) mustParseConfig(path string, userQuery ...string) pongo2.Context {
	data, err := b.ReadFile(path)
	if err != nil {
		b.Logger.Fatalf("failed to read config file: %v", err)
	}

	var parsed map[string]any
	if err := yaml.Unmarshal(data, &parsed); err != nil {
		b.Logger.Fatalf("failed to parse YAML config: %v", err)
	}

	if len(userQuery) > 0 && userQuery[0] != "" {
		parsed["user_query"] = userQuery[0]
	}

	return pongo2.Context(parsed)
}

func (b *Builder) renderAndWrite(tpl *pongo2.Template, ctx pongo2.Context, outPath string) {
	out, err := tpl.Execute(ctx)
	if err != nil {
		b.Logger.Fatalf("template rendering failed: %v", err)
	}
	paths.EnsureDirectoryExists(outPath)
	if err := b.WriteFile(outPath, []byte(out), 0644); err != nil {
		b.Logger.Fatalf("failed to write output file: %v", err)
	}
}
