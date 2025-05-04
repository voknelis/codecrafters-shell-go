package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/command"
)

func main() {
	// infinite loop implements REPL (Read-Eval-Print Loop)
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
			return
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		command, err := command.NewCommand(input)
		if err != nil {
			fmt.Println(input + ": command not found")
			continue
		}

		command.Exec()
	}
}
