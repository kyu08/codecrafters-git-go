package main

import (
	"os"

	"github.com/codecrafters-io/git-starter-go/cmd/mygit/handler"
)

func main() {
	handler.Handler(os.Args)
}
