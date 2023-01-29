package main

import (
	"github.com/codecrafters-io/git-starter-go/cmd/mygit/handler"
)

func main() {
	// TODO: テストかく
	// TODO: コード綺麗にする
	// TODO: エラーどうするとよさそうか
	cmd := handler.Command()
	cmd.Execute()
}
