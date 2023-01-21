package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/git-starter-go/cmd/mygit/handler"
)

func main() {
	// コマンドライン引数の受け取りにcobraを使ってみたい
	if err := handler.Handler(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
}
