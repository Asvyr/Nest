package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) == 2 {
		switch strings.ToLower(args[1]) {
		case "pull":
			fmt.Print("Pulling changes")
		case "push":
			fmt.Println("Pushing changes")
		case "init":
			fmt.Println("Initilizing project")
		case "help":
			fmt.Println("Args: 'pull', 'push', 'init'")
		default:
			fmt.Println("Invalid arguments. type owl help for a list of commands")
		}
	}
}
