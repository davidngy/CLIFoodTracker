package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found")
	}
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	ctx := context.Background()
	db, err := connectDB(ctx)
	if err != nil {
		fmt.Println("DB error", err)
		return
	}

	defer db.Close()
	for {
		fmt.Print("diet >	")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		words := strings.Fields(input)
		if len(words) == 0 {
			println("missing args")
			continue
		}
		cmdName := words[0]
		cmd, ok := commands[cmdName]
		if !ok {
			fmt.Println("command not found")
			continue
		}

		args := words[1:]

		err := cmd.Callback(args)
		if err != nil {
			fmt.Println(err)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("input error:", err)
	}
}
