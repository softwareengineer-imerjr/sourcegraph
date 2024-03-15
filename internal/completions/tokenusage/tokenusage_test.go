package tokenusage_test

import (
	"testing"

	"github.com/sourcegraph/sourcegraph/internal/completions/tokenusage"
	"github.com/sourcegraph/sourcegraph/internal/rcache"
)

func TestTokenizeAndCalculateUsage(t *testing.T) {
	rcache.SetupForTest(t)
	mockCache := rcache.NewWithTTL("LLMUsage", 1800)
	manager := tokenusage.NewTokenUsageManager()

	err := manager.TokenizeAndCalculateUsage("input text", "output text", "anthropic", "feature1", true)
	if err != nil {
		t.Fatalf("TokenizeAndCalculateUsage returned an error: %v", err)
	}

	// Verify that token counts are updated in the cache
	inputKey := "anthropic:feature1:stream:input"
	outputKey := "anthropic:feature1:stream:output"

	if val, exists := mockCache.GetInt(inputKey); !exists || val <= 0 {
		t.Errorf("Expected input token count to be updated in cache, but key %s was not found or value is not positive", inputKey)
	}

	if val, exists := mockCache.GetInt(outputKey); !exists || val <= 0 {
		t.Errorf("Expected output token count to be updated in cache, but key %s was not found or value is not positive", outputKey)
	}
}

func TestGetAllTokenUsageData(t *testing.T) {
	rcache.SetupForTest(t)
	manager := tokenusage.NewTokenUsageManager()
	cache := rcache.NewWithTTL("LLMUsage", 1800)
	cache.SetInt("LLMUsage:model1:feature1:stream:input", 10)
	cache.SetInt("LLMUsage:model1:feature1:stream:output", 20)

	usageData := manager.GetAllTokenUsageData()

	if len(usageData) != 2 {
		t.Errorf("Expected 2 items in usage data, got %d", len(usageData))
	}
	if usageData["LLMUsage:model1:feature1:stream:input"] != 10 {
		t.Errorf("Expected 10 tokens for input, got %d", usageData["LLMUsage:model1:feature1:stream:input"])
	}

	if usageData["LLMUsage:model1:feature1:stream:output"] != 20 {
		t.Errorf("Expected 20 tokens for output, got %d", usageData["LLMUsage:model1:feature1:stream:input"])
	}
}
