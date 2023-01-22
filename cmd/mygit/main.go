package main

import (
	"github.com/codecrafters-io/git-starter-go/cmd/mygit/handler"
)

func main() {
	// TODO: コマンドライン引数の受け取りにcobraを使ってみる
	// if err := handler.Handler(os.Args); err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s", err)
	// }

	cmd := handler.Command()
	cmd.Execute()
}
