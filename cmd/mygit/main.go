package main

import (
	"fmt"
	"io/ioutil"
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
			// hashをファイルパスに変換する
			blobHash := optValue
			filePath := ".git/objects/" + *blobHashToFilePath(*blobHash)

			// ファイル内容を取得する
			b, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Printf("ReadFile failed: %s", err)
				os.Exit(1)
			}
			lines := string(b)

			// ファイル内容をcompress/zlibを使って解凍する
			unzippedLines := unzipLines(lines)

			// ファイル内容を標準出力に出力する
			fmt.Printf(unzippedLines)
		}

	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
		os.Exit(1)
	}
}

func getElem(args []string, index int64) *string {
	if int64(len(args)) >= (index + 1) {
		return &args[index]
	}

	return nil
}

func blobHashToFilePath(hash string) *string {
	if len(hash) < 3 {
		return nil
	}

	path := hash[0:2] + "/" + hash[2:]
	return &path
}

func unzipLines(lines string) string {
	return "TODO: implement unzipLines"
}
