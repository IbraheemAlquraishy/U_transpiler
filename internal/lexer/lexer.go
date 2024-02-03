package lexer

import (
	token "github.com/IbraheemAlquraishy/U_transpiler/internal/modules/token"
)

type Lexer struct {
	Input string
	Pos   int
	Rpos  int
	Ch    byte
}

// the main func in the lexer object
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipwhitespace()
	l.skipcomments()
	switch l.Ch {
	case '=':
		if l.peekChar() == '=' {
			Ch := l.Ch
			l.readChar()
			lit := string(Ch) + string(l.Ch)
			tok = token.Token{Type: token.Isequal, Lit: lit}
		} else {
			tok = newtoken(token.Assign, l.Ch)
		}
	case ';':
		tok = newtoken(token.SEMICOLON, l.Ch)
	case '(':
		tok = newtoken(token.LPAREN, l.Ch)
	case ')':
		tok = newtoken(token.RPAREN, l.Ch)
	case ',':
		tok = newtoken(token.COMMA, l.Ch)
	case '+':
		if l.peekChar() == '=' {
			Ch := l.Ch
			l.readChar()
			lit := string(Ch) + string(l.Ch)
			tok = token.Token{Type: token.Plusequal, Lit: lit}
		} else if l.peekChar() == '+' {
			Ch := l.Ch
			l.readChar()
			lit := string(Ch) + string(l.Ch)
			tok = token.Token{Type: token.Inc, Lit: lit}
		} else {
			tok = newtoken(token.Plus, l.Ch)
		}
	case '-':
		if l.peekChar() == '=' {
			Ch := l.Ch
			l.readChar()
			lit := string(Ch) + string(l.Ch)
			tok = token.Token{Type: token.Subequal, Lit: lit}
		} else if l.peekChar() == '-' {
			Ch := l.Ch
			l.readChar()
			lit := string(Ch) + string(l.Ch)
			tok = token.Token{Type: token.Dec, Lit: lit}
		} else {
			tok = newtoken(token.Sub, l.Ch)
		}
	case '*':
		if l.peekChar() == '=' {
			Ch := l.Ch
			l.readChar()
			lit := string(Ch) + string(l.Ch)
			tok = token.Token{Type: token.Multiequal, Lit: lit}
		} else {
			tok = newtoken(token.Multi, l.Ch)
		}
	case '/':
		if l.peekChar() == '=' {
			Ch := l.Ch
			l.readChar()
			lit := string(Ch) + string(l.Ch)
			tok = token.Token{Type: token.Divequal, Lit: lit}
		} else {
			tok = newtoken(token.Div, l.Ch)
		}
	case '{':
		tok = newtoken(token.LBRACE, l.Ch)
	case '}':
		tok = newtoken(token.RBRACE, l.Ch)
	case '"':
		tok.Type = token.Str
		tok.Lit = l.readstring()
	case '>':
		if l.peekChar() == '=' {
			Ch := l.Ch
			l.readChar()
			lit := string(Ch) + string(l.Ch)
			tok = token.Token{Type: token.Greaterorequal, Lit: lit}
		} else {
			tok = newtoken(token.Greater, l.Ch)
		}
	case '<':
		if l.peekChar() == '=' {
			Ch := l.Ch
			l.readChar()
			lit := string(Ch) + string(l.Ch)
			tok = token.Token{Type: token.Lowerorequal, Lit: lit}
		} else {
			tok = newtoken(token.Lower, l.Ch)
		}
	case '!':
		if l.peekChar() == '=' {
			Ch := l.Ch
			l.readChar()
			lit := string(Ch) + string(l.Ch)
			tok = token.Token{Type: token.Notequal, Lit: lit}
		} else {
			tok = newtoken(token.Not, l.Ch)
		}
	case '^':
		if l.peekChar() == '=' {
			Ch := l.Ch
			l.readChar()
			lit := string(Ch) + string(l.Ch)
			tok = token.Token{Type: token.Powerequal, Lit: lit}
		} else {
			tok = newtoken(token.Power, l.Ch)
		}
	case ':':
		if l.peekChar() == '=' {
			Ch := l.Ch
			l.readChar()
			lit := string(Ch) + string(l.Ch)
			tok = token.Token{Type: token.COLONEqual, Lit: lit}
		} else {
			tok = newtoken(token.Illegal, l.Ch)
		}
	case 0:
		tok.Lit = ""
		tok.Type = token.EOF
	default:
		if isletter(l.Ch) {
			tok.Lit = l.readident()
			tok.Type = token.Lookupident(tok.Lit)
			return tok
		} else if isDigit(l.Ch) {
			if l.isfloat() {
				tok.Type = token.Float
				tok.Lit = l.readfloat()
			} else {
				tok.Type = token.Int
				tok.Lit = l.readnumber()
			}
			return tok
		} else {
			tok = newtoken(token.Illegal, l.Ch)
		}
	}
	l.readChar()
	return tok
}

// constractor
func New(Input string) *Lexer {
	l := &Lexer{Input: Input}
	l.readChar()
	return l
}

// move to the next Char
func (l *Lexer) readChar() {
	if l.Rpos >= len(l.Input) {
		l.Ch = 0
	} else {
		l.Ch = l.Input[l.Rpos]
	}
	l.Pos = l.Rpos
	l.Rpos += 1
}

func (l *Lexer) peekChar() byte {
	if l.Rpos >= len(l.Input) {
		return 0
	} else {
		return l.Input[l.Rpos]
	}
}

// read funcs
func (l *Lexer) readnumber() string {
	Pos := l.Pos
	for isDigit(l.Ch) {
		l.readChar()
	}
	return l.Input[Pos:l.Pos]
}

func (l *Lexer) readfloat() string {
	Pos := l.Pos

	for isDigit(l.Ch) || l.Ch == '.' {
		l.readChar()
	}

	return l.Input[Pos:l.Pos]
}

func (l *Lexer) readstring() string {
	l.readChar()
	Pos := l.Pos
	for l.Ch != '"' {
		l.readChar()
	}
	return l.Input[Pos:l.Pos]
}

func (l *Lexer) readident() string {
	Pos := l.Pos
	for isletter(l.Ch) {
		l.readChar()
	}
	return l.Input[Pos:l.Pos]
}

// is funcs
func isletter(Ch byte) bool {
	return 'a' <= Ch && Ch <= 'z' || 'A' <= Ch && Ch <= 'Z' || Ch == '_'
}

func (l *Lexer) isfloat() bool {
	Pos := l.Pos
	for isDigit(l.Input[Pos]) {
		if Pos < len(l.Input)-1 {
			Pos += 1
		} else {
			break
		}
	}
	if l.Input[Pos] == '.' {
		return true
	}
	return false
}

func isDigit(Ch byte) bool {
	return '0' <= Ch && Ch <= '9'
}

// others
func (l *Lexer) skipwhitespace() {
	for l.Ch == ' ' || l.Ch == '\t' || l.Ch == '\n' || l.Ch == '\r' {
		l.readChar()
	}
}

func newtoken(tokentype token.Tokentype, Ch byte) token.Token {
	return token.Token{Type: tokentype, Lit: string(Ch)}
}

func (l *Lexer) skipcomments() {
	if l.Ch == '/' {

		if l.peekChar() == '/' {
			for l.Ch != '\n' {
				l.readChar()
			}
			l.skipwhitespace()
		}
	}
}
