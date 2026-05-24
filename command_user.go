package main

import (
	"context"
	"fmt"
)

func commandCreateUser(cfg *Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: create-user <name>")
	}
	name := args[0]
	user, err := CreateUser(context.Background(), cfg.DB, name)
	if err != nil {
		return err
	}

	fmt.Printf("created use: %s (id: %d)\n", user.Name, user.ID)
	return nil
}
