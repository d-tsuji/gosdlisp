package gosdlisp

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEval_Evaluate(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  T
	}{
		{"add", `(+ 1 2)`, &Integer{3}},
		{"sub", `(- 1 2)`, &Integer{-1}},
		{"mul", `(* 1 2)`, &Integer{2}},
		{"div", `(/ 10 2)`, &Integer{5}},
		{"defun", `(defun fact (n) (1))`, &Symbol{
			Name:  "FACT",
			Value: nil,
			Function: &Cons{
				Car: &Symbol{
					Name:     "LAMBDA",
					Value:    nil,
					Function: nil,
				},
				Cdr: &Cons{
					Car: &Cons{
						Car: &Symbol{
							Name:     "N",
							Value:    nil,
							Function: nil,
						},
						Cdr: nil,
					},
					Cdr: &Cons{
						Car: &Cons{
							Car: &Integer{1},
							Cdr: nil,
						},
						Cdr: nil,
					},
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEval()
			r := NewReader(strings.NewReader(tt.input))
			sexp := r.Read()
			got := e.Evaluate(sexp)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Evaluate() differs: (-got +want)\n%s", diff)
			}
		})
	}
}
