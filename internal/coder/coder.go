package coder

import (
	"fmt"

	"github.com/IbraheemAlquraishy/U_transpiler/internal/lexer"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/ast"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/token"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/parser"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/typechecker"
)

func Code(file string) string {
	l := lexer.New(string(file))
	p := parser.New(l)
	program := p.ParseProgram()
	cf := typechecker.New(program, p.Errors())
	cf.Checkfuncs()
	cr := typechecker.Newf(program, p.Errors(), cf.Funcs)
	cr.CheckRest()
	c := `#include <iostream>
	#include <math.h>
		using namespace std;


	`
	c += writefuncs(cf.ReturnCheckedprogram())
	c += `
	
	int main(){

	`
	bs := &ast.BlockStatement{Statements: cr.Program.Statements}
	c += writestatment(bs)

	c += "}"
	return c
}

func writefuncs(program *ast.Program) string {
	s := ""
	for _, i := range program.Statements {
		j, _ := i.(*ast.Functionstatment)
		s += string(j.Name.Type) + " " + j.Name.Value + "("
		for n, p := range j.Param {
			s += string(p.Type) + " " + p.Value
			if n != len(j.Param)-1 {
				s += ","
			}
		}
		s += "){"
		s += writestatment(j.Body)
		s += "}"
	}
	return s
}

func writestatment(s ast.Statement) string {
	st := ""
	if i, ok := s.(*ast.Declarestatment); ok {
		st += string(i.Name.Type) + " " + i.Name.Value + " "
		if i.Value != nil {
			st += "="
			st += writeexpression(i.Value)
		}
		st += ";"
	} else if i, ok := s.(*ast.Retrunstatment); ok {
		st += "return "
		st += writeexpression(i.ReturnValue)
		st += ";"
	} else if i, ok := s.(*ast.BlockStatement); ok {
		for _, j := range i.Statements {
			st += writestatment(j)
		}
	} else if _, ok := s.(*ast.Functionstatment); ok {
		//not ganna be used
	} else if i, ok := s.(*ast.ExpressionStatement); ok {
		st += writeexpression(i.Expression)
		st += ";"
	} else {
		i, _ := s.(*ast.PrintStatment)
		if i.Tokentype() == token.Println {
			st += "cout<<"
			st += writeexpression(i.Data)
			st += "<<endl"

		} else if i.Tokentype() == token.Print {
			st += "cout<<"
			st += writeexpression(i.Data)
		} else {
			st += "cin>>"
			st += writeexpression(i.Data)
		}

		st += ";"
	}
	return st
}

func writeexpression(e ast.Expression) string {
	s := ""
	if i, ok := e.(*ast.PrefixExpression); ok {
		s += i.Operator
		s += writeexpression(i.Right)
	} else if i, ok := e.(*ast.InfixExpression); ok {
		if i.Operator == "^" {
			s += "("
			s += "pow("
			s += writeexpression(i.Left)
			s += ","
			s += writeexpression(i.Right)
			s += "))"
		} else if i.Operator == "^=" {
			s += "("
			s += writeexpression(i.Left)
			s += "="
			s += "pow("
			s += writeexpression(i.Left)
			s += ","
			s += writeexpression(i.Right)
			s += "))"
		} else {
			s += "("
			s += writeexpression(i.Left)
			s += i.Operator
			s += writeexpression(i.Right)
			s += ")"
		}
	} else if i, ok := e.(*ast.IncExpression); ok {
		s += "("
		s += writeexpression(i.Left)
		s += i.Tokenliteral()
		s += ")"
	} else if i, ok := e.(*ast.IfExpression); ok {
		s += "if"
		s += writeexpression(i.Condition)
		s += "{"
		s += writestatment(i.Consequence)
		s += "}"
		if i.Alternative != nil {
			s += "else{"
			s += writestatment(i.Alternative)
			s += "}"
		}
	} else if i, ok := e.(*ast.ForExpression); ok {
		s += "for ("
		if i.Var != nil {
			s += writestatment(i.Var)
		}
		s += writeexpression(i.Condition)
		s += ";"
		s += writeexpression(i.Relation)
		s += ")"
		s += "{"
		s += writestatment(i.Body)
		s += "}"
	} else if i, ok := e.(*ast.CallExpression); ok {
		s += i.Function.Name.Value + "("
		for n, p := range i.Arg {
			s += writeexpression(p)
			if n != len(i.Arg)-1 {
				s += ","
			}

		}
		s += ")"
	} else if i, ok := e.(*ast.Identity); ok {
		s += i.Value
	} else if i, ok := e.(*ast.IntegerLit); ok {
		s += fmt.Sprintf("%v", i.Value)
	} else if i, ok := e.(*ast.Boolean); ok {
		if i.Value {
			s += "true"
		} else {
			s += "false"
		}
	} else if i, ok := e.(*ast.Stringlit); ok {
		s += `"` + i.Value + `"`
	} else if i, ok := e.(*ast.Floatlit); ok {
		s += fmt.Sprintf("%f", i.Value)
	}

	return s
}
