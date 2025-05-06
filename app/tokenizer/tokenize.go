package tokenizer

import (
	"slices"
	"strings"
)

var backslashEscapedCharacters = []rune{'\\', '$', '"', '`', '\n'}

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
				}
				inDoubleQuote = !inDoubleQuote
			}
		case '\\':
			if inDoubleQuote {
				if isBackslash {
					isBackslash = false
					currentToken.WriteRune(char)
				} else {
					isBackslash = true
				}
			} else {
				currentToken.WriteRune(char)
			}
		case ' ':
			if inSingleQuote || inDoubleQuote {
				currentToken.WriteRune(char)
			} else if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
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
