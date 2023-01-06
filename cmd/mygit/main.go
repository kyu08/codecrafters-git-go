package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
		os.Exit(1)
	}

	command := os.Args[1]
	opt := getElem(os.Args, 2)
	optValue := getElem(os.Args, 3)

	switch command {
	case "init":
		gitInit()
	case "cat-file":
		catFile(opt, optValue)
	case "hash-object":
		hashObject(opt, optValue)
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
		os.Exit(1)
	}
}
