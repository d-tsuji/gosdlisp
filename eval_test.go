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
		{"if", `(> 2 1)`, &Symbol{"T", nil, nil}},
		{"defun", `(defun fact (n) (1))`, &Symbol{
			"FACT", nil, &Cons{
				Car: &Symbol{"LAMBDA", nil, nil},
				Cdr: &Cons{
					&Cons{&Symbol{"N", nil, nil}, nil},
					&Cons{&Cons{&Integer{1}, nil}, nil},
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
