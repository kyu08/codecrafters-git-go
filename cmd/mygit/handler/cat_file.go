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
			param := usecase.CatFileParam{Hash: &hash}
			if err := usecase.CatFile(os.Stdout, param); err != nil {
				fmt.Fprintf(os.Stderr, "%s", err)
			}
		},
	}

	cmd.Flags().StringVarP(
		&hash,
		"p",
		"p",
		"",
		"usage: git cat-file (-t [--allow-unknown-type] | -s [--allow-unknown-type] | -e | -p | <type> | --textconv | --filters) [--path=<path>] <object>",
	)

	return cmd
}
