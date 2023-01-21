package main

import (
	"os"

	"github.com/codecrafters-io/git-starter-go/cmd/mygit/handler"
)

func main() {
	// fmt.Printで標準出力固定で出力するのではなく、fmt.Fprintfとかを使って出力先を外から渡すようにすればよりテスタブルになりそう
	// コマンドライン引数の受け取りにcobraを使ってみたい
	handler.Handler(os.Args)
}
