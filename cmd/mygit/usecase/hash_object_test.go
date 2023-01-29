package usecase

import (
	"errors"
	"github.com/codecrafters-io/git-starter-go/cmd/mygit/testable"
	"testing"
)

func TestHashObjectParamValidate(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		FilePath *string
		Want     error
	}{
		"非nilのときnilを返す": {
			FilePath: testable.StrToPtr("/path/to/file.txt"),
			Want:     nil,
		},
		"nilのときエラーを返す": {
			FilePath: nil,
			Want:     errors.New("file name is empty"),
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			p := HashObjectParam{
				FilePath: tt.FilePath,
			}
			got := p.validate()
			want := tt.Want
			if diff := testable.CompareError(t, want, got); diff != "" {
				t.Errorf("mismatch. (-expect +got)\n%s", diff)
			}
		})
	}
}
