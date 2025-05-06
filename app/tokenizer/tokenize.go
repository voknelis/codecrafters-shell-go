package tokenizer

import "strings"

func Tokenize(input string) []string {
	tokens := make([]string, 0)
	currentToken := strings.Builder{}
	isQuote := false

	for _, char := range input {
		if char == '\'' {
			isQuote = !isQuote
			continue
		}

		if char == ' ' && !isQuote {
			if currentToken.Len() == 0 {
				continue
			}

			tokens = append(tokens, currentToken.String())
			currentToken.Reset()
			continue
		}

		currentToken.WriteRune(char)
	}

	// flush the rest of the last token data
	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}

	return tokens
}
