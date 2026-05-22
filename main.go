package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

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
