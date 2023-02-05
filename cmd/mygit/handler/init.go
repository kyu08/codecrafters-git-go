package handler

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/git-starter-go/cmd/mygit/usecase"
	"github.com/spf13/cobra"
)

func Init() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "init",
		Long:  "init",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if err := usecase.Init(os.Stdout); err != nil {
				fmt.Fprintf(os.Stderr, "%s", err)
			}
		},
	}
}
