package gosdlisp

import (
	"bufio"
	"log"
	"os"
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

func NewReader() *Reader {
	return &Reader{
		ru:          0,
		line:        nil,
		indexOfLine: 0,
		lineLength:  -1,
		br:          bufio.NewReader(os.Stdin),
	}
}

func (r *Reader) read() T {
	line, _, err := r.br.ReadLine()
	if err != nil {
		log.Fatalf("cannot read line: %v", err)
	}
	r.line = []rune(string(line) + `\0`)
	return r.getSexp()
}

func (r *Reader) getRune() {
	r.ru = r.line[r.indexOfLine]
	r.indexOfLine++
}

func (r *Reader) getSexp() T {
	for {
		r.skipSpace()
		switch r.ru {
		case '(':
		default:
			if unicode.IsDigit(r.ru) {

			}
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

func (r *Reader) skipSpace() {
	for unicode.IsSpace(r.ru) {
		r.getRune()
	}
}
