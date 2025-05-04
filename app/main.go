package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		fmt.Println(command + ": command not found")
	}
}
