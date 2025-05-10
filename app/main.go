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
	"github.com/codecrafters-io/shell-starter-go/app/redirect"
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
	cmd, err := command.NewCommand(input)
	if err != nil {
		if errors.Is(err, command.ErrUnknownCommand) {
			fmt.Println(input + ": command not found")
			return
		}

		fmt.Println(err)
		return
	}

	var stdout command.Writer = os.Stdout
	var stderr command.Writer = os.Stderr

	for _, r := range cmd.Redirection {
		switch r.FileDescriptor {
		case 1:
			stdout = redirect.NewRedirectStd(r.Target, r.Type == redirect.RedirectionOutputAppend)
		case 2:
			stderr = redirect.NewRedirectStd(r.Target, r.Type == redirect.RedirectionOutputAppend)
		}
	}

	err = cmd.Executable.Exec(stdout, stderr)
	if err != nil {
		fmt.Println("-shell:", err)
	}
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
