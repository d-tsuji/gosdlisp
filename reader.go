package gosdlisp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Reader struct {
	ru          rune
	line        []rune
	indexOfLine int
	lineLength  int
	br          *bufio.Reader
}

func NewReader(in io.Reader) *Reader { return &Reader{br: bufio.NewReader(in)} }

// read is an S-expression reader.
func (r *Reader) Read() (T, error) {
	line, _, err := r.br.ReadLine()
	if err != nil {
		return nil, fmt.Errorf("cannot read line: %v", err)
	}
	r.lineLength = utf8.RuneCountInString(string(line))
	r.line = []rune(string(line) + `\0`)
	r.getRune()
	return r.getSexp()
}

// getRune reads one character from the rune slice,
// sets a value to rune, and proceeds to indexOfLine.
func (r *Reader) getRune() {
	r.ru = r.line[r.indexOfLine]
	r.indexOfLine++
}

// getSexp reads an S-expression.
func (r *Reader) getSexp() (T, error) {
	for {
		r.skipSpace()
		switch r.ru {
		case '(':
			return r.makeList()
		case '\'':
			return r.makeQuote()
		case '-':
			return r.makeMinusNumber()
		default:
			if unicode.IsDigit(r.ru) {
				return r.makeNumber()
			}
			return r.makeSymbol()
		}
	}
}

// makeNumber reads the Number.
func (r *Reader) makeNumber() (T, error) {
	var str strings.Builder
	if r.ru == '-' {
		str.WriteRune('-')
		r.getRune()
	}
	for ; r.indexOfLine <= r.lineLength; r.getRune() {
		if r.ru == '(' || r.ru == ')' {
			break
		}
		if unicode.IsSpace(r.ru) {
			break
		}
		if !unicode.IsDigit(r.ru) {
			r.indexOfLine--
			return r.makeSymbolInternal(str.String()), nil
		}
		str.WriteRune(r.ru)
	}
	value, err := strconv.Atoi(str.String())
	if err != nil {
		return nil, fmt.Errorf("cannot convert int: %v", err)
	}
	return NewInteger(value), nil
}

// makeMinusNumber reads a negative number.
func (r *Reader) makeMinusNumber() (T, error) {
	nru := r.line[r.indexOfLine]
	if !unicode.IsDigit(nru) {
		var str strings.Builder
		str.WriteRune(r.ru)
		return r.makeSymbolInternal(str.String()), nil
	}
	return r.makeNumber()
}

// makeSymbol reads a symbol.
func (r *Reader) makeSymbol() (T, error) {
	r.ru = unicode.ToUpper(r.ru)
	var str strings.Builder
	str.WriteRune(r.ru)
	return r.makeSymbolInternal(str.String()), nil
}

// makeSymbolInternal reads a symbol in the middle of a string.
func (r *Reader) makeSymbolInternal(str string) T {
	for r.indexOfLine < r.lineLength {
		r.getRune()
		if r.ru == '(' || r.ru == ')' {
			break
		}
		if unicode.IsSpace(r.ru) {
			break
		}
		r.ru = unicode.ToUpper(r.ru)
		str += string(r.ru)
	}

	return NewSymbol(str)
}

// makeList reads the list.
func (r *Reader) makeList() (T, error) {
	r.getRune()
	r.skipSpace()
	if r.ru == ')' {
		r.getRune()
		return nil, nil
	}
	top := NewCons(nil, nil)
	list := top
	for {
		sexp, err := r.getSexp()
		if err != nil {
			return nil, err
		}
		list.Car = sexp
		r.skipSpace()
		if r.indexOfLine > r.lineLength {
			return nil, nil
		}
		if r.ru == ')' {
			break
		}
		if r.ru == '.' {
			r.getRune()
			sexp, err := r.getSexp()
			if err != nil {
				return nil, err
			}
			list.Cdr = sexp
			r.skipSpace()
			r.getRune()
			return top, nil
		}
		list.Cdr = NewCons(nil, nil)
		l, ok := list.Cdr.(*Cons)
		if !ok {
			return nil, fmt.Errorf("cannot convert Cons: %v", list.Cdr)
		}
		list = l
	}
	r.getRune()
	return top, nil
}

// makeQuote reads the quote.
func (r *Reader) makeQuote() (T, error) {
	top := NewCons(nil, nil)
	list := top
	list.Car = NewSymbol("QUOTE")
	list.Cdr = NewCons(nil, nil)
	l, ok := list.Cdr.(*Cons)
	if !ok {
		return nil, fmt.Errorf("cannot convert Cons: %v", list.Cdr)
	}
	list = l
	r.getRune()
	list.Car, _ = r.getSexp()
	return top, nil
}

// SkipSpace skips whitespace.
func (r *Reader) skipSpace() {
	for unicode.IsSpace(r.ru) {
		r.getRune()
	}
}
