package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// Scanner struct
type Scanner struct {
	r *bufio.Reader
}

type fn func(rune) bool

var LexicalAnalyisResult [][]string

// NewScanner returns a new instance of the scanner struct
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// Scan returns the token and lit value
func (s *Scanner) Scan() [][]string {
	// Peek at the next rune.
	peekCh, _ := s.peek()

	// Check if whitespace/string/int
	// if isWhitespace(peekCh) {
	// 	buffer := s.scanBuffer(isWhitespace, false)
	// 	LexicalAnalyisResult = append(LexicalAnalyisResult, []string{WS, buffer})
	// 	// s.scanWhitespace()
	// } else if isLetter(peekCh) {
	// 	// s.scanIdent()
	// 	buffer := s.scanBuffer(isLetter, false)
	// 	isKeyword(buffer)
	// } else if isDigit(peekCh) {
	// 	// return s.scanDigits()
	// } else if isString(peekCh) {
	// 	// s.scanString()
	// 	buffer := s.scanBuffer(isString, true)
	// 	LexicalAnalyisResult = append(LexicalAnalyisResult, []string{STRING, buffer})
	// } else {
	// 	ch := s.read()
	// 	isSpecialChar(ch)
	// }

	switch {
		case isWhitespace(peekCh):
			buffer := s.scanBuffer(isWhitespace, false)
			LexicalAnalyisResult = append(LexicalAnalyisResult, []string{WS, buffer})
		case isLetter(peekCh):
			buffer := s.scanBuffer(isLetter, false)
			isKeyword(buffer) 
		case isDigit(peekCh):
			// return s.scanDigits() 
		case isString(peekCh): 
			buffer := s.scanBuffer(isString, true)
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
	}

	return LexicalAnalyisResult

}

func (s *Scanner) scanBuffer(f fn, b bool) string {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	//Keep looping and writing chars to buffer until char is not a int.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !f(ch) {
			if b == true {
				_, _ = buf.WriteRune(ch)
				// break
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

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// Peeks at the rune from the buffered reader
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

// eof represents a marker rune for the end of the reader.
var eof = rune(0)

// func (s *Scanner) scanWhitespace() {
// 	// Create a buffer and read the current character into it.
// 	var buf bytes.Buffer
// 	buf.WriteRune(s.read())
//
// 	//Keep looping and writing chars to buffer until char is not whitepace
// 	for {
// 		if ch := s.read(); ch == eof {
// 			break
// 		} else if !isWhitespace(ch) {
// 			s.unread()
// 			break
// 		} else {
// 			buf.WriteRune(ch)
// 		}
// 	}
// 	// Return the buffer
// 	LexicalAnalyisResult = append(LexicalAnalyisResult, []string{STRING, " "})
// }
//
// func (s *Scanner) scanIdent() {
// 	// Create a buffer and read the current character into it.
// 	var buf bytes.Buffer
// 	buf.WriteRune(s.read())
//
// 	//Keep looping and writing chars to buffer until char is not a letter.
// 	for {
// 		if ch := s.read(); ch == eof {
// 			break
// 		} else if !isLetter(ch) && ch != '_' {
// 			s.unread()
// 			break
// 		} else {
// 			_, _ = buf.WriteRune(ch)
// 		}
// 	}
//
// 	isKeyword(buf.String())
// }
//
// func (s *Scanner) scanString() {
// 	// Create a buffer and read the current character into it.
// 	var buf bytes.Buffer
// 	buf.WriteRune(s.read())
//
// 	//Keep looping and writing chars to buffer until char is not a letter.
// 	for {
// 		if ch := s.read(); ch == eof {
// 			break
// 		} else if isString(ch) {
// 			_, _ = buf.WriteRune(ch)
// 			break
// 		} else {
// 			_, _ = buf.WriteRune(ch)
// 		}
// 	}
//
// 	LexicalAnalyisResult = append(LexicalAnalyisResult, []string{STRING, buf.String()})
// }

func (s *Scanner) scanDigits() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	//Keep looping and writing chars to buffer until char is not a int.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isDigit(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// Otherwise return as a regular identifier.
	return DIGIT, buf.String()
}
