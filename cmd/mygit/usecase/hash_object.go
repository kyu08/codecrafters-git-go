package usecase

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"os"
)

type HashObjectParam struct {
	FilePath *string
}

func (p HashObjectParam) validate() error {
	if p.FilePath == nil {
		return errors.New("file name is empty")
	}
	return nil
}

func HashObject(writer io.Writer, param HashObjectParam) error {
	if err := param.validate(); err != nil {
		return err
	}

	store, err := getStore(*param.FilePath)
	if err != nil {
		return err
	}

	hash := getHash(store)
	if err = saveBlob(store, hash); err != nil {
		return err
	}

	writer.Write([]byte(hash))
	return nil
}

func getStore(sourceFilePath string) (string, error) {
	// ファイル内容の取得
	sourceFile, err := os.Open(sourceFilePath)
	defer sourceFile.Close()
	if err != nil {
		return "", fmt.Errorf("os.Open failed: %w", err)
	}

	// ファイル内容の読み出し
	contentByte := make([]byte, 1024)
	count, err := sourceFile.Read(contentByte)
	if err != nil {
		return "", fmt.Errorf("sourceFile.Read failed: %w", err)

	}

	// headerを計算
	contentStr := string(contentByte[:count])
	header := fmt.Sprintf("blob %d\x00", len(contentStr))

	return header + contentStr, nil
}

// TODO: add test
func getHash(store string) string {
	h := sha1.New()
	h.Write([]byte(store))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

// saveBlob blobデータを圧縮して.git/objectsに格納
// TODO: 責務が大きすぎるので分割する
func saveBlob(store, hash string) error {
	// file contentの圧縮
	dirPath := fmt.Sprintf(".git/objects/%s", hash[:2])
	blobFilePath := fmt.Sprintf("%s/%s", dirPath, hash[2:])

	if err := os.MkdirAll(dirPath, 0777); err != nil {
		return fmt.Errorf("os.MkdirAll failed. err:%w", err)
	}

	f, err := os.Create(blobFilePath)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("os.Create failed. err:%w", err)
	}

	// 圧縮
	var buf bytes.Buffer
	zw := zlib.NewWriter(&buf)
	zw.Write([]byte(store))
	zw.Close()

	// .git/objects以下にファイル書き込み
	if _, err := f.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("f.Write failed. err:%w", err)
	}

	return nil
}
