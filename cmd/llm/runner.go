package llm

import (
	"fmt"
	"io"
	"log"
)

// LLMRunner handles prompt loading, LLM interaction, and output.
type LLMRunner struct {
	Out          io.Writer
	PromptPath   string
	OutputPath   string
	GetPrompt    func(string) (string, error)
	RunLLM       func(string) (string, error)
	SaveResponse func(string, string) error
}

// Run executes the LLM flow.
// Run executes the LLM flow.
func (r *LLMRunner) Run() {
	fmt.Fprintln(r.Out, "[llm] Reading prompt...")
	prompt, err := r.GetPrompt(r.PromptPath)
	if err != nil {
		log.Fatalf("[llm] Prompt error: %v", err)
	}

	fmt.Fprintln(r.Out, "[llm] Running LLM...")
	resp, err := r.RunLLM(prompt)
	if err != nil {
		log.Fatalf("[llm] LLM error: %v", err)
	}

	if r.OutputPath != "" {
		err := r.SaveResponse(resp, r.OutputPath)
		if err != nil {
			log.Fatalf("[llm] Failed to save response: %v", err)
		}
		fmt.Fprintf(r.Out, "[llm] 💾 LLM response saved to: %s\n", r.OutputPath)
	}
}
