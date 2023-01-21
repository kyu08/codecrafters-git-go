package usecase

import (
	"fmt"
	"os"
)

func Init() error {
	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		// 0755: rwxr-xr-x
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("fail: create directory: %w", err)
		}
	}

	headFileContents := []byte("ref: refs/heads/master\n")
	if err := os.WriteFile(".git/HEAD", headFileContents, 0644); err != nil {
		return fmt.Errorf("fail: write file: %w", err)
	}

	fmt.Println("Initialized git directory")
	return nil
}
