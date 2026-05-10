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

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '(':
		tok = newToken(token.LPARA, l.ch)
	case ')':
		tok = newToken(token.RPARA, l.ch)
	case '+':
		tok = newToken(token.ADD, l.ch)
	case '{':
		tok = newToken(token.LBRAC, l.ch)
	case '}':
		tok = newToken(token.RBRAC, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	}

	//Update the token line before returning
	l.ReadChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
