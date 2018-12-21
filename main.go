package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	fmt.Println("Hello", args[0], "!")
	TapeInit()

	usefulArgs := os.Args[1:]
	stringInput := strings.Join(usefulArgs, "")
	commands := sanityze(stringInput)
	fmt.Println("Commands: ", commands)

	if validate(commands) {
		nodes := parse(commands)
		execute(nodes)
	} else {
		fmt.Println("Invalid Code! Hint: verify blocks ;)'[]'")
		// os.Exit(1)
	}
}
