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
		{"eq", `(eq 'a 'a)`, NewSymbol("T")},
		{"eq", `(eq 'a 'b)`, nil},
		{"eq", `(eq 1 1)`, NewSymbol("T")},
		{"eq", `(eq 1 2)`, nil},
		{"if", `(> 2 1)`, NewSymbol("T")},
		{"if", `(> 2 2)`, nil},
		{"if", `(< 1 2)`, NewSymbol("T")},
		{"if", `(< 2 2)`, nil},
		{"if", `(>= 2 2)`, NewSymbol("T")},
		{"if", `(<= 2 2)`, NewSymbol("T")},
		{"quote", `(cons 1 '(2 3))`, &Cons{
			&Integer{1}, &Cons{&Integer{2}, &Cons{&Integer{3}, nil}},
		}}, // (1 2 3)
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

func TestEval_EvaluateFunc(t *testing.T) {
	tests := []struct {
		name    string
		defun   string
		execCmd string
		want    T
	}{
		{"defun", `(defun 1+ (n) (+ n 1))`, `(1+ 10)`, &Integer{11}},
		{"defun", `(defun abs (n) (if (< n 0) (- 0 n) n))`, `(abs -1)`, &Integer{1}},
		{"defun", `(defun fact (n) (if (< n 1) 1 (* n (fact (- n 1)))))`, `(fact 10)`, &Integer{3628800}},
		{"defun", `(defun fib (n) (if (<= n 1) n (+ (fib (- n 1)) (fib (- n 2)))))`, `(fib 11)`, &Integer{89}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEval()
			r := NewReader(strings.NewReader(tt.defun))
			e.Evaluate(r.Read())

			got := e.Evaluate(NewReader(strings.NewReader(tt.execCmd)).Read())
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Evaluate() differs: (-got +want)\n%s", diff)
			}
		})
	}
}
