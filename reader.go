package gosdlisp

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type Reader struct {
	ru          rune
	line        []rune
	indexOfLine int
	lineLength  int
	br          *bufio.Reader
}

func NewReader(in io.Reader) *Reader {
	return &Reader{
		ru:          0,
		line:        nil,
		indexOfLine: 0,
		lineLength:  -1,
		br:          bufio.NewReader(in),
	}
}

// read is an S-expression reader.
func (r *Reader) read() T {
	line, _, err := r.br.ReadLine()
	if err != nil {
		log.Fatalf("cannot read line: %v", err)
	}
	r.line = []rune(string(line) + `\0`)
	return r.getSexp()
}

// getRune reads one character from the rune slice,
// sets a value to rune, and proceeds to indexOfLine.
func (r *Reader) getRune() {
	r.ru = r.line[r.indexOfLine]
	r.indexOfLine++
}

// getSexp reads an S-expression.
func (r *Reader) getSexp() T {
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
func (r *Reader) makeNumber() T {
	var str strings.Builder
	if r.ru == '-' {
		str.WriteRune('-')
		r.getRune()
	}
	for ; r.indexOfLine <= r.lineLength; r.getRune() {
		if r.ru == '(' || r.ru == ')' {
			break
		}
		if !unicode.IsDigit(r.ru) {
			r.indexOfLine--
			return r.makeSymbolInternal(str)
		}
		str.WriteRune(r.ru)
	}
	value, err := strconv.Atoi(str.String())
	if err != nil {
		log.Fatalf("cannot convert int: %v", err)
	}
	return NewInteger(value)
}

// makeMinusNumber reads a negative number.
func (r *Reader) makeMinusNumber() T {
	nru := r.line[r.indexOfLine]
	if !unicode.IsDigit(nru) {
		var str strings.Builder
		str.WriteRune(r.ru)
		return r.makeSymbolInternal(str)
	}
	return r.makeNumber()
}

// makeSymbol reads a symbol.
func (r *Reader) makeSymbol() T {
	r.ru = unicode.ToUpper(r.ru)
	var str strings.Builder
	str.WriteRune(r.ru)
	return r.makeSymbolInternal(str)
}

// makeSymbolInternal reads a symbol in the middle of a string.
func (r *Reader) makeSymbolInternal(str strings.Builder) T {
	for r.indexOfLine < r.lineLength {
		r.getRune()
		if r.ru == '(' || r.ru == ')' {
			break
		}
		if unicode.IsSpace(r.ru) {
			break
		}
		r.ru = unicode.ToUpper(r.ru)
		str.WriteRune(r.ru)
	}

	symStr := "" + str.String()

	if symStr == "NIL" {
		return Null{}
	}
	return NewSymbol(symStr)
}

// makeList reads the list.
func (r *Reader) makeList() T {
	r.getRune()
	r.skipSpace()
	if r.ru == ')' {
		r.getRune()
		return &Null{}
	}
	top := NewCons(&Null{}, &Null{})
	list := top
	for {
		list.car = r.getSexp()
		r.skipSpace()
		if r.indexOfLine > r.lineLength {
			return &Null{}
		}
		if r.ru == ')' {
			break
		}
		if r.ru == '.' {
			r.getRune()
			list.cdr = r.getSexp()
			r.skipSpace()
			r.getRune()
			return top
		}
		list.cdr = NewCons(&Null{}, &Null{})
		l, ok := list.cdr.(Cons)
		if !ok {
			log.Fatalf("cannot convert Cons: %v", list.cdr)
		}
		list.cdr = l
	}
	r.getRune()
	return top
}

// makeQuote reads the quote.
func (r *Reader) makeQuote() T {
	// TODO
	return nil
}

// SkipSpace skips whitespace.
func (r *Reader) skipSpace() {
	for unicode.IsSpace(r.ru) {
		r.getRune()
	}
}
