package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"os"
)

func hashObject(opt, optValue *string) {
	switch *opt {
	case "-w":
		if optValue == nil {
			fmt.Print("file name is empty")
			os.Exit(1)
		}

		// TODO: hashString取得メソッドを別定義する
		hash := someFunc1(*optValue)
		someFunc2(hash)
	}
}

// someFunc1 データをblobとして.git/objectsに格納
func someFunc1(sourceFilePath string) string {
	// ファイル内容の取得
	sourceFile, err := os.Open(sourceFilePath)
	defer sourceFile.Close()
	if err != nil {
		fmt.Printf("os.Open failed: %s\n", err)
		os.Exit(1)
	}

	contentByte := make([]byte, 1024)
	count, err := sourceFile.Read(contentByte)
	if err != nil {
		fmt.Printf("sourceFile.Read failed: %s\n", err)
		os.Exit(1)
	}

	contentStr := string(contentByte[:count])
	header := fmt.Sprintf("blob %d\x00", len(contentStr))
	store := header + contentStr

	// sha1を計算
	h := sha1.New()
	h.Write([]byte(store))
	bs := h.Sum(nil)
	hash := fmt.Sprintf("%x", bs)
	dirName := hash[:2]
	fileName := hash[2:]

	// file contentの圧縮
	dirPath := fmt.Sprintf(".git/objects/%s", dirName)
	blobFilePath := fmt.Sprintf("%s/%s", dirPath, fileName)

	if err := os.MkdirAll(dirPath, 0777); err != nil {
		fmt.Printf("os.MkdirAll failed. err:%s", err)
		os.Exit(1)
	}

	f, err := os.Create(blobFilePath)
	defer f.Close()
	if err != nil {
		fmt.Printf("os.Create failed. err:%s", err)
		os.Exit(1)
	}

	var buf bytes.Buffer
	zw := zlib.NewWriter(&buf)
	defer zw.Close()

	if _, err = zw.Write(contentByte[:count]); err != nil {
		fmt.Printf("zw.Write failed. err:%s", err)
		os.Exit(1)
	}

	// ↑と↓まとめれたりする？
	// zwの圧縮 & 書き込み周りがおかしそうなので修正する
	if count, err = f.Write(buf.Bytes()); err != nil {
		fmt.Printf("f.Write failed. err:%s", err)
		os.Exit(1)
	}

	return hash
}

// someFunc2 SHAをstdoutに出力
func someFunc2(sha string) {
	fmt.Print(sha)
}
