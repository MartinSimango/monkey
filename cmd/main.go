package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/MartinSimango/monkey/repl"
)

func main() {

	us, err := user.Current()
	if err != nil {
		fmt.Println("Failed to get user")
	}
	fmt.Printf("Welcome %s enter some monkey\n", us.Name)
	repl.Start(os.Stdin, os.Stdout)
}
