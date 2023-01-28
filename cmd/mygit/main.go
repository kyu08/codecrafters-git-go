package main

import (
	"github.com/codecrafters-io/git-starter-go/cmd/mygit/handler"
)

func main() {
	// TODO: validationを関数に切り出す
	// TODO: コード綺麗にする
	// TODO: エラーどうするとよさそうか
	// TODO: テストかく
	cmd := handler.Command()
	cmd.Execute()
}
