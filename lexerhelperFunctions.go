package main

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool {
	return ch == ' '
}

// isLetter returns true if the rune is a letter.
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') && (ch < '0' && ch > '9')
}

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

func isString(ch rune) bool {
	return (ch == '"')
}

func isOperator(ch rune) bool {
	return (ch == '=' || ch == '>' || ch == '<')
}

func isKeyword(buffer string) {
	//Check than the string in the buffer is not a keyword
	switch buffer {
	case "function":
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: FUNCTION2, val: buffer})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{FUNCTION, buffer})
	case "as":
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: AS2, val: buffer})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{AS, buffer})
	case "end":
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: END2, val: buffer})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{END, buffer})
	case "sub":
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: SUB2, val: buffer})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{SUB, buffer})
	case "for":
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: FOR2, val: buffer})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{FOR, buffer})
	case "if":
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: IF2, val: buffer})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{IF, buffer})
	case "while":
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: WHILE2, val: buffer})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{WHILE, buffer})
	case "return":
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: RETURN2, val: buffer})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{RETURN, buffer})
	case "this":
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: THIS2, val: buffer})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{THIS, buffer})
	case "object", "int", "string", "void":
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: PARAMS2, val: buffer})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{PARAMS, buffer})
	default:
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: IDENT2, val: buffer})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{IDENT, buffer})
	}

}

func isSpecialChar(ch rune) {
	//Process special Chars
	switch ch {
	case '*':
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: ASTERISK2, val: string(ch)})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{ASTERISK, string(ch)})
	case ',':
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: COMMA2, val: string(ch)})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{COMMA, string(ch)})
	case ':':
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: COLON2, val: string(ch)})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{COLON, string(ch)})
	case ';':
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: SEMICOLON2, val: string(ch)})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{SEMICOLON, string(ch)})
	case '(':
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: OPENBRACKET2, val: string(ch)})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{OPENBRACKET, string(ch)})
	case ')':
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: CLOSEBRACKET2, val: string(ch)})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{CLOSEBRACKET, string(ch)})
	case '{':
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: OPENCURLYBRACKET2, val: string(ch)})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{OPENCURLYBRACKET, string(ch)})
	case '}':
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: CLOSECURLYBRACKET2, val: string(ch)})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{CLOSECURLYBRACKET, string(ch)})
	case '.':
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: DOT2, val: string(ch)})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{DOT, string(ch)})
	case '_':
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: UNDERSCORE2, val: string(ch)})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{UNDERSCORE, string(ch)})
	case '=':
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: EQUALS2, val: string(ch)})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{EQUALS, string(ch)})
	case '\n':
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: NEWLINE2, val: "\\n"})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{NEWLINE, "\\n"})
	case '\t':
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: TAB2, val: string(ch)})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{TAB, string(ch)})
	default:
		LexicalAnalyisResult2 = append(LexicalAnalyisResult2, &item{item: ILLEGAL2, val: string(ch)})
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{ILLEGAL, string(ch)})
	}

}
