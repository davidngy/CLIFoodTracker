package main

import (
	"fmt"
	"os"
)

type Command struct {
	Name        string
	Description string
	Callback    func(*Config, []string) error
}

func getCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "Show available commands",
			Callback:    commandHelp,
		},
		"create-user": {
			Name:        "user setup",
			Description: "Creating a User",
			Callback:    commandCreateUser,
		},
		"select-user": {},
		"exit": {
			Name:        "exit",
			Description: "closing the application",
			Callback: func(config *Config, args []string) error {
				fmt.Println("bye!")
				os.Exit(0)
				return nil
			},
		},
	}
}
