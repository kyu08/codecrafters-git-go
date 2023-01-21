package usecase

import (
	"fmt"
	"io/ioutil"
)

func LsTree(opt, optValue *string) error {
	switch *opt {
	case "--name-only":
		// tree-shaをファイルパスに変換
		hash := *optValue
		filePath := fmt.Sprintf(".git/objects/%s/%s", hash[:2], hash[2:])

		// ファイル内容を取得
		b, err := ioutil.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("ReadFile failed: %w", err)
		}

		// 解凍
		fmt.Printf("string(b): %v\n", string(b))

		// treeオブジェクトのファイル内容を取得
		// blobならファイル名をtreeなら再帰で処理
		// アルファベット順にソート
		// 標準出力に出力
		return nil
	default:
		return fmt.Errorf("Invalid option %s", *opt)
	}
}
