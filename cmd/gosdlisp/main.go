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
	fmt.Fprint(os.Stdout, "> ")
	scanner := bufio.NewScanner(os.Stdin)
	eval := gosdlisp.NewEval()

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "error occured: %v", r)
		}
	}()

	for scanner.Scan() {
		fmt.Fprint(os.Stdout, "> ")
		line := scanner.Text()
		if line == "" {
			continue
		}
		r := gosdlisp.NewReader(strings.NewReader(line))
		sexp, err := r.Read()
		if err != nil {
			fmt.Fprintln(os.Stderr, "read:", err)
			continue
		}
		if sexp == gosdlisp.NewSymbol("QUIT") || sexp == gosdlisp.NewSymbol("EXIT") {
			break
		}
		fmt.Fprintln(os.Stdout, eval.Evaluate(sexp))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Fprintln(os.Stdout, "bye!")
}
