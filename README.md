Mini Go Lisp
============

[![Go Report Card](https://goreportcard.com/badge/github.com/d-tsuji/gosdlisp)](https://goreportcard.com/report/github.com/d-tsuji/gosdlisp)
[![Actions Status](https://github.com/d-tsuji/gosdlisp/workflows/test/badge.svg)](https://github.com/d-tsuji/gosdlisp/actions)

Mini lisp interpreter written in Go. It is implemented with reference to the [d-tsuji/SDLisp](https://github.com/d-tsuji/SDLisp) repository written in Java.

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
> (defun fact (n) (if (< n 1) 1 (* n (fact (- n 1)))))
FACT
> (fact 10)
3628800
```
