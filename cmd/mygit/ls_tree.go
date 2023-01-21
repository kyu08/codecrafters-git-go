package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func lsTree(opt, optValue *string) {
	switch *opt {
	case "--name-only":
		// tree-shaをファイルパスに変換
		hash := *optValue
		filePath := fmt.Sprintf(".git/objects/%s/%s", hash[:2], hash[2:])

		// ファイル内容を取得
		b, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("ReadFile failed: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("string(b): %v\n", string(b))

		// treeオブジェクトのファイル内容を取得
		// blobならファイル名をtreeなら再帰で処理
		// アルファベット順にソート
		// 標準出力に出力
	default:
		fmt.Fprintf(os.Stderr, "Invalid option %s\n", *opt)
		os.Exit(1)
	}
}
