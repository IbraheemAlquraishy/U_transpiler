package ast

import (
	"bytes"

	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/token"
)

type Node interface {
	Tokenliteral() string
	String() string
}

type Statement interface {
	Node
	statementnode()
	Tokenliteral() string
	Tokentype() token.Tokentype
}

type Expression interface {
	Node
	expressionnode()
	Tokentype() token.Tokentype
}

type Program struct {
	Statements []Statement
}

func (p *Program) Tokenliteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].Tokenliteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type Identity struct {
	Token token.Token
	Type  token.Tokentype
	Value string
}

func (i *Identity) expressionnode() {}

func (i *Identity) Tokenliteral() string {
	return i.Token.Lit
}

func (i *Identity) Tokentype() token.Tokentype {
	return i.Type
}
func (i *Identity) String() string {
	switch i.Type {
	case token.Intt:
		return i.Value + " int"
	case token.Floatt:
		return i.Value + " float"
	case token.Boolt:
		return i.Value + " bool"
	case token.Strt:
		return i.Value + " string"
	default:
		return i.Value
	}
}

type IntegerLit struct {
	Token token.Token
	Value int
}

func (il *IntegerLit) expressionnode()            {}
func (il *IntegerLit) Tokenliteral() string       { return il.Token.Lit }
func (il *IntegerLit) Tokentype() token.Tokentype { return il.Token.Type }
func (il *IntegerLit) String() string {
	return il.Token.Lit
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionnode()            {}
func (b *Boolean) Tokenliteral() string       { return b.Token.Lit }
func (b *Boolean) Tokentype() token.Tokentype { return b.Token.Type }
func (b *Boolean) String() string             { return b.Token.Lit }

//TODO stringlit struct and its parser

type Stringlit struct {
	Token token.Token
	Value string
}

func (s *Stringlit) expressionnode()            {}
func (s *Stringlit) Tokenliteral() string       { return s.Token.Lit }
func (s *Stringlit) Tokentype() token.Tokentype { return s.Token.Type }
func (s *Stringlit) String() string {
	return s.Value
}

type Floatlit struct {
	Token token.Token
	Value float64
}

func (f *Floatlit) expressionnode()            {}
func (f *Floatlit) Tokenliteral() string       { return f.Token.Lit }
func (f *Floatlit) Tokentype() token.Tokentype { return f.Token.Type }
func (f *Floatlit) String() string {
	return f.Token.Lit
}
