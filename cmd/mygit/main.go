package main

import (
	"github.com/codecrafters-io/git-starter-go/cmd/mygit/handler"
)

// TODO: usecase、標準出力固定で出力するのではなく、書き込み先を外から渡す構成にしてテストを書く
// TODO: その他TODOコメント消化する
func main() {
	cmd := handler.Command()
	cmd.Execute()
}
