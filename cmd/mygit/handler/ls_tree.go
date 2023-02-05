package handler

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/git-starter-go/cmd/mygit/usecase"
	"github.com/spf13/cobra"
)

func LSTree() *cobra.Command {
	var treeSha string
	cmd := &cobra.Command{
		Use:   "ls-tree",
		Short: "ls-tree",
		Long:  "ls-tree",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if err := usecase.LSTree(os.Stdout, &treeSha); err != nil {
				fmt.Fprintf(os.Stderr, "%s", err)
			}
		},
	}

	cmd.Flags().StringVarP(
		&treeSha,
		"name-only",
		"n",
		"",
		"git ls-tree [-d] [-r] [-t] [-l] [-z] [--name-only] [--name-status] [--object-only] [--full-name] [--full-tree] [--abbrev[=<n>]] [--format=<format>] <tree-ish> [<path>…]",
	)

	return cmd
}
