package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/kanatti/just-go/interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, Test Interpreter v0.0.1\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
