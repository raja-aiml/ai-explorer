package llm

import (
	"fmt"

	llmConfig "raja.aiml/ai.explorer/llm/config"
	llmWrapper "raja.aiml/ai.explorer/llm/wrapper"
)

// initLLMProvider initializes the LLM model using the langchaingo wrapper.
func InitLLMProvider(cfg llmConfig.Config) (llmWrapper.Model, error) {
	provider := &llmWrapper.LangchaingoProvider{}
	model, err := provider.Init(cfg.Provider, cfg.Model.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize LLM provider: %w", err)
	}
	return model, nil
}
