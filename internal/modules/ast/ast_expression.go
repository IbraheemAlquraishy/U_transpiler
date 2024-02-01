package ast

import (
	"bytes"
	"strings"

	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/token"
)

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionnode()      {}
func (pe *PrefixExpression) Tokenliteral() string { return pe.Token.Lit }
func (pe *PrefixExpression) Tokentype() token.Tokentype {
	return pe.Token.Type
}
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionnode() {}
func (ie *InfixExpression) Tokenliteral() string {
	return ie.Token.Lit
}
func (ie *InfixExpression) Tokentype() token.Tokentype {
	return ie.Token.Type
}
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}

type IncExpression struct {
	Token token.Token
	Left  Expression
}

func (in *IncExpression) expressionnode() {}
func (in *IncExpression) Tokenliteral() string {
	return in.Token.Lit
}
func (in *IncExpression) Tokentype() token.Tokentype {
	return in.Token.Type
}
func (in *IncExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(" + in.Left.String())
	out.WriteString(string(in.Token.Type) + ")")

	return out.String()
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionnode()            {}
func (ie *IfExpression) Tokenliteral() string       { return ie.Token.Lit }
func (ie *IfExpression) Tokentype() token.Tokentype { return ie.Token.Type }
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" {")
	out.WriteString(ie.Consequence.String())
	out.WriteString(" }")
	if ie.Alternative != nil {
		out.WriteString("else {")
		out.WriteString(ie.Alternative.String())
		out.WriteString(" }")
	}
	return out.String()
}

type ForExpression struct {
	Token     token.Token
	Var       *Declarestatment
	Condition Expression
	Relation  Expression
	Body      *BlockStatement
}

func (fe *ForExpression) expressionnode()            {}
func (fe *ForExpression) Tokenliteral() string       { return fe.Token.Lit }
func (fe *ForExpression) Tokentype() token.Tokentype { return fe.Token.Type }
func (fe *ForExpression) String() string {
	var out bytes.Buffer
	out.WriteString("for ")
	out.WriteString(fe.Var.String())
	out.WriteString(fe.Condition.String())
	out.WriteString(fe.Relation.String())
	out.WriteString(" {")
	out.WriteString(fe.Body.String())
	out.WriteString(" }")
	return out.String()
}

type CallExpression struct {
	Token    token.Token
	Function Functionstatment
	Arg      []Expression
}

func (ce *CallExpression) expressionnode()            {}
func (ce *CallExpression) Tokenliteral() string       { return ce.Token.Lit }
func (ce *CallExpression) Tokentype() token.Tokentype { return ce.Token.Type }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range ce.Arg {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.Name.Value)
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}
