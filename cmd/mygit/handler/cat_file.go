package handler

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/git-starter-go/cmd/mygit/usecase"
	"github.com/spf13/cobra"
)

func CatFile() *cobra.Command {
	var hash string
	cmd := &cobra.Command{
		Use:   "cat-file",
		Short: "cat-file",
		Long:  "cat-file",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if err := usecase.CatFile(&hash); err != nil {
				fmt.Fprintf(os.Stderr, "%s", err)
			}
		},
	}

	const pFlagUsage = "usage: git cat-file (-t [--allow-unknown-type] | -s [--allow-unknown-type] | -e | -p | <type> | --textconv | --filters) [--path=<path>] <object>"
	cmd.Flags().StringVarP(&hash, "p", "p", "", pFlagUsage)

	return cmd
}
