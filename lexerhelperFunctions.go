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
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{FUNCTION, buffer})
	case "as":
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{AS, buffer})
	case "end":
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{END, buffer})
	case "sub":
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{SUB, buffer})
	case "for":
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{FOR, buffer})
	case "if":
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{IF, buffer})
	case "while":
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{WHILE, buffer})
	case "return":
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{RETURN, buffer})
	case "this":
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{THIS, buffer})
	case "object", "int", "string", "void":
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{PARAMS, buffer})
	default:
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{IDENT, buffer})
	}

}

func isSpecialChar(ch rune) {
	//Process special Chars
	switch ch {
	case '*':
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{ASTERISK, string(ch)})
	case ',':
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{COMMA, string(ch)})
	case ':':
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{COLON, string(ch)})
	case ';':
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{SEMICOLON, string(ch)})
	case '(':
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{OPENBRACKET, string(ch)})
	case ')':
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{CLOSEBRACKET, string(ch)})
	case '{':
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{OPENCURLYBRACKET, string(ch)})
	case '}':
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{CLOSECURLYBRACKET, string(ch)})
	case '.':
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{DOT, string(ch)})
	case '_':
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{UNDERSCORE, string(ch)})
	case '=':
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{EQUALS, string(ch)})
	case '\n':
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{NEWLINE, "\\n"})
	case '\t':
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{TAB, string(ch)})
	default:
		LexicalAnalyisResult = append(LexicalAnalyisResult, []string{ILLEGAL, string(ch)})
	}

}
