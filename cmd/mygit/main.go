package main

import (
	"fmt"
	"os"
)

// Usage: your_git.sh <command> <arg1> <arg2> ...
func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	fmt.Printf("os.Args: %v\n", os.Args)
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
		os.Exit(1)
	}

	switch command, opt, optValue := os.Args[1], getElem(os.Args, 2), getElem(os.Args, 3); command {
	case "init":
		for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
			// MEMO: 0755: rwxr-xr-x
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
			}
		}

		headFileContents := []byte("ref: refs/heads/master\n")
		if err := os.WriteFile(".git/HEAD", headFileContents, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
		}

		fmt.Println("Initialized git directory")

	case "cat-file":
		switch *opt {
		case "-p":
			// TODO: テストケースとgit cat-fileの定義を読んで挙動を把握
			fmt.Println("-p")
			fmt.Printf("optValue: %v\n", optValue)
		}

	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
		os.Exit(1)
	}
}

func getElem(args []string, index int) *string {
	if len(args) > index {
		return &args[index]
	}

	return nil
}
