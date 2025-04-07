package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathResolver_Derive(t *testing.T) {
	tests := []struct {
		name           string
		promptCategory string
		topic          string
		wantTmpl       string
		wantCfg        string
		wantOut        string
	}{
		{
			name:           "topics type - shared template outside topic folder",
			promptCategory: "topics",
			topic:          "demo",
			wantTmpl:       "resources/topics/template.yaml",
			wantCfg:        "resources/topics/demo/config.yaml",
			wantOut:        "resources/topics/demo/prompt.txt",
		},
		{
			name:           "topics type -  template inside topic folder",
			promptCategory: "chart",
			topic:          "flowchart",
			wantTmpl:       "resources/chart/flowchart/template.yaml",
			wantCfg:        "resources/chart/flowchart/config.yaml",
			wantOut:        "resources/chart/flowchart/prompt.txt",
		},
		{
			name:           "empty promptType",
			promptCategory: "",
			topic:          "basic",
			wantTmpl:       "resources//basic/template.yaml",
			wantCfg:        "resources//basic/config.yaml",
			wantOut:        "resources//basic/prompt.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resolver := PathResolver{PromptCategory: tt.promptCategory}
			gotTmpl, gotCfg, gotOut := resolver.Derive(tt.topic)

			assert.Equal(t, tt.wantTmpl, gotTmpl)
			assert.Equal(t, tt.wantCfg, gotCfg)
			assert.Equal(t, tt.wantOut, gotOut)
		})
	}
}
