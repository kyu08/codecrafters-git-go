package main

import (
	"github.com/codecrafters-io/git-starter-go/cmd/mygit/handler"
)

func main() {
	// TODO: cobra移植やる
	// ls-treeだけ
	// TODO: validationを関数に切り出す
	// TODO: コード綺麗にする
	// TODO: テストかく
	cmd := handler.Command()
	cmd.Execute()
}
