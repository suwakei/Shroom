package main

import (
	"Shroom/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(user)
	}
	fmt.Printf("Hello!! %s! This is repl of programming langage Shroom!!\n", user.Username)
	fmt.Printf("Press any commands!\n")
	repl.Start(os.Stdin, os.Stdout)
}
