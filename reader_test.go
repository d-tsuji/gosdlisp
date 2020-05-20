package gosdlisp

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReader_read(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  T
	}{
		{name: "", input: `(+ 1 2)`, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewReader(strings.NewReader(tt.input))
			got := r.read()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("read() differs: (-got +want)\n%s", diff)
			}
		})
	}
}
