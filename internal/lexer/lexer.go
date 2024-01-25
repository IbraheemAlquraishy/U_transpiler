package lexer

import (
	token "github.com/IbraheemAlquraishy/U_transpiler/internal/modules"
)

type Lexer struct {
	input string
	pos   int
	rpos  int
	ch    byte
}

// the main func in the lexer object
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipwhitespace()

	switch l.ch {
	case '=':
		if l.peekchar() == '=' {
			ch := l.ch
			l.readchar()
			lit := string(ch) + string(l.ch)
			tok = token.Token{Type: token.Isequal, Lit: lit}
		} else {
			tok = newtoken(token.Assign, l.ch)
		}
	case ';':
		tok = newtoken(token.SEMICOLON, l.ch)
	case '(':
		tok = newtoken(token.LPAREN, l.ch)
	case ')':
		tok = newtoken(token.RPAREN, l.ch)
	case ',':
		tok = newtoken(token.COMMA, l.ch)
	case '+':
		if l.peekchar() == '=' {
			ch := l.ch
			l.readchar()
			lit := string(ch) + string(l.ch)
			tok = token.Token{Type: token.Plusequal, Lit: lit}
		} else if l.peekchar() == '+' {
			ch := l.ch
			l.readchar()
			lit := string(ch) + string(l.ch)
			tok = token.Token{Type: token.Inc, Lit: lit}
		} else {
			tok = newtoken(token.Plus, l.ch)
		}
	case '-':
		if l.peekchar() == '=' {
			ch := l.ch
			l.readchar()
			lit := string(ch) + string(l.ch)
			tok = token.Token{Type: token.Subequal, Lit: lit}
		} else if l.peekchar() == '-' {
			ch := l.ch
			l.readchar()
			lit := string(ch) + string(l.ch)
			tok = token.Token{Type: token.Dec, Lit: lit}
		} else {
			tok = newtoken(token.Sub, l.ch)
		}
	case '*':
		if l.peekchar() == '=' {
			ch := l.ch
			l.readchar()
			lit := string(ch) + string(l.ch)
			tok = token.Token{Type: token.Multiequal, Lit: lit}
		} else {
			tok = newtoken(token.Multi, l.ch)
		}
	case '/':
		if l.peekchar() == '=' {
			ch := l.ch
			l.readchar()
			lit := string(ch) + string(l.ch)
			tok = token.Token{Type: token.Divequal, Lit: lit}
		} else {
			tok = newtoken(token.Div, l.ch)
		}
	case '{':
		tok = newtoken(token.LBRACE, l.ch)
	case '}':
		tok = newtoken(token.RBRACE, l.ch)
	case '"':
		tok.Type = token.Str
		tok.Lit = l.readstring()
	case '>':
		if l.peekchar() == '=' {
			ch := l.ch
			l.readchar()
			lit := string(ch) + string(l.ch)
			tok = token.Token{Type: token.Greaterorequal, Lit: lit}
		} else {
			tok = newtoken(token.Greater, l.ch)
		}
	case '<':
		if l.peekchar() == '=' {
			ch := l.ch
			l.readchar()
			lit := string(ch) + string(l.ch)
			tok = token.Token{Type: token.Lowerorequal, Lit: lit}
		} else {
			tok = newtoken(token.Lower, l.ch)
		}
	case '!':
		if l.peekchar() == '=' {
			ch := l.ch
			l.readchar()
			lit := string(ch) + string(l.ch)
			tok = token.Token{Type: token.Notequal, Lit: lit}
		} else {
			tok = newtoken(token.Not, l.ch)
		}
	case '^':
		if l.peekchar() == '=' {
			ch := l.ch
			l.readchar()
			lit := string(ch) + string(l.ch)
			tok = token.Token{Type: token.Powerequal, Lit: lit}
		} else {
			tok = newtoken(token.Power, l.ch)
		}
	case 0:
		tok.Lit = ""
		tok.Type = token.EOF
	default:
		if isletter(l.ch) {
			tok.Lit = l.readident()
			tok.Type = token.Lookupident(tok.Lit)
			return tok
		} else if isDigit(l.ch) {
			if l.isfloat() {
				tok.Type = token.Float
				tok.Lit = l.readfloat()
			} else {
				tok.Type = token.Int
				tok.Lit = l.readnumber()
			}
			return tok
		} else {
			tok = newtoken(token.Illegal, l.ch)
		}
	}
	l.readchar()
	return tok
}

// constractor
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readchar()
	return l
}

// move to the next char
func (l *Lexer) readchar() {
	if l.rpos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.rpos]
	}
	l.pos = l.rpos
	l.rpos += 1
}

func (l *Lexer) peekchar() byte {
	if l.rpos >= len(l.input) {
		return 0
	} else {
		return l.input[l.rpos]
	}
}

// read funcs
func (l *Lexer) readnumber() string {
	pos := l.pos
	for isDigit(l.ch) {
		l.readchar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readfloat() string {
	pos := l.pos

	for isDigit(l.ch) || l.ch == '.' {
		l.readchar()
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) readstring() string {
	l.readchar()
	pos := l.pos
	for l.ch != '"' {
		l.readchar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readident() string {
	pos := l.pos
	for isletter(l.ch) {
		l.readchar()
	}
	return l.input[pos:l.pos]
}

// is funcs
func isletter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) isfloat() bool {
	pos := l.pos
	for isDigit(l.input[pos]) {
		pos += 1
	}
	if l.input[pos] == '.' {
		return true
	}
	return false
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// others
func (l *Lexer) skipwhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readchar()
	}
}

func newtoken(tokentype token.Tokentype, ch byte) token.Token {
	return token.Token{Type: tokentype, Lit: string(ch)}
}
