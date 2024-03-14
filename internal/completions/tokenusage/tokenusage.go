package tokenusage

import (
	"fmt"

	"github.com/sourcegraph/sourcegraph/internal/completions/tokenizer"
	"github.com/sourcegraph/sourcegraph/internal/rcache"
)

type TokenUsageManager struct {
	Cache *rcache.Cache
}

func NewTokenUsageManager() *TokenUsageManager {
	return &TokenUsageManager{
		Cache: rcache.NewWithTTL("LLMUsage", 1800),
	}
}

func (m *TokenUsageManager) TokenizeAndCalculateUsage(inputText, outputText, model, feature string, stream bool) error {
	tokenizer, err := tokenizer.NewTokenizer(model)
	if err != nil {
		return err
	}

	inputTokens, _ := tokenizer.Tokenize(inputText)
	outputTokens, _ := tokenizer.Tokenize(outputText)

	baseKey := fmt.Sprintf("%s:%s:%s:", model, feature, streamDescription(stream))

	// Calculate and update token counts in Redis
	m.updateTokenCounts(baseKey+"input", len(inputTokens))
	m.updateTokenCounts(baseKey+"output", len(outputTokens))
	return nil
}

// Helper function to get the description of the request type
func streamDescription(stream bool) string {
	if stream {
		return "stream"
	}
	return "non-stream"
}

func (m *TokenUsageManager) updateTokenCounts(key string, tokenCount int) {
	currentTokens, _ := m.Cache.GetInt(key)
	newTokens := currentTokens + tokenCount
	m.Cache.SetInt(key, newTokens)
}

func (m *TokenUsageManager) GetTokenCounts(key string) (int, bool) {
	return m.Cache.GetInt(key)
}
