package handler

import (
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	// TODO: テストだったらbufを、それ以外だったら標準出力をusecaseに渡す
	var rootCmd = &cobra.Command{Use: "mygit usage"}
	rootCmd.AddCommand(Init())
	rootCmd.AddCommand(CatFile())
	rootCmd.AddCommand(HashObject())
	rootCmd.AddCommand(LSTree())

	return rootCmd
}
