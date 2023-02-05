package usecase

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type CatFileParam struct {
	Hash *string
}

func (p CatFileParam) validate() error {
	const hashLen = 40

	if len(*p.Hash) != hashLen {
		return errors.New("invalid hash format.")
	}
	return nil
}

func CatFile(write io.Writer, param CatFileParam) error {
	if err := param.validate(); err != nil {
		return err
	}

	// hashをファイルパスに変換
	// TODO: add test?
	filePath := fmt.Sprintf(".git/objects/%s/%s", (*param.Hash)[:2], (*param.Hash)[2:])

	// ファイル内容を取得
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("fail: read file: %w", err)
	}

	result, err := unzip(b)
	if err != nil {
		return fmt.Errorf("fail: unzipLines: %w", err)
	}

	write.Write([]byte(result))
	return nil
}

// zlibで圧縮されたバイト列を解凍
// TODO: add test
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
