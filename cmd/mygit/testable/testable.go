package testable

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// 必要になったら適切な場所にうつした方がよさそう
func CompareError(t *testing.T, want, got error) string {
	t.Helper()
	wantStr, gotStr := "", ""
	if want != nil {
		wantStr = want.Error()
	}

	if got != nil {
		gotStr = got.Error()
	}

	return cmp.Diff(wantStr, gotStr)
}

func StrToPtr(str string) *string {
	return &str
}
