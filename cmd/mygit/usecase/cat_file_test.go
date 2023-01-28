package usecase

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCatFileParamValidate(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		Hash string
		Want error
	}{
		"40文字のときnilを返す": {
			Hash: "1234567890123456789012345678901234567890",
			Want: nil,
		},
		"39文字のときエラーを返す": {
			Hash: "123456789012345678901234567890123456789",
			Want: errors.New("invalid hash format."),
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			p := CatFileParam{
				hash: &tt.Hash,
			}
			got := p.validate()
			want := tt.Want
			if diff := compareError(t, want, got); diff != "" {
				t.Errorf("mismatch. (-expect +got)\n%s", diff)
			}
		})
	}
}

// 必要になったら適切な場所にうつした方がよさそう
func compareError(t *testing.T, want, got error) string {
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
