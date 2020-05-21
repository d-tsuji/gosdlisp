package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/d-tsuji/gosdlisp"
)

func main() {
	fmt.Fprintln(os.Stdout, "Welcome to GoSDLisp!")
	fmt.Fprintln(os.Stdout, "> Copyright (C) Tsuji Daishiro 2020.")
	fmt.Fprintln(os.Stdout, "> Type quit and hit Enter for leaving GoSDLisp.")
	fmt.Fprintln(os.Stdout, "> ")
	scanner := bufio.NewScanner(os.Stdin)
	eval := gosdlisp.NewEval()

	for scanner.Scan() {
		fmt.Fprint(os.Stdout, "> ")
		line := scanner.Text()
		if line == "" {
			continue
		}
		r := gosdlisp.NewReader(strings.NewReader(line))
		sexp := r.Read()
		if sexp == gosdlisp.NewSymbol("QUIT") {
			break
		}
		fmt.Fprintln(os.Stdout, eval.Evaluate(sexp))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Fprintln(os.Stdout, "bye!")
}
