package tokenizer

import (
	_ "embed"

	"github.com/pkoukk/tiktoken-go"
)

// Tokenizer is an interface that all tokenizer implementations must satisfy.
type Tokenizer interface {
	Tokenize(text string) ([]int, error)
}

type TiktokenTokenizer struct {
	tk *tiktoken.Tiktoken
}

func (t *TiktokenTokenizer) Tokenize(text string) ([]int, error) {
	return t.tk.Encode(text, []string{"all"}, nil), nil
}

// NewTokenizer returns a Tokenizer instance based on the provided model.
func NewTokenizer(model string) (Tokenizer, error) {
	switch model {
	case "anthropic/claude-2":
		// Assuming NewAnthropicClaudeTokenizer() is correctly implemented elsewhere
		return NewAnthropicClaudeTokenizer()
	case "openai":
		// Example: Returning a TiktokenTokenizer for the "openai/gpt-3" model
		// Initialize your tiktoken.Tiktoken instance as needed
		return NewOpenAITokenizer()
	default:
		return nil, nil
	}
}

// NewAnthropicClaudeTokenizer is a tokenizer that emulates Anthropic's
// tokenization for Claude.
func NewOpenAITokenizer() (*TiktokenTokenizer, error) {
	encoding := "cl100k_base"
	tke, _ := tiktoken.GetEncoding(encoding)
	return &TiktokenTokenizer{
		tk: tke,
	}, nil
}
