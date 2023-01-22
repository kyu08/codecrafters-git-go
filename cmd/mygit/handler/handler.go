package handler

import (
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/git-starter-go/cmd/mygit/usecase"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {

	var rootCmd = &cobra.Command{Use: "TODO: write usage"}
	rootCmd.AddCommand(Init())
	rootCmd.AddCommand(CatFile())

	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	rootCmd.AddCommand(cmdEcho)
	cmdEcho.AddCommand(cmdTimes)

	return rootCmd
}

// TODO: あとでけす
var cmdEcho = &cobra.Command{
	Use:   "echo [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, " "))
	},
}
var echoTimes int

var cmdTimes = &cobra.Command{
	Use:   "times [# times] [string to echo]",
	Short: "Echo anything to the screen more times",
	Long: `echo things multiple times back to the user by providing
a count and a string.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < echoTimes; i++ {
			fmt.Println("Echo: " + strings.Join(args, " "))
		}
	},
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
