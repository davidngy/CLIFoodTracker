package main

import "fmt"

func commandHelp(args []string) error {
	commands := getCommands()
	for _, cmd := range commands {
		fmt.Println(cmd.Name, cmd.Description)
	}

	return nil
}
