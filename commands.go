package main

import "fmt"

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
		"user setup": {
			Name:        "user setup",
			Description: "Creating a User",
			Callback: func(args []string) error {
				fmt.Println("created")
				return nil
			},
		},
	}
}
