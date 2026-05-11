package lexer

import "go-interpreter/token"

type Lexer struct {
	input        string
	position     int  // curr pos in input
	readPosition int  // curr reading pos in input
	ch           byte //current character examining
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.ReadChar()
	return l
}

// Method to read the Character in the code :
func (l *Lexer) ReadChar() {
	// Check if end of char , if yes change the char to 0:
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		//else current char = Character in reading buffer.
		l.ch = l.input[l.readPosition]
	}
	//Move the pointer buffers :

	//Last read position
	l.position = l.readPosition

	//Next reading position
	l.readPosition += 1
}

// Method to send next token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			//Temp to store curr char
			ch := l.ch
			//GO to next char
			l.ReadChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '(':
		tok = newToken(token.LPARA, l.ch)
	case ')':
		tok = newToken(token.RPARA, l.ch)
	case '+':
		tok = newToken(token.ADD, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '{':
		tok = newToken(token.LBRAC, l.ch)
	case '}':
		tok = newToken(token.RBRAC, l.ch)
	case '/':
		tok = newToken(token.FORSLASH, l.ch)
	case '-':
		tok = newToken(token.SUBTRACT, l.ch)
	case '*':
		tok = newToken(token.STAR, l.ch)
	case '<':
		tok = newToken(token.LSIGN, l.ch)
	case '>':
		tok = newToken(token.RSIGN, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.ReadChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.EXCLA, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok

		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}
	//Update the token line before returning
	l.ReadChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.ReadChar()
	}

	//if not a letter slice from postion -> current readingPos
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.ReadChar()
	}

	return l.input[position:l.position]
}

// Check if its a letter or _
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.ReadChar()
	}
}

func (l *Lexer) peekChar() byte {

	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
