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
			panic("file name is empty")
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

	// file contentの圧縮
	blobFilePath := fmt.Sprintf(".git/objects/%s/%s", hash[:2], hash[2:])

	// TODO: bを圧縮する
	f, err := os.Create(blobFilePath)
	defer f.Close()

	buf := new(bytes.Buffer)
	zw := zlib.NewWriter(buf)
	zw.Close()

	if _, err := io.Copy(zw, f); err != nil {
		panic("io.Copy failed.")
	}

	byte := buf.Bytes()

	_, err = f.Write(byte)
	if err != nil {
		panic(err.Error())
	}

	return hash
}

// someFunc2 SHAをstdoutに出力
func someFunc2(sha string) {
	fmt.Print(sha)
}