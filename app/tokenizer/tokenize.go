package tokenizer

import (
	"slices"
	"strings"
)

type tokenizerState struct {
	inSingleQuote bool
	inDoubleQuote bool
	isEscaped     bool
	currentToken  strings.Builder
	tokens        []string
}

const (
	singleQuote = '\''
	doubleQuote = '"'
	backslash   = '\\'
	space       = ' '
)

var escapedCharactersInDoubleQuotes = []rune{'\\', '$', '"', '`', '\n'}

func Tokenize(input string) []string {
	state := tokenizerState{
		tokens: []string{},
	}

	for _, char := range input {
		state.processChar(char)
	}

	state.flushToken()

	return state.tokens
}

func (t *tokenizerState) processChar(char rune) {
	switch {
	case t.inSingleQuote:
		t.processCharInSingleQuote(char)
	case t.inDoubleQuote:
		t.processCharInDoubleQuote(char)
	default:
		t.processUnquotedChar(char)
	}
}

// inside single quote:
//   - no escaping is performed at all
//   - every character between the single quotes is literal
//   - you cannot escape anything inside single quotes — not even another single quote
//     except to include a single quote, you have to end the quote, escape it, and reopen:
//     'It'\”s fine' -> It's fine
func (t *tokenizerState) processCharInSingleQuote(char rune) {
	if char == singleQuote {
		t.inSingleQuote = false
		return
	}

	t.currentToken.WriteRune(char)
}

// inside double quote
//   - backslash (\) only escapes these four characters: $, `, ", \
//     "Hello \$USER" -> Hello $USER
//     "Hello \"World\"" -> Hello "World"
//     "Slash: \\" -> Slash: \
func (t *tokenizerState) processCharInDoubleQuote(char rune) {
	if t.isEscaped {
		if slices.Contains(escapedCharactersInDoubleQuotes, char) {
			t.processEscapedChar(char)
		} else {
			t.isEscaped = false
			t.currentToken.WriteRune('\\')
			t.currentToken.WriteRune(char)
		}
		return
	}

	switch char {
	case backslash:
		t.isEscaped = true
	case doubleQuote:
		t.inDoubleQuote = false
	default:
		t.currentToken.WriteRune(char)
	}

}

// unquoted (bare words):
// - the backslash (\):
//   - escapes the next character
//   - prevents special characters from being interpreted (like space, *, $, etc.)
func (t *tokenizerState) processUnquotedChar(char rune) {
	if t.isEscaped {
		t.processEscapedChar(char)
		return
	}

	switch char {
	case backslash:
		t.isEscaped = true
	case singleQuote:
		t.inSingleQuote = true
	case doubleQuote:
		t.inDoubleQuote = true
	case space:
		t.flushToken()
	default:
		t.currentToken.WriteRune(char)
	}
}

func (t *tokenizerState) processEscapedChar(char rune) {
	t.currentToken.WriteRune(char)
	t.isEscaped = false
}

func (t *tokenizerState) flushToken() {
	if t.currentToken.Len() > 0 {
		t.tokens = append(t.tokens, t.currentToken.String())
		t.currentToken.Reset()
	}
}
