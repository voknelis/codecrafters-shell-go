package tokenizer_test

import (
	"testing"

	"github.com/codecrafters-io/shell-starter-go/app/tokenizer"
	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	t.Run("Unquoted", func(t *testing.T) {
		t.Run("Empty arguments", func(t *testing.T) {
			tokens := tokenizer.Tokenize("")
			assert.Equal(t, []string{}, tokens)
		})

		t.Run("Single argument", func(t *testing.T) {
			tokens := tokenizer.Tokenize("arg1")
			assert.Equal(t, []string{"arg1"}, tokens)
		})

		t.Run("Multiple arguments", func(t *testing.T) {
			tokens := tokenizer.Tokenize("arg1 arg2 arg3")
			assert.Equal(t, []string{"arg1", "arg2", "arg3"}, tokens)
		})

		t.Run("Multiple spaces between arguments", func(t *testing.T) {
			tokens := tokenizer.Tokenize("arg1  arg2   arg3")
			assert.Equal(t, []string{"arg1", "arg2", "arg3"}, tokens)
		})

		t.Run("Escaped charachers", func(t *testing.T) {
			// Preserve escaped spaces
			tokens := tokenizer.Tokenize(`arg1\ \ \ arg2`)
			assert.Equal(t, []string{"arg1   arg2"}, tokens)

			// preserve single quote
			tokens = tokenizer.Tokenize(`\'arg1\'"`)
			assert.Equal(t, []string{`'arg1'`}, tokens)

			// preserve double quotes
			tokens = tokenizer.Tokenize(`\"arg1\"`)
			assert.Equal(t, []string{`"arg1"`}, tokens)
		})
	})

	t.Run("Single quotes", func(t *testing.T) {
		t.Run("Argument with single quote", func(t *testing.T) {
			tokens := tokenizer.Tokenize("'arg1'")
			assert.Equal(t, []string{"arg1"}, tokens)
		})

		t.Run("Multiple arguments with single quote", func(t *testing.T) {
			tokens := tokenizer.Tokenize("'arg1' 'arg 2' 'arg 3 4'")
			assert.Equal(t, []string{"arg1", "arg 2", "arg 3 4"}, tokens)
		})

		t.Run("Unclosed single quote", func(t *testing.T) {
			tokens := tokenizer.Tokenize("'arg1")
			assert.Equal(t, []string{"arg1"}, tokens)

			tokens = tokenizer.Tokenize("arg1'")
			assert.Equal(t, []string{"arg1"}, tokens)
		})

		t.Run("Non-breaking single quote", func(t *testing.T) {
			tokens := tokenizer.Tokenize("arg1''arg2")
			assert.Equal(t, []string{"arg1arg2"}, tokens)
		})

		t.Run("Preserve empty spaces inside single quotes", func(t *testing.T) {
			tokens := tokenizer.Tokenize("'arg1   arg 2'")
			assert.Equal(t, []string{"arg1   arg 2"}, tokens)
		})
	})

	t.Run("Double quotes", func(t *testing.T) {
		t.Run("Argument with double quote", func(t *testing.T) {
			tokens := tokenizer.Tokenize(`"arg1"`)
			assert.Equal(t, []string{"arg1"}, tokens)
		})

		t.Run("Escaped charachers", func(t *testing.T) {
			// preserve single quote
			tokens := tokenizer.Tokenize(`"arg1'"`)
			assert.Equal(t, []string{"arg1'"}, tokens)

			// preserve single backslash
			tokens = tokenizer.Tokenize(`"arg1\arg2"`)
			assert.Equal(t, []string{"arg1\\arg2"}, tokens)
			tokens = tokenizer.Tokenize(`"arg1\ arg2"`)
			assert.Equal(t, []string{"arg1\\ arg2"}, tokens)

			// single quote inside double quote cannot be escaped
			tokens = tokenizer.Tokenize(`"arg1\'"`)
			assert.Equal(t, []string{`arg1\'`}, tokens)

			// escape double quote
			tokens = tokenizer.Tokenize(`"arg1\""`)
			assert.Equal(t, []string{"arg1\""}, tokens)

			// escape double quote and preserve following single quote
			tokens = tokenizer.Tokenize(`"arg1\"arg2'arg3"`)
			assert.Equal(t, []string{`arg1"arg2'arg3`}, tokens)

			// escape backslash
			tokens = tokenizer.Tokenize(`"arg1\\"`)
			assert.Equal(t, []string{"arg1\\"}, tokens)

			// escape dollar sign
			tokens = tokenizer.Tokenize(`"arg1\$"`)
			assert.Equal(t, []string{"arg1$"}, tokens)

			// escape backtick symbol
			tokens = tokenizer.Tokenize("\"arg1`\"")
			assert.Equal(t, []string{"arg1`"}, tokens)

			// escape new line
			tokens = tokenizer.Tokenize("\"arg1\n\"")
			assert.Equal(t, []string{"arg1\n"}, tokens)
		})

		t.Run("Multiple arguments with double quote", func(t *testing.T) {
			tokens := tokenizer.Tokenize(`"arg1" "arg 2" "arg \3" "arg \\"`)
			assert.Equal(t, []string{"arg1", "arg 2", "arg \\3", "arg \\"}, tokens)
		})

		t.Run("Unclosed double quote", func(t *testing.T) {
			tokens := tokenizer.Tokenize(`"arg1`)
			assert.Equal(t, []string{"arg1"}, tokens)

			tokens = tokenizer.Tokenize(`arg1"`)
			assert.Equal(t, []string{"arg1"}, tokens)
		})

		t.Run("Non-breaking double quote", func(t *testing.T) {
			tokens := tokenizer.Tokenize(`arg1""arg2`)
			assert.Equal(t, []string{"arg1arg2"}, tokens)
		})

		t.Run("Preserve empty spaces inside dougle quotes", func(t *testing.T) {
			tokens := tokenizer.Tokenize(`"arg1   arg 2"`)
			assert.Equal(t, []string{"arg1   arg 2"}, tokens)
		})
	})

	t.Run("Multiple mixed arguments", func(t *testing.T) {
		tokens := tokenizer.Tokenize("arg1 arg 2 'arg 3 4'")
		assert.Equal(t, []string{"arg1", "arg", "2", "arg 3 4"}, tokens)
	})
}
