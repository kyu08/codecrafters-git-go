package main

import (
	"fmt"
	"os"
)

// Usage: your_git.sh <command> <arg1> <arg2> ...
func main() {
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
			// 引数を元にファイルの中身を返す
			// blob sha をファイルパスに変換する
			// ファイル内容を標準出力に出力する
			// TODO: テストケースとgit cat-fileの定義を読んで挙動を把握
			fmt.Print("dumpty doo dumpty dumpty yikes donkey")
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
