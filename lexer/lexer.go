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

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.character {
	case '=':
		if l.peekCharacter() == '=' {
			ch := l.character
			l.readCharacter()
			literal := string(ch) + string(l.character)
			tok = token.New(token.EQ, literal)
		} else {
			tok = token.New(token.ASSIGN, string(l.character))
		}
	case '!':
		if l.peekCharacter() == '=' {
			ch := l.character
			l.readCharacter()
			literal := string(ch) + string(l.character)
			tok = token.New(token.NOT_EQ, literal)
		} else {
			tok = token.New(token.BANG, string(l.character))
		}
	case '+':
		tok = token.New(token.PLUS, string(l.character))
	case '-':
		tok = token.New(token.MINUS, string(l.character))
	case '/':
		tok = token.New(token.SLASH, string(l.character))
	case '*':
		tok = token.New(token.ASTERISK, string(l.character))
	case '<':
		tok = token.New(token.LT, string(l.character))
	case '>':
		tok = token.New(token.GT, string(l.character))
	case '(':
		tok = token.New(token.LPAREN, string(l.character))
	case ')':
		tok = token.New(token.RPAREN, string(l.character))
	case '{':
		tok = token.New(token.LBRACE, string(l.character))
	case '}':
		tok = token.New(token.RBRACE, string(l.character))
	case '[':
		tok = token.New(token.LBRACKET, string(l.character))
	case ']':
		tok = token.New(token.RBRACKET, string(l.character))
	case ',':
		tok = token.New(token.COMMA, string(l.character))
	case ';':
		tok = token.New(token.SEMICOLON, string(l.character))
	case ':':
		tok = token.New(token.COLON, string(l.character))
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
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
			tok = token.New(token.ILLEGAL, string(l.character))
		}
	}

	l.readCharacter()
	return tok
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

func (l *Lexer) skipWhitespace() {
	for l.character == ' ' || l.character == '\t' || l.character == '\n' || l.character == '\r' {
		l.readCharacter()
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

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readCharacter()
		if l.character == '"' || l.character == 0 {
			break
		}
	}

	return l.input[position:l.position]
}
