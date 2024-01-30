package ast

import (
	"bytes"

	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/token"
)

type Declarestatment struct {
	Name *Identity

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

	ds.Name = i
}

func (ds *Declarestatment) String() string {
	var out bytes.Buffer
	out.WriteString(ds.Name.String())
	out.WriteString("=")
	if ds.Value != nil {
		out.WriteString(ds.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

type Retrunstatment struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *Retrunstatment) statementnode()             {}
func (rs *Retrunstatment) Tokenliteral() string       { return rs.Token.Lit }
func (rs *Retrunstatment) Tokentype() token.Tokentype { return rs.Token.Type }
func (rs *Retrunstatment) String() string {
	var out bytes.Buffer
	out.WriteString(rs.Tokenliteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	return out.String()
}

type Functionstatment struct {
	Token token.Token
	Name  *Identity
	Param []*Identity
	Body  *BlockStatement
}

func (f *Functionstatment) statementnode() {}
func (f *Functionstatment) Tokenliteral() string {
	return f.Token.Lit
}

func (f *Functionstatment) Tokentype() token.Tokentype {
	return f.Token.Type
}

func (f *Functionstatment) String() string {
	var out bytes.Buffer
	out.WriteString(f.Tokenliteral() + " ")
	out.WriteString(f.Name.Token.Lit)
	out.WriteString("(")
	for i, p := range f.Param {
		out.WriteString(p.String())
		if i < len(f.Param)-1 {
			out.WriteString(",")
		}
	}
	out.WriteString(")")
	out.WriteString(string(f.Name.Tokentype()))
	out.WriteString("{")
	if f.Body != nil {
		out.WriteString(f.Body.String())
	}
	out.WriteString("}")
	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementnode()             {}
func (es *ExpressionStatement) Tokenliteral() string       { return es.Token.Lit }
func (es *ExpressionStatement) Tokentype() token.Tokentype { return es.Token.Type }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementnode()             {}
func (bs *BlockStatement) Tokenliteral() string       { return bs.Token.Lit }
func (bs *BlockStatement) Tokentype() token.Tokentype { return bs.Token.Type }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
		out.WriteString(";")
	}
	return out.String()
}

type PrintStatment struct {
	Token token.Token
	Data  Expression
}

func (ps *PrintStatment) statementnode()             {}
func (ps *PrintStatment) Tokenliteral() string       { return ps.Token.Lit }
func (ps *PrintStatment) Tokentype() token.Tokentype { return ps.Token.Type }
func (ps *PrintStatment) String() string {
	var out bytes.Buffer
	out.WriteString(ps.Token.Lit)
	out.WriteString(" ")
	out.WriteString(ps.Data.String())
	return out.String()
}
