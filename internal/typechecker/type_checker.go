package typechecker

import (
	"errors"
	"fmt"
	"log"

	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/ast"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/funcmap"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/token"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/parser"
)

type Checker struct {
	Program *ast.Program
	Errs    []error
	Ids     map[string]ast.Identity
	Funcs   funcmap.Fns
}

func New(p *parser.Parser) Checker {
	var c Checker
	c.Program = p.ParseProgram()
	e := p.Errors()
	if len(e) != 0 {
		log.Fatalln(e)

	}
	c.Ids = make(map[string]ast.Identity)
	c.Funcs.New()

	return c
}

func (c *Checker) Returnprogram() *ast.Program {
	return c.Program
}

func (c *Checker) Printprogram() {
	for _, i := range c.Program.Statements {
		fmt.Println(i.String())
	}

}

func (c *Checker) Info() {
	fmt.Println(len(c.Program.Statements))
	for _, i := range c.Program.Statements {
		c.EvalStatment(i)
	}
}

func (c *Checker) EvalExpression(e ast.Expression) string {
	var t string

	if _, ok := e.(*ast.PrefixExpression); ok {
		t = "prefix"
	} else if _, ok := e.(*ast.InfixExpression); ok {
		t = "infix"

	} else if _, ok := e.(*ast.IncExpression); ok {
		t = "inc"
	} else if _, ok := e.(*ast.IfExpression); ok {
		t = "if"
	} else if _, ok := e.(*ast.ForExpression); ok {
		t = "for"
	} else if i, ok := e.(*ast.CallExpression); ok {
		t = "call"
		t = string(i.Function.Name.Type)
	} else if _, ok := e.(*ast.Identity); ok {
		t = "ident"
	} else if _, ok := e.(*ast.IntegerLit); ok {
		t = "int"
	} else if _, ok := e.(*ast.Boolean); ok {
		t = "bool"
	} else if _, ok := e.(*ast.Stringlit); ok {
		t = "string"
	} else if _, ok := e.(*ast.Floatlit); ok {
		t = "float"
	} else {

		e := errors.New("not a valid expression")
		c.Errs = append(c.Errs, e)
	}
	if t == "" {
		c.ThrowErrs()
	}
	return t

}

func (c *Checker) ThrowErrs() {
	log.Fatal(c.Errs)
}

func (c *Checker) EvalStatment(s ast.Statement) {
	var t string
	i := s
	if i, ok := s.(*ast.Declarestatment); ok {
		t = "declare"
		tr := c.EvalExpression(i.Value)
		var typ token.Tokentype
		switch tr {
		case "int":
			typ = token.Intt

		case "string":
			typ = token.Strt

		case "bool":
			typ = token.Boolt

		case "float":
			typ = token.Floatt

		}
		if i.Name.Type == token.UNdefined {
			i.Name.Type = typ
			c.registerId(*i.Name)
		} else if i.Name.Type == typ {
			c.registerId(*i.Name)
		} else {
			e := errors.New("this function returns diffrent datatype")
			c.Errs = append(c.Errs, e)
			c.ThrowErrs()
		}

	} else if i, ok := s.(*ast.Retrunstatment); ok {
		t = "return"
		c.EvalExpression(i.ReturnValue)
	} else if i, ok := s.(*ast.BlockStatement); ok {
		t = "block"
		for _, j := range i.Statements {
			c.EvalStatment(j)
		}
	} else if i, ok := s.(*ast.Functionstatment); ok {
		t = "function"
		if c.Funcs.Exists(i.Name.Value, i.Param) {
			e := errors.New("this function exist with the same params")
			c.Errs = append(c.Errs, e)
			c.ThrowErrs()
		}
		c.EvalStatment(i.Body)
		c.Funcs.Add(i.Name.Value, i.Name.Type, i)
	} else if i, ok := s.(*ast.ExpressionStatement); ok {
		t = "expression"
		if j, ok := i.Expression.(*ast.InfixExpression); ok {
			c.evalInfix(j)
		}

	} else if i, ok := s.(*ast.PrintStatment); ok {
		t = "print"
		c.EvalExpression(i.Data)
	} else {
		e := errors.New("not a valid statment")
		c.Errs = append(c.Errs, e)
	}
	if t == "" {
		c.ThrowErrs()
	} else {
		fmt.Print(t + "\t" + i.String())
	}
}

func (c *Checker) registerId(i ast.Identity) {
	r := c.lookupid(i.Value, i.Type)
	if r == 2 {
		c.Ids[i.Value] = i
	} else if r == 1 {
		e := errors.New("ident " + i.Value + " is defiend with a defrend data type")
		c.Errs = append(c.Errs, e)
		c.ThrowErrs()
	} else if r == 3 {
		c.Ids[i.Value] = i
	} else {
		e := errors.New("ident " + i.Value + " is defiend ")
		c.Errs = append(c.Errs, e)
		c.ThrowErrs()
	}
}

func (c *Checker) lookupid(v string, t token.Tokentype) int {
	i, ok := c.Ids[v]

	if ok {
		if i.Type == token.UNdefined {
			return 3
		} else if i.Type == t {
			return 0
		}
		return 1
	}
	return 2
}

func (c *Checker) getidtype(v string) token.Tokentype {
	i, _ := c.Ids[v]
	return i.Type
}

func (c *Checker) getid(v string) ast.Identity {

	i, ok := c.Ids[v]
	if !ok {

	}
	return i
}

func (c *Checker) evalInfix(in *ast.InfixExpression) string {
	tl := c.EvalExpression(in.Left)
	if tl == "infix" {
		j, _ := in.Left.(*ast.InfixExpression)
		tl = c.evalInfix(j)
	}
	tr := c.EvalExpression(in.Right)
	if tr == "infix" {
		j, _ := in.Right.(*ast.InfixExpression)
		tr = c.evalInfix(j)
	}

	if tl == "ident" {
		j, _ := in.Left.(*ast.Identity)
		r := c.lookupid(j.Value, j.Type)
		var typ token.Tokentype
		switch tr {
		case "int":
			typ = token.Intt

		case "string":
			typ = token.Strt

		case "bool":
			typ = token.Boolt

		case "float":
			typ = token.Floatt

		}
		if r == 3 {
			j.Type = typ
			c.registerId(*j)
		} else if r == 1 {
			ty := c.getidtype(j.Value)
			if ty != typ {
				err := errors.New("not a valid assign expression ")
				c.Errs = append(c.Errs, err)
				c.ThrowErrs()
			}
		}
	} else if tr != tl {
		err := errors.New("not a valid infix expression")
		c.Errs = append(c.Errs, err)
		c.ThrowErrs()
	} else if (in.Operator != "+" && in.Operator != "==") && (tl == "string" || tr == "string") {
		err := errors.New("not a valid string expression")
		c.Errs = append(c.Errs, err)
		c.ThrowErrs()
	}
	return tr
}
