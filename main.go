package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/SebbieMzingKe/intepreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("hello %s! this is the mzing programming language1\n", user.Username)
	fmt.Printf("be free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}