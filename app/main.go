package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/codecrafters-io/shell-starter-go/app/command"
)

func main() {
	handleSysCall()

	// infinite loop implements REPL (Read-Eval-Print Loop)
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			// handle Ctrl+D
			if errors.Is(err, io.EOF) {
				fmt.Println("\nexiting shell")
				os.Exit(0)
			}

			fmt.Fprintln(os.Stderr, "\nerror reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		executeCommand(input)
	}
}

func executeCommand(input string) {
	command, err := command.NewCommand(input)
	if err != nil {
		fmt.Println(input + ": command not found")
		return
	}

	command.Exec()
}

func handleSysCall() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for sig := range signalChannel {
			switch sig {
			// handle Ctrl+C
			case syscall.SIGINT:
				// just print a newline and continue
				fmt.Fprint(os.Stdout, "\n$ ")
			// graceful shutdown
			case syscall.SIGTERM:
				fmt.Println("\nreceived SIGTERM, exiting...")
				os.Exit(0)
			}
		}
	}()
}
