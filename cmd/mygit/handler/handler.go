package handler

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/git-starter-go/cmd/mygit/plumbing"
)

func Handler(args []string) {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
		os.Exit(1)
	}

	command := os.Args[1]
	opt := plumbing.GetElem(os.Args, 2)
	optValue := plumbing.GetElem(os.Args, 3)

	switch command {
	case "init":
		plumbing.GitInit()
	case "cat-file":
		plumbing.CatFile(opt, optValue)
	case "hash-object":
		plumbing.HashObject(opt, optValue)
	case "ls-tree":
		plumbing.LsTree(opt, optValue)
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
		os.Exit(1)
	}
}
