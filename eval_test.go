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
		{"car", `(car '(a b c))`, &Symbol{Name: "A"}},
		{"cdr", `(cdr '(a b c))`, &Cons{Car: &Symbol{Name: "B"}, Cdr: &Cons{Car: &Symbol{Name: "C"}}}},
		{"add", `(+ 1 2)`, &Integer{3}},
		{"add", `(+ 1 2)`, &Integer{3}},
		{"sub", `(- 1 2)`, &Integer{-1}},
		{"mul", `(* 1 2)`, &Integer{2}},
		{"div", `(/ 10 2)`, &Integer{5}},
		{"mod", `(mod 10 3)`, &Integer{1}},
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
		{"if", `(= 1 1)`, NewSymbol("T")},
		{"if", `(= 1 2)`, nil},
		{"cons", `(cons 1 '(2 3))`, &Cons{
			&Integer{1}, &Cons{&Integer{2}, &Cons{&Integer{3}, nil}},
		}},
		{"defun", `(defun fact (n) (1))`, &Symbol{
			"FACT", nil, &Cons{
				Car: &Symbol{Name: "LAMBDA"},
				Cdr: &Cons{&Cons{Car: &Symbol{Name: "N"}}, &Cons{Car: &Cons{Car: &Integer{1}}}},
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

func TestEval_EvaluateValue(t *testing.T) {
	tests := []struct {
		name     string
		firstCmd string
		nextCmd  string
		want     string
	}{
		{"car", `(car '(a b c))`, "", `A`},
		{"car", `(cdr '(a b c))`, "", `(B C)`},
		{"cons", `(cons 1 2)`, "", `(1 . 2)`},
		{"cons", `(cons 3 (cons 1 2))`, "", `(3 1 . 2)`},
		{"quote", `(quote (+ 1 2 3))`, "", `(+ 1 2 3)`},
		{"quote", `'(+ 1 2 3)`, "", `(+ 1 2 3)`},
		{"setq", `(setq x 1)`, `(+ x 2)`, "3"},
		{"symbol-function", `(defun x (n) (+ n 2))`, `(symbol-function 'x)`, `(LAMBDA (N) (+ N 2))`},
		{"defun", `(defun zerop (n) (= n 0))`, `(zerop 0)`, "T"},
		{"defun", `(defun 1+ (n) (+ n 1))`, `(1+ 10)`, "11"},
		{"defun", `(defun abs (n) (if (< n 0) (- 0 n) n))`, `(abs -1)`, "1"},
		{"defun", `(defun gcd (m n) (if (= (mod m n) 0) n (gcd n (mod m n))))`, `(gcd 12 18)`, "6"},
		{"defun", `(defun fact (n) (if (< n 1) 1 (* n (fact (- n 1)))))`, `(fact 10)`, "3628800"},
		{"defun", `(defun fib (n) (if (<= n 1) n (+ (fib (- n 1)) (fib (- n 2)))))`, `(fib 11)`, "89"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEval()
			r := NewReader(strings.NewReader(tt.firstCmd))
			sexp := r.Read()
			got := e.Evaluate(sexp)
			if tt.nextCmd != "" {
				got = e.Evaluate(NewReader(strings.NewReader(tt.nextCmd)).Read())
			}
			if diff := cmp.Diff(got.String(), tt.want); diff != "" {
				t.Errorf("Evaluate() differs: (-got +want)\n%s", diff)
			}
		})
	}
}
