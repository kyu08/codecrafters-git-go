package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func catFile(opt, optValue *string) {
	switch *opt {
	case "-p":
		const hashLen = 40

		if optValue == nil {
			panic("optValue not given.")
		}

		// hashをファイルパスに変換
		blobHash := *optValue
		if len(blobHash) != hashLen {
			panic("invalid hash format.")
		}
		filePath := fmt.Sprintf(".git/objects/%s/%s", blobHash[:2], blobHash[2:])

		// ファイル内容を取得
		b, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("ReadFile failed: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("👺\n")
		// ファイル内容を解凍
		result, err := unzip(b)
		if err != nil {
			fmt.Printf("unzipLines failed: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf(result)
	}
}

func getElem(args []string, index int64) *string {
	if int64(len(args)) >= (index + 1) {
		return &args[index]
	}

	return nil
}

func unzip(b []byte) (string, error) {
	r, err := zlib.NewReader(bytes.NewReader(b))
	if err != nil {
		return "", err
	}

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	result := func() string {
		s := strings.Split(string(buf), "\x00")
		if len(s) == 1 {
			return s[0]
		}
		return s[1]
	}()

	return result, nil
}
