package gosdlisp

import (
	"bufio"
	"log"
	"os"
	"unicode/utf8"
)

var (
	defaultRuneBuffSize = 256
)

type Reader struct {
	runeBuffSize int
	runeBuff     []rune
	ru           rune
	line         string
	indexOfLine  int
	lineLength   int
	br           *bufio.Reader
}

func NewReader() *Reader {
	return &Reader{
		runeBuffSize: defaultRuneBuffSize,
		runeBuff:     make([]rune, defaultRuneBuffSize),
		ru:           0,
		line:         "",
		indexOfLine:  0,
		lineLength:   -1,
		br:           bufio.NewReader(os.Stdin),
	}
}

func (r *Reader) read() T {
	line, _, err := r.br.ReadLine()
	if err != nil {
		log.Fatalf("cannot read line: %v", err)
	}
	r.line = string(line)
}

func (r *Reader) prepare() {
	r.lineLength = utf8.RuneCountInString(r.line)
	r.runeBuff = []rune(r.line)
	// TODO
	// charBuff[lineLength] = '\0'; // 終了マーク
	// getChar();
}
