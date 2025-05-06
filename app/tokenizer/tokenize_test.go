package tokenizer_test

import (
	"testing"

	"github.com/codecrafters-io/shell-starter-go/app/tokenizer"
	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	t.Run("Empty arguments", func(t *testing.T) {
		tokens := tokenizer.Tokenize("")
		assert.Equal(t, tokens, []string{})
	})

	t.Run("Single argument", func(t *testing.T) {
		tokens := tokenizer.Tokenize("arg1")
		assert.Equal(t, tokens, []string{"arg1"})
	})

	t.Run("Multiple arguments", func(t *testing.T) {
		tokens := tokenizer.Tokenize("arg1 arg2 arg3")
		assert.Equal(t, tokens, []string{"arg1", "arg2", "arg3"})
	})

	t.Run("Multiple spaces between arguments", func(t *testing.T) {
		tokens := tokenizer.Tokenize("arg1  arg2   arg3")
		assert.Equal(t, tokens, []string{"arg1", "arg2", "arg3"})
	})

	t.Run("Single argument with quote", func(t *testing.T) {
		tokens := tokenizer.Tokenize("'arg1'")
		assert.Equal(t, tokens, []string{"arg1"})
	})

	t.Run("Multiple arguments with quote", func(t *testing.T) {
		tokens := tokenizer.Tokenize("'arg1' 'arg 2' 'arg 3 4'")
		assert.Equal(t, tokens, []string{"arg1", "arg 2", "arg 3 4"})
	})

	t.Run("Unclosed quote", func(t *testing.T) {
		tokens := tokenizer.Tokenize("'arg1")
		assert.Equal(t, tokens, []string{"arg1"})

		tokens = tokenizer.Tokenize("arg1'")
		assert.Equal(t, tokens, []string{"arg1"})
	})

	t.Run("Non-breaking quote", func(t *testing.T) {
		tokens := tokenizer.Tokenize("arg1''arg2")
		assert.Equal(t, tokens, []string{"arg1arg2"})
	})

	t.Run("Multiple mixed arguments", func(t *testing.T) {
		tokens := tokenizer.Tokenize("arg1 arg 2 'arg 3 4'")
		assert.Equal(t, tokens, []string{"arg1", "arg", "2", "arg 3 4"})
	})

	t.Run("Preserve empty spaces inside quotes", func(t *testing.T) {
		tokens := tokenizer.Tokenize("'arg1   arg 2'")
		assert.Equal(t, tokens, []string{"arg1   arg 2"})
	})
}
