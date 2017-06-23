//////////////////////////////////////////
//////////// LEXICAL SCANNER /////////////
//////////////////////////////////////////

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// Scanner struct
type Scanner struct {
	r     *bufio.Reader
	items chan item
}

// Custom type to allow passing of rune checking functions
type fn func(rune) bool

type item struct {
	item itemType
	val  string
}

// Array to haold the result
var LexicalAnalyisResult [][]string
var LexicalAnalyisResult2 []*item

// NewScanner returns a new instance of the scanner struct
// @param r: The text to be scanned
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r:     bufio.NewReader(r),
		items: make(chan item),
	}
}

// Scan returns the result of the scanning
// @return LexicalAnalyisResult
func (s *Scanner) Scan() [][]string {
	// Peek at the next rune.
	peekCh, _ := s.peek()

	switch {
	case isWhitespace(peekCh):
		buffer := s.scanBuffer(isWhitespace, false)
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: WS2, val: buffer})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{WS, buffer})
	case isLetter(peekCh):
		buffer := s.scanBuffer(isLetter, false)
		isKeyword(buffer)
	case isDigit(peekCh):
		// buffer := s.scanBuffer(isDigit, false)
		// LexicalAnalyisResult = append(LexicalAnalyisResult, []string{DIGIT, buffer})
	case isString(peekCh):
		buffer := s.scanBuffer(isString, true)
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: STRING2, val: buffer})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{STRING, buffer})
	default:
		ch := s.read()
		isSpecialChar(ch)
	}

	_, checkNext := s.peek()
	if checkNext {
		s.Scan()
	} else {
		fmt.Println(LexicalAnalyisResult)
		//
		for i := range LexicalAnalyisResult2 {
			fmt.Println(LexicalAnalyisResult2[i].item)
			fmt.Println(LexicalAnalyisResult2[i].val)
		}
	}

	return LexicalAnalyisResult

}

// scanBuffer loop through the runes and build a token
// @params f: The checking function to be used
// @params b: A boolean to check whether the token to be build uses a start and end indentifer
// @return string: A string version of the token
func (s *Scanner) scanBuffer(f fn, b bool) string {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	//Keep looping and writing chars to buffer until char is not what we want.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !f(ch) {
			if b == true {
				_, _ = buf.WriteRune(ch)
			} else {
				s.unread()
				break
			}
		} else {
			_, _ = buf.WriteRune(ch)
			if b == true {
				break
			}
		}
	}

	return buf.String()
}

// read reads the next rune from the bufferred reader
// @return rune: The next rune in the buffered reader
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// peeks looks at the next rune from the buffered reader without progressing the reader
func (s *Scanner) peek() (rune, bool) {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof, false
	}
	s.unread()
	return ch, true
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

//Emit the token back to the client (Whoever calls the scanner)
func (s *Scanner) emit(t itemType, v string) {
	s.items <- item{item: t, val: v}
}

// eof represents a marker rune for the end of the reader.
var eof = rune(0)
