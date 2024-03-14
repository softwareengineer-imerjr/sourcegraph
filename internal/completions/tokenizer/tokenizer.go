package tokenizer

import (
	_ "embed"
	"strings"

	"github.com/pkoukk/tiktoken-go"
)

// Tokenizer is an interface that all tokenizer implementations must satisfy.
type Tokenizer interface {
	Tokenize(text string) ([]int, error)
}

type TiktokenTokenizer struct {
	tk    *tiktoken.Tiktoken
	model string
}

func (t *TiktokenTokenizer) Tokenize(text string) ([]int, error) {
	return t.tk.Encode(text, []string{"all"}, nil), nil
}

// NewTokenizer returns a Tokenizer instance based on the provided model.
func NewTokenizer(model string) (Tokenizer, error) {
	switch {
	case strings.Contains(model, "anthropic"):
		return NewAnthropicClaudeTokenizer(model)
	case strings.Contains(model, "azure"), strings.Contains(model, "openai"):
		// Returning a TiktokenTokenizer for models related to "azure" or "openai"
		return NewOpenAITokenizer(model)
	default:
		return nil, nil
	}
}

func NewOpenAITokenizer(model string) (*TiktokenTokenizer, error) {
	tkm, _ := tiktoken.EncodingForModel(model)
	return &TiktokenTokenizer{
		tk:    tkm,
		model: model,
	}, nil
}
