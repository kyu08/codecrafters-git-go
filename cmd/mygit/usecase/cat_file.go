package usecase

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func CatFile(opt, optValue *string) error {
	switch *opt {
	case "-p":
		const hashLen = 40

		if optValue == nil {
			return errors.New("optValue not given.")
		}

		// hashをファイルパスに変換
		blobHash := *optValue
		if len(blobHash) != hashLen {
			return errors.New("invalid hash format.")
		}
		filePath := fmt.Sprintf(".git/objects/%s/%s", blobHash[:2], blobHash[2:])

		// ファイル内容を取得
		b, err := ioutil.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("fail: read file: %w", err)
		}

		result, err := unzip(b)
		if err != nil {
			return fmt.Errorf("fail: unzipLines: %w", err)
		}

		fmt.Printf(result)
		return nil
	default:
		return fmt.Errorf("invalid option: %s", *opt)
	}
}

// zlibで圧縮されたバイト列を解凍
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
