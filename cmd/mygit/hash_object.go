package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
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
	b, err := ioutil.ReadFile(sourceFilePath)
	if err != nil {
		fmt.Printf("ReadFile failed: %s\n", err)
		os.Exit(1)
	}

	content := string(b)
	header := fmt.Sprintf("blob %d\x00", len(content))
	store := header + content

	// sha1を計算
	h := sha1.New()
	h.Write([]byte(store))
	bs := h.Sum(nil)
	hash := fmt.Sprintf("%x\n", bs)
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
	if err != nil {
		fmt.Printf("os.Create failed. err:%s", err)
		os.Exit(1)
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	zw := zlib.NewWriter(buf)
	zw.Close()

	if _, err := io.Copy(zw, f); err != nil {
		fmt.Printf("io.Copy failed. err:%s", err)
		os.Exit(1)
	}

	byte := buf.Bytes()

	_, err = f.Write(byte)
	if err != nil {
		fmt.Printf("f.Write failed. err:%s", err)
		os.Exit(1)
	}

	return hash
}

// someFunc2 SHAをstdoutに出力
func someFunc2(sha string) {
	fmt.Print(sha)
}
