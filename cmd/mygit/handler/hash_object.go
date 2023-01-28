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
			if err := usecase.HashObject(&filePath); err != nil {
				fmt.Fprintf(os.Stderr, "%s", err)
			}
		},
	}

	const flagUsage = "git hash-object [-t <type>] [-w] [--path=<file> | --no-filters] [--stdin [--literally]] [--] <file>…"
	cmd.Flags().StringVarP(&filePath, "w", "w", "", flagUsage)

	return cmd
}
