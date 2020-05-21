Mini Go Lisp
============

[![Go Report Card](https://goreportcard.com/badge/github.com/d-tsuji/gosdlisp)](https://goreportcard.com/report/github.com/d-tsuji/gosdlisp)
[![Actions Status](https://github.com/d-tsuji/gosdlisp/workflows/test/badge.svg)](https://github.com/d-tsuji/gosdlisp/actions)

Mini lisp interpreter written in Go. It is implemented with reference to the [d-tsuji/SDLisp](https://github.com/d-tsuji/SDLisp) repository written in Java.

## Support

- System Functions
    - `car`
    - `cdr`
    - `cons`
    - `eq`
    - `if` 
    - Arithmetic operations(`+`, `-`, `*`, `/`)
    - comparative operation(`>`, `<`, `>=`, `<=`, `=`)
- Special
    - `symbol-function`
    - `quote` or `'`
    - `setq`
    - `defun`

## Usage

```
$ go run github.com/d-tsuji/gosdlisp/cmd/gosdlisp
```

## Some examples

```lisp
> (+ 1 2)
3
```

```lisp
> (cons 1 '(2 3))
(1 2 3)
```

```lisp
> (defun 1+ (n) (+ n 1))
1+
> (1+ 10)
11
```

```lisp
> (defun abs (n) (if (< n 0) (- 0 n) n))
ABS
> (abs -1)
1
```

```lisp
> (defun fact (n) (if (< n 1) 1 (* n (fact (- n 1)))))
FACT
> (fact 10)
3628800
```

```lisp
> (defun fib (n) (if (<= n 1) n (+ (fib (- n 1)) (fib (- n 2)))))
FIB
> (fib 11)
89
```

And see [eval_test.go](https://github.com/d-tsuji/gosdlisp/blob/master/eval_test.go) for other examples of how it works.
