package handler

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/git-starter-go/cmd/mygit/usecase"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {

	var rootCmd = &cobra.Command{Use: "TODO: write usage"}
	rootCmd.AddCommand(Init())
	rootCmd.AddCommand(CatFile())

	return rootCmd
}

// fmt.Printで標準出力固定で出力するのではなく、fmt.Fprintfとかを使って出力先を外から渡すようにすればよりテスタブルになりそう
func Handler(args []string) error {
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: mygit <command> [<args>...]")
	}

	command := os.Args[1]
	opt := getElem(os.Args, 2)
	optValue := getElem(os.Args, 3)

	switch command {
	case "hash-object":

		// ↓これのcobra移植をするところから！！！！！！！！

		return usecase.HashObject(opt, optValue)
	case "ls-tree":
		return usecase.LsTree(opt, optValue)
	default:
		return fmt.Errorf("Unknown command %s", command)
	}
}

func getElem(args []string, index int64) *string {
	if int64(len(args)) >= (index + 1) {
		return &args[index]
	}

	return nil
}
