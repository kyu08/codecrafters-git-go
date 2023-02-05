package handler

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/git-starter-go/cmd/mygit/usecase"
	"github.com/spf13/cobra"
)

func HashObject() *cobra.Command {
	var filePath string
	cmd := &cobra.Command{
		Use:   "hash-object",
		Short: "hash-object",
		Long:  "hash-object",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			param := usecase.HashObjectParam{FilePath: &filePath}
			if err := usecase.HashObject(os.Stdout, param); err != nil {
				fmt.Fprintf(os.Stderr, "%s", err)
			}
		},
	}

	cmd.Flags().StringVarP(
		&filePath,
		"w",
		"w",
		"",
		"git hash-object [-t <type>] [-w] [--path=<file> | --no-filters] [--stdin [--literally]] [--] <file>…",
	)

	return cmd
}
