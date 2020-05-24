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
			sexp, err := r.Read()
			if err != nil {
				t.Errorf("read: %v", err)
			}
			got, err := e.Evaluate(sexp)
			if err != nil {
				t.Errorf("evaluate: %v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Evaluate() differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestEval_EvaluateValue(t *testing.T) {
	tests := []struct {
		name  string
		cmds  []string
		wants []string
	}{
		{"car", []string{`(car '(a b c))`}, []string{`A`}},
		{"car", []string{`(cdr '(a b c))`}, []string{`(B C)`}},
		{"cons", []string{`(cons 1 2)`}, []string{`(1 . 2)`}},
		{"cons", []string{`(cons 3 (cons 1 2))`}, []string{`(3 1 . 2)`}},
		{"quote", []string{`(quote (+ 1 2 3))`}, []string{`(+ 1 2 3)`}},
		{"quote", []string{`'(+ 1 2 3)`}, []string{`(+ 1 2 3)`}},
		{"setq", []string{`(setq x 1)`, `(+ x 2)`}, []string{"X", "3"}},
		{"symbol-function", []string{`(defun x (n) (+ n 2))`, `(symbol-function 'x)`}, []string{"X", `(LAMBDA (N) (+ N 2))`}},
		{"defun", []string{`(defun zerop (n) (= n 0))`, `(zerop 0)`}, []string{"ZEROP", "T"}},
		{"defun", []string{`(defun 1+ (n) (+ n 1))`, `(1+ 10)`}, []string{"1+", "11"}},
		{"defun", []string{`(defun abs (n) (if (< n 0) (- 0 n) n))`, `(abs -1)`}, []string{"ABS", "1"}},
		{"defun", []string{`(defun gcd (m n) (if (= (mod m n) 0) n (gcd n (mod m n))))`, `(gcd 12 18)`}, []string{"GCD", "6"}},
		{"defun", []string{`(defun fact (n) (if (< n 1) 1 (* n (fact (- n 1)))))`, `(fact 10)`}, []string{"FACT", "3628800"}},
		{"defun", []string{`(defun fib (n) (if (<= n 1) n (+ (fib (- n 1)) (fib (- n 2)))))`, `(fib 11)`}, []string{"FIB", "89"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEval()
			for i := 0; i < len(tt.cmds); i++ {
				r, err := NewReader(strings.NewReader(tt.cmds[i])).Read()
				if err != nil {
					t.Errorf("read: %v", err)
				}
				got, err := e.Evaluate(r)
				if err != nil {
					t.Errorf("evaluate: %v", err)
				}
				if diff := cmp.Diff(got.String(), tt.wants[i]); diff != "" {
					t.Errorf("Evaluate() differs cmds[%d]: (-got +want)\n%s", i, diff)
				}
			}
		})
	}
}
