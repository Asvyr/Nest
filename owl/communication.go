package main

import (
	"fmt"
	"net"
)

func HandleCommunication(protocol string) {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	switch protocol {
	case "pull":
		fmt.Println("Pulling changes")
	case "push":
		fmt.Println("Pushing changes")
	case "init":
		// Does this actually need to be here?
		// this will probs fall into the pull/push portions
	}

}
