package scanner

import (
	"fmt"
	"hoax/token"
)

type Scanner struct {
	Source  string
	Start   int
	Pointer int
	Line    int
	Tokens  []token.Token
}

func (s *Scanner) ScanTokens() {
	fmt.Println("Scanning tokens")
	for !s.isAtEnd() {
		s.Start = s.Pointer
		s.ScanToken()
	}
	s.Tokens = append(s.Tokens, token.NewToken(token.EOF, "", "", s.Line))
}

func (s *Scanner) ScanToken() {
	fmt.Println("Scanning...", "at index", s.Pointer)
	phoneme := s.Advance()
	fmt.Println("Detected phoneme", phoneme)
	switch phoneme {
	// 1 characters Tokens
	case "(":
		s.Tokens = append(s.Tokens, token.NewToken(token.LEFT_PAREN, "(", "", s.Line))
	case ")":
		s.Tokens = append(s.Tokens, token.NewToken(token.RIGHT_PAREN, ")", "", s.Line))
	case "{":
		s.Tokens = append(s.Tokens, token.NewToken(token.LEFT_BRACE, "{", "", s.Line))
	case "}":
		s.Tokens = append(s.Tokens, token.NewToken(token.RIGHT_BRACE, "}", "", s.Line))
	case ",":
		s.Tokens = append(s.Tokens, token.NewToken(token.COMMA, ",", "", s.Line))
	case ".":
		s.Tokens = append(s.Tokens, token.NewToken(token.DOT, ".", "", s.Line))
	case "-":
		s.Tokens = append(s.Tokens, token.NewToken(token.MINUS, "-", "", s.Line))
	case "+":
		s.Tokens = append(s.Tokens, token.NewToken(token.PLUS, "+", "", s.Line))
	case ";":
		s.Tokens = append(s.Tokens, token.NewToken(token.SEMICOLON, ";", "", s.Line))
	case "*":
		s.Tokens = append(s.Tokens, token.NewToken(token.STAR, "*", "", s.Line))

	// 2 characters Tokens
	case "!":
		peek := s.Peek()
		fmt.Println("Peeked", peek, s.Pointer)
		switch peek {
		case "=":
			s.Tokens = append(s.Tokens, token.NewToken(token.BANG_EQUAL, "!=", "", s.Line))
			s.Advance()
		default:
			s.Tokens = append(s.Tokens, token.NewToken(token.BANG, "!", "", s.Line))
		}
	case "=":
		peek := s.Peek()
		switch peek {
		case "=":
			s.Tokens = append(s.Tokens, token.NewToken(token.EQUAL_EQUAL, "==", "", s.Line))
			s.Advance()
		default:
			s.Tokens = append(s.Tokens, token.NewToken(token.EQUAL, "=", "", s.Line))
		}
	case ">":
		peek := s.Peek()
		switch peek {
		case "=":
			s.Tokens = append(s.Tokens, token.NewToken(token.GREATER_EQUAL, ">=", "", s.Line))
			s.Advance()
		default:
			s.Tokens = append(s.Tokens, token.NewToken(token.GREATER, ">", "", s.Line))
		}
	case "<":
		peek := s.Peek()
		switch peek {
		case "=":
			s.Tokens = append(s.Tokens, token.NewToken(token.LESS_EQUAL, "<=", "", s.Line))
			s.Advance()
		default:
			s.Tokens = append(s.Tokens, token.NewToken(token.LESS, "<", "", s.Line))
		}
	// literals
	// string
	case "\"", "'":
		isAtQuote := false
		for !isAtQuote {
			curr := s.Advance()
			// matches quote phoneme
			if curr == phoneme {
				stringLiteral := s.Source[s.Start:s.Pointer]
				s.Tokens = append(s.Tokens, token.NewToken(token.STRING, stringLiteral, "", s.Line))
				isAtQuote = true
			} else if curr == "\000" {
				// error
			} else {
				curr = s.Advance()
			}
		}
	case " ", "\r", "\t":
		fmt.Println("Unrecognized character", phoneme)
	default: // whitespace or unrecognized characters
		if isAlpha(phoneme) {
			// Identify known keywords
			for isAlphaNumeric(s.Peek()) {
				s.Advance()
			}
			text := s.Source[s.Start:s.Pointer]
			tokType, keyword := checkKeyword(text)
			if keyword {
				s.Tokens = append(s.Tokens, token.NewToken(tokType, text, "", s.Line))
			} else {
				s.Tokens = append(s.Tokens, token.NewToken(token.IDENTIFIER, text, "", s.Line))
			}

		} else if isDigit(phoneme) {
			//integer or float
			isFloat := false

			for isDigit(s.Peek()) {
				s.Advance()
			}

			// look for decimal
			if s.Peek() == "." {
				isFloat = true
				for isDigit(s.Peek()) {
					s.Advance()
				}
			}

			number := s.Source[s.Start:s.Pointer]
			if isFloat {
				s.Tokens = append(s.Tokens, token.NewToken(token.FLOAT, number, "", s.Line))
			} else {
				s.Tokens = append(s.Tokens, token.NewToken(token.NUMBER, number, "", s.Line))
			}

		} else {
			fmt.Println("Unrecognized character", phoneme)
		}
	}
}

func isAlpha(c string) bool {
	return (c >= "a" && c <= "z") || (c >= "A" && c <= "Z") || c == "_"
}

func isDigit(c string) bool {
	return c >= "0" && c <= "9"
}

func isAlphaNumeric(c string) bool {
	return isAlpha(c) || isDigit(c)
}

func checkKeyword(text string) (token.TokenType, bool) {
	switch text {
	case "and":
		return token.AND, true
	case "class":
		return token.CLASS, true
	case "else":
		return token.ELSE, true
	case "false":
		return token.FALSE, true
	case "fun":
		return token.FUN, true
	case "for":
		return token.FOR, true
	case "if":
		return token.IF, true
	case "nil":
		return token.NIL, true
	case "or":
		return token.OR, true
	case "print":
		return token.PRINT, true
	case "return":
		return token.RETURN, true
	case "super":
		return token.SUPER, true
	case "this":
		return token.THIS, true
	case "true":
		return token.TRUE, true
	case "var":
		return token.VAR, true
	case "while":
		return token.WHILE, true
	default:
		return token.IDENTIFIER, false // assuming 0 is the zero value for token.TokenType; adjust as necessary
	}
}

// loop utilities
func (s *Scanner) Advance() string {
	s.Pointer = s.Pointer + 1
	literal := string(s.Source[s.Pointer-1]) // -1 because we call advance() at start of scanToken()
	return literal
}

func (s *Scanner) Peek() string {
	if s.isAtEnd() {
		return "\000"
	}
	return string(s.Source[s.Pointer]) // not -1 = next because we call advance() at start of scanToken()
}

func (s *Scanner) Literal() string {
	if s.isAtEnd() {
		return "\000"
	}
	return s.Source[s.Pointer:s.Pointer]
}

func (s *Scanner) isAtEnd() bool {
	return s.Pointer >= len(s.Source)
}
