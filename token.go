package main

// Token represents a lexical token.
type Token string

type itemType int

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
const (
	// Special tokens
	ILLEGAL2 itemType = iota
	EOF2
	WS2
	// Literals
	IDENT2
	STRING2
	DIGIT2
	// Misc characters
	ASTERISK2
	COMMA2
	OPENBRACKET2
	CLOSEBRACKET2
	OPENCURLYBRACKET2
	CLOSECURLYBRACKET2
	COLON2
	SEMICOLON2
	NEWLINE2
	TAB2
	EQUALS2
	UNDERSCORE2
	DOT2
	// Keywords
	FUNCTION2
	PARAMS2
	AS2
	SUB2
	END2
	IF2
	FOR2
	WHILE2
	RETURN2
	THIS2
)

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
