package main

import (
	"Shroom/repl"
	"fmt"
	"os"
	"os/user"
)
//  モジュール名を全部srにしろ

func main() {
	if args := os.Args; len(args) - 1 == 0 {
	user, err := user.Current()
	if err != nil {
		panic(user)
	}
	fmt.Printf("Hello!! %s! This is repl of programming langage Shroom!!\n", user.Username)
	fmt.Printf("Type any commands!\n")
	repl.Start(os.Stdin, os.Stdout)
	}
}
