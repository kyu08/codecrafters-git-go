package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
)

func hashObject(opt, optValue *string) {
	switch *opt {
	case "-w":
		if optValue == nil {
			panic("file name is empty")
		}

		someFunc1(*optValue)
		// someFunc2(sha)
	}
}

// someFunc1 データをblobとして.git/objectsに格納
func someFunc1(filePath string) string {
	// SHAの計算
	// ファイル内容の取得
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("ReadFile failed: %s\n", err)
		os.Exit(1)
	}

	content := string(b)
	header := fmt.Sprintf("blob %d\u0000", len(content))
	store := header + content

	// sha1を計算
	h := sha1.New()
	h.Write([]byte(store))
	bs := h.Sum(nil)
	fmt.Printf("%x\n", bs)

	// file contentの圧縮
	// headerとcontentを連結したやつをzlibで圧縮して格納
	fmt.Printf("store: %v\n", store)
	sha := "sha"
	return sha
}

// someFunc2 SHAをstdoutに出力
func someFunc2(sha string) {
	fmt.Print(sha)
}
