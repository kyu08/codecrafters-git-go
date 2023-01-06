package main

import (
	"testing"
)

func Test_blobHashToFilePath(t *testing.T) {
	t.Parallel()
	hogefugaResult := "ho/gefuga"
	testTxtResult := "te/st.txt"

	tests := map[string]struct {
		hash string
		want *string
	}{
		"hogefuga": {
			hash: "hogefuga",
			want: &hogefugaResult,
		},
		"test.txt": {
			hash: "test.txt",
			want: &testTxtResult,
		},
		"a": {
			hash: "a",
			want: nil,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := blobHashToFilePath(tt.hash)

			if tt.want == nil {
				if got == nil {
					return
				}
				t.Errorf("blobHashToFilePath(): %v, want: %v", *got, tt.want)
			}

			if *got != *tt.want {
				t.Errorf("blobHashToFilePath(): %v, want: %v", *got, *tt.want)
			}
		})
	}
}
func Test_getElem(t *testing.T) {
	t.Parallel()

	hoge := "hoge"

	tests := map[string]struct {
		args  []string
		index int64
		want  *string
	}{
		"存在するインデックス": {
			args:  []string{"hoge"},
			index: 0,
			want:  &hoge,
		},
		"存在しないインデックス": {
			args:  []string{"hoge"},
			index: 1,
			want:  nil,
		},
		"空配列": {
			args:  []string{},
			index: 1,
			want:  nil,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := getElem(tt.args, tt.index)

			if tt.want == nil {
				if got == nil {
					return
				}
				t.Errorf("got: %v, want: %v", *got, tt.want)
			}

			if *got != *tt.want {
				t.Errorf("got: %v, want: %v", *got, *tt.want)
			}
		})
	}
}
