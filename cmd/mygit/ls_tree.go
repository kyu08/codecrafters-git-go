package main

import (
	"fmt"
	"os"
)

func lsTree(opt, optValue *string) {
	switch *opt {
	case "ls-tree":
		fmt.Fprintf(os.Stderr, "Invalid option %s\n", *opt)
		// objectを取得
		// NOTE:git hash-objectでディレクトリも追加できるようにする必要がある？
		// アルファベット順にソート
		// 標準出力に出力
	default:
		fmt.Fprintf(os.Stderr, "Invalid option %s\n", *opt)
		os.Exit(1)
	}
}
