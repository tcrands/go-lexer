package main

// Token represents a lexical token.
type Token string

// type Token int

// const (
// 	// Special tokens
// 	ILLEGAL Token = iota
// 	EOF
// 	WS
// 	// Literals
// 	IDENT // main
// 	// Misc characters
// 	ASTERISK // *
// 	COMMA    // ,
//
// 	// KEYWORDS
// 	FUNCTION
// )
//
// const (
// 	// Special tokens
// 	ILLEGAL Token = iota
// 	EOF
// 	WS
// 	// Literals
// 	IDENT
// 	STRING
// 	DIGIT
// 	// Misc characters
// 	ASTERISK
// 	COMMA
// 	OPENBRACKET
// 	CLOSEBRACKET
// 	OPENCURLYBRACKET
// 	CLOSECURLYBRACKET
// 	COLON
// 	SEMICOLON
// 	NEWLINE
// 	TAB
// 	// Keywords
// 	FUNCTION
// 	PARAMS
// 	AS
// 	SUB
// 	END
// 	IF
// 	FOR
// 	WHILE
// 	RETURN
// 	THIS
// )

const (
	// Special tokens
	ILLEGAL = "illegal"
	EOF     = ""
	WS      = "whitespace"
	TAB     = "tab"
	NEWLINE = "new line"
	// Literals
	IDENT  = "ident"
	STRING = "string"
	DIGIT  = "int"
	// Misc characters
	ASTERISK          = "asterisk"
	COMMA             = "comma"
	OPENBRACKET       = "open bracket"
	CLOSEBRACKET      = "close bracket"
	OPENCURLYBRACKET  = "open curly bracket"
	CLOSECURLYBRACKET = "close curly bracket"
	COLON             = "colon"
	SEMICOLON         = "semi colon"
	DOT               = "dot"
	UNDERSCORE        = "underscore"
	EQUALS            = "equals"
	// Keywords
	FUNCTION = "function"
	PARAMS   = "params"
	AS       = "param definer"
	SUB      = "subroutine"
	END      = "block ender"
	IF       = "if statement"
	FOR      = "for statement"
	WHILE    = "while statement"
	RETURN   = "return statement"
	THIS     = "this"
)
