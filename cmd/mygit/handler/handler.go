package handler

import (
	"github.com/spf13/cobra"
)

// fmt.Printで標準出力固定で出力するのではなく、fmt.Fprintfとかを使って出力先を外から渡すようにすればよりテスタブルになりそう
func Command() *cobra.Command {
	var rootCmd = &cobra.Command{Use: "mygit usage"}
	rootCmd.AddCommand(Init())
	rootCmd.AddCommand(CatFile())
	rootCmd.AddCommand(HashObject())
	rootCmd.AddCommand(LSTree())

	return rootCmd
}

func getElem(args []string, index int64) *string {
	if int64(len(args)) >= (index + 1) {
		return &args[index]
	}

	return nil
}
