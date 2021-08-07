package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	character    byte
}

func New(input string) *Lexer {
	l := Lexer{input: input}
	l.readCharacter()
	return &l
}

func (l *Lexer) readCharacter() {
	if l.readPosition >= len(l.input) {
		l.character = 0
	} else {
		l.character = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.character {
	case '=':
		if l.peekCharacter() == '=' {
			ch := l.character
			l.readCharacter()
			literal := string(ch) + string(l.character)
			//TODO flo: refactor to use newToken method, or overload the method if needed for byte
			tok = token.Token{
				Type:    token.EQ,
				Literal: literal,
			}
		} else {
			tok = newToken(token.ASSIGN, l.character)
		}
	case '!':
		if l.peekCharacter() == '=' {
			ch := l.character
			l.readCharacter()
			literal := string(ch) + string(l.character)
			//TODO flo: refactor to use newToken method, or overload the method if needed for byte
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: literal,
			}
		} else {
			tok = newToken(token.BANG, l.character)
		}
	case '+':
		tok = newToken(token.PLUS, l.character)
	case '-':
		tok = newToken(token.MINUS, l.character)
	case '/':
		tok = newToken(token.SLASH, l.character)
	case '*':
		tok = newToken(token.ASTERISK, l.character)
	case '<':
		tok = newToken(token.LT, l.character)
	case '>':
		tok = newToken(token.GT, l.character)
	case '(':
		tok = newToken(token.LPAREN, l.character)
	case ')':
		tok = newToken(token.RPAREN, l.character)
	case '{':
		tok = newToken(token.LBRACE, l.character)
	case '}':
		tok = newToken(token.RBRACE, l.character)
	case ',':
		tok = newToken(token.COMMA, l.character)
	case ';':
		tok = newToken(token.SEMICOLON, l.character)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.character) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.character) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.character)
		}
	}

	l.readCharacter()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.character == ' ' || l.character == '\t' || l.character == '\n' || l.character == '\r' {
		l.readCharacter()
	}
}

//TODO flo: think about correct place for this function, wouldn't it make sense to place it in the 'token' package
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (l *Lexer) peekCharacter() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.character) {
		l.readCharacter()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.character) {
		l.readCharacter()
	}
	return l.input[position:l.position]
}
