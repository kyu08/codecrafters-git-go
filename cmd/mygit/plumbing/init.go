package plumbing

import (
	"fmt"
	"os"
)

func GitInit() {
	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		// MEMO: 0755: rwxr-xr-x
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
		}
	}

	headFileContents := []byte("ref: refs/heads/master\n")
	if err := os.WriteFile(".git/HEAD", headFileContents, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
	}

	fmt.Println("Initialized git directory")
}