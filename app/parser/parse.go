package parser

import (
	"github.com/codecrafters-io/shell-starter-go/app/redirect"
)

type CommandNode struct {
	Executable string
	Arguments  []string

	Redirection []redirect.Redirection
}

func Parse(tokens []string) (*CommandNode, error) {
	commandNode := &CommandNode{
		Arguments:   []string{},
		Redirection: []redirect.Redirection{},
	}

	if len(tokens) == 0 {
		return commandNode, nil
	}

	commandNode.Executable = tokens[0]

	redirectToken := ""
	for _, token := range tokens[1:] {
		// set redirect operator processing on next iteration
		if redirect.IsRedirectionOperator(token) {
			redirectToken = token
			continue
		}

		// process redirect operator and its argument
		if redirectToken != "" {
			redirection, err := redirect.ParseRedirection(redirectToken, token)
			if err != nil {
				return commandNode, err
			}

			err = redirect.CheckRedirectionTarget(token)
			if err != nil {
				return commandNode, err
			}

			commandNode.Redirection = append(commandNode.Redirection, *redirection)
			redirectToken = ""
			continue
		}

		// handle other tokens as arguments
		commandNode.Arguments = append(commandNode.Arguments, token)
	}

	return commandNode, nil
}
