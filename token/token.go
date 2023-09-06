package token

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}

func NewToken(tokenType TokenType, lexeme string, literal any, line int) Token {
	return Token{tokenType, lexeme, literal, line}
}

func (t Token) String() string {
	return t.Type.String() + " " + t.Lexeme + " " + t.Literal.(string)
}
