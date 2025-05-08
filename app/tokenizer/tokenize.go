package tokenizer

import (
	"slices"
	"strings"
)

var backslashEscapedCharacters = []rune{'\\', '$', '"', '`', '\n'}

// Shell Quoting Modes and Escape Behavior explained
//
// inside single quote:
// - no escaping is performed at all
// - every character between the single quotes is literal
// - you cannot escape anything inside single quotes — not even another single quote
//   except to include a single quote, you have to end the quote, escape it, and reopen:
//   'It'\''s fine' -> It's fine
//
// inside double quote
// - backslash (\) only escapes these four characters: $, `, ", \
// 	 "Hello \$USER" -> Hello $USER
//   "Hello \"World\"" -> Hello "World"
//   "Slash: \\" -> Slash: \
//
// unquoted (bare words):
// - the backslash (\) escapes the next character.
// - prevents special characters from being interpreted (like space, *, $, etc.).

func Tokenize(input string) []string {
	tokens := make([]string, 0)
	currentToken := strings.Builder{}
	inSingleQuote := false
	inDoubleQuote := false
	isBackslash := false

	for _, char := range input {
		switch char {
		case '\'':
			if inDoubleQuote {
				if isBackslash {
					isBackslash = false
					currentToken.WriteRune('\\')
				}
				currentToken.WriteRune(char)
			} else if isBackslash && !inSingleQuote && !inDoubleQuote {
				currentToken.WriteRune(char)
				isBackslash = false
			} else {
				inSingleQuote = !inSingleQuote
			}
		case '"':
			if inSingleQuote {
				currentToken.WriteRune(char)
			} else {
				if isBackslash {
					isBackslash = false
					currentToken.WriteRune(char)
				} else {
					inDoubleQuote = !inDoubleQuote
				}
			}
		case '\\':
			if inDoubleQuote {
				if isBackslash {
					isBackslash = false
					currentToken.WriteRune(char)
				} else {
					isBackslash = true
				}
			} else if inSingleQuote {
				// supposed to throw error, but for not just write it
				currentToken.WriteRune(char)
			} else if !inSingleQuote && !inDoubleQuote {
				isBackslash = true
			} else {
				currentToken.WriteRune(char)
			}
		case ' ':
			if isBackslash && !inDoubleQuote {
				currentToken.WriteRune(char)
			} else if inSingleQuote || inDoubleQuote {
				if isBackslash {
					currentToken.WriteRune('\\')
				}
				currentToken.WriteRune(char)
			} else if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}

			if isBackslash {
				isBackslash = false
			}
		default:
			if inDoubleQuote && isBackslash && !slices.Contains(backslashEscapedCharacters, char) {
				currentToken.WriteRune('\\')
			}
			if isBackslash {
				isBackslash = false
			}
			currentToken.WriteRune(char)
		}
	}

	// flush the rest of the last token data
	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}

	return tokens
}
