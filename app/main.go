package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	COMMAND_EXIT = "exit"
)

func main() {
	// infinite loop implements REPL (Read-Eval-Print Loop)
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
			return
		}

		command = strings.TrimSpace(command)

		if strings.HasPrefix(command, COMMAND_EXIT) {
			splits := strings.SplitN(command, " ", 2)

			exitCode := 0

			if len(splits) == 2 {
				code, err := strconv.Atoi(splits[1])
				if err == nil {
					exitCode = code
				}
			}

			os.Exit(exitCode)
			return
		}

		fmt.Println(command + ": command not found")
	}
}
