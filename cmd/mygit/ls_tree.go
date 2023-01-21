package main

import (
	"fmt"
	"os"
)

func lsTree(opt, optValue *string) {
	switch *opt {
	case "--name-only":
		fmt.Fprintf(os.Stderr, "sha: %s\n", *optValue)
		// tree-shaをファイルパスに変換
		// treeオブジェクトのファイル内容を取得
		// blobならファイル名をtreeなら再帰で処理
		// アルファベット順にソート
		// 標準出力に出力
	default:
		fmt.Fprintf(os.Stderr, "Invalid option %s\n", *opt)
		os.Exit(1)
	}
}
