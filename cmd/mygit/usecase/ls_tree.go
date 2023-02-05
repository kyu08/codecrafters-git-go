package usecase

import (
	"fmt"
	"io"
	"io/ioutil"
)

func LSTree(writer io.Writer, treeSha *string) error {
	// tree-shaをファイルパスに変換
	filePath := fmt.Sprintf(".git/objects/%s/%s", (*treeSha)[:2], (*treeSha)[2:])

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
}
