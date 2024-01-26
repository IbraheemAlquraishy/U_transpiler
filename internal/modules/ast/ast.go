package ast

import "github.com/IbraheemAlquraishy/U_transpiler/internal/modules/token"

type Node interface {
	Tokenliteral() string
}

type StatmentType string

const (
	DeclareStatment = "declare"
	AssignStatment  = "assign"
)

type Statement interface {
	Node
	statementnode()

	Tokentype() token.Tokentype
}

type Expression interface {
	Node
	expressionnode()
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

type Declarestatment struct {
	Stattype StatmentType
	Name     *Identity

	Value Expression
}

func (ds *Declarestatment) statementnode() {}

func (ds *Declarestatment) Tokenliteral() string {
	return ds.Name.Tokenliteral()
}
func (ds *Declarestatment) Tokentype() token.Tokentype {
	return ds.Name.Tokentype()
}

func (ds *Declarestatment) New(i *Identity) {
	ds.Stattype = DeclareStatment
	ds.Name = i
}

type Retrunstatment struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *Retrunstatment) statementnode()             {}
func (rs *Retrunstatment) Tokenliteral() string       { return rs.Token.Lit }
func (rs *Retrunstatment) Tokentype() token.Tokentype { return rs.Token.Type }

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
