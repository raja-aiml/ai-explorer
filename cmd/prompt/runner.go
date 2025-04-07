package cmd

import (
	"fmt"
	"io"

	"raja.aiml/ai.explorer/prompt"
	"raja.aiml/ai.explorer/resources"
)

// Runner defines the interface for anything that can generate a prompt.
type Runner interface {
	Run()
}

// PromptRunner holds values required to generate a prompt.
type PromptRunner struct {
	Out            io.Writer
	Renderer       prompt.Renderer
	PromptCategory string
	Topic          string
	Template       string
	Config         string
	Output         string
	Preview        bool
	UserQuery      string
}

// ResolvePaths infers missing paths from category and topic.
func (r *PromptRunner) ResolvePaths() (tmpl, cfg, out string) {
	if r.PromptCategory == "" {
		r.PromptCategory = defaultPromptCategory
	}
	if r.Topic == "" {
		r.Topic = defaultTopic
	}

	tmpl, cfg, out = r.Template, r.Config, r.Output
	if tmpl == "" || cfg == "" || out == "" {
		resolver := resources.PathResolver{PromptCategory: r.PromptCategory}
		defTmpl, defCfg, defOut := resolver.Derive(r.Topic)
		if tmpl == "" {
			tmpl = defTmpl
		}
		if cfg == "" {
			cfg = defCfg
		}
		if out == "" {
			out = defOut
		}
	}
	return
}

// Run generates or previews the prompt.
func (r *PromptRunner) Run() {
	tmpl, cfg, out := r.ResolvePaths()
	queryArgs := r.queryArgs()

	if r.Preview {
		r.Renderer.RenderToStdout(tmpl, cfg, queryArgs...)
		return
	}

	r.Renderer.RenderToFile(tmpl, cfg, out, queryArgs...)
	fmt.Fprintf(r.Out, "Prompt saved to: %s\n", out)
}

// queryArgs returns userQuery as a variadic string slice if non-empty.
func (r *PromptRunner) queryArgs() []string {
	if r.UserQuery != "" {
		return []string{r.UserQuery}
	}
	return nil
}
