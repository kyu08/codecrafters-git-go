package handler

import (
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var rootCmd = &cobra.Command{Use: "mygit usage"}
	rootCmd.AddCommand(Init())
	rootCmd.AddCommand(CatFile())
	rootCmd.AddCommand(HashObject())
	rootCmd.AddCommand(LSTree())

	return rootCmd
}
