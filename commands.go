package main

import (
	"fmt"
	"os"
)

type Command struct {
	Name        string
	Description string
	Callback    func(args []string) error
}

func getCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "Show available commands",
			Callback:    commandHelp,
		},
		"user-setup": {
			Name:        "user setup",
			Description: "Creating a User",
			Callback: func(args []string) error {
				fmt.Println("created")
				return nil
			},
		},
		"exit": {
			Name:        "exit",
			Description: "closing the application",
			Callback: func(args []string) error {
				fmt.Println("bye!")
				os.Exit(0)
				return nil
			},
		},
	}
}
