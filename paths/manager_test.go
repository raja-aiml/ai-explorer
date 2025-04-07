package paths

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigPath_Custom(t *testing.T) {
	custom := "/custom/config.yaml"
	result := GetConfigPath("anytopic", custom)
	if result != custom {
		t.Errorf("Expected custom path %q, got %q", custom, result)
	}
}

func TestGetConfigPath_Default(t *testing.T) {
	topic := "testtopic"
	expected := fmt.Sprintf(ConfigPathFormat, topic)
	result := GetConfigPath(topic, "")
	if result != expected {
		t.Errorf("Expected default path %q, got %q", expected, result)
	}
}

func TestGetOutputPath_Custom(t *testing.T) {
	custom := "/custom/output.txt"
	result := GetOutputPath("anytopic", custom)
	if result != custom {
		t.Errorf("Expected custom path %q, got %q", custom, result)
	}
}

func TestGetOutputPath_Default(t *testing.T) {
	topic := "testtopic"
	expected := fmt.Sprintf(OutputPathFormat, topic)
	result := GetOutputPath(topic, "")
	if result != expected {
		t.Errorf("Expected default path %q, got %q", expected, result)
	}
}

func TestGetAnswerPath_Custom(t *testing.T) {
	custom := "/custom/answer.md"
	result := GetAnswerPath("anytopic", custom)
	if result != custom {
		t.Errorf("Expected custom path %q, got %q", custom, result)
	}
}

func TestGetAnswerPath_Default(t *testing.T) {
	topic := "testtopic"
	expected := fmt.Sprintf(AnswerPathFormat, topic)
	result := GetAnswerPath(topic, "")
	if result != expected {
		t.Errorf("Expected default path %q, got %q", expected, result)
	}
}

func TestGetTemplatePath_Custom(t *testing.T) {
	custom := "/custom/topic.yaml"
	result := GetTemplatePath(custom)
	if result != custom {
		t.Errorf("Expected custom template path %q, got %q", custom, result)
	}
}

func TestGetTemplatePath_Default(t *testing.T) {
	expected := TemplateFilePath
	result := GetTemplatePath("")
	if result != expected {
		t.Errorf("Expected default template path %q, got %q", expected, result)
	}
}

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
