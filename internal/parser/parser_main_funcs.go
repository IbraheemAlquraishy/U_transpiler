package parser

import (
	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/ast"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/token"
)

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curtoken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curtoken.Type {
	case token.Ident:
		if !p.PeekTokenIs(token.Intt) && !p.PeekTokenIs(token.COLONEqual) && !p.PeekTokenIs(token.Strt) && !p.PeekTokenIs(token.Floatt) && !p.PeekTokenIs(token.Boolt) {
			return p.parseExpressionstatment()
		} else {
			return p.parsedeclarestatment()
		}
	case token.Return:
		return p.parseReturnstatment()
	case token.FUNCTION:
		return p.parseFunctionstatment()
	case token.Print:
		return p.parsePrintStatment()
	case token.Println:
		return p.parsePrintStatment()
	case token.Input:
		return p.parsePrintStatment()
	default:
		return p.parseExpressionstatment()
	}
}

func (p *Parser) parsedeclarestatment() *ast.Declarestatment {
	var it *ast.Identity

	if p.PeekTokenIs(token.COLONEqual) {
		it = &ast.Identity{Token: p.curtoken, Value: p.curtoken.Lit}
		p.nextToken()
		it.Type = token.UNdefined

	} else {
		it = &ast.Identity{Token: p.curtoken, Value: p.curtoken.Lit, Type: p.peektoken.Type}
		p.nextToken()
		p.nextToken()
	}

	var stmt ast.Declarestatment
	stmt.New(it)

	for !p.CurtokenIs(token.SEMICOLON) {

		if p.CurtokenIs(token.Assign) || p.CurtokenIs(token.COLONEqual) {
			p.nextToken()
			if p.CurtokenIs(token.Ident) && p.PeekTokenIs(token.LPAREN) {
				t := &ast.Identity{Token: p.curtoken, Value: p.curtoken.Lit, Type: token.UNdefined}
				f, err := p.lookupfuncs(*t)
				if err != "" {
					p.Peekerror("no such function")
					return nil
				}
				if p.checkifInfixnotcall() {
					stmt.Value = p.parseInfixcall(t, f)
				} else {
					stmt.Value = p.parseCallexpression(t, f)
				}
			} else {
				stmt.Value = p.parseExpression(LOWEST)
			}
		} else if p.CurtokenIs(token.If) || p.CurtokenIs(token.For) || p.CurtokenIs(token.FUNCTION) || p.CurtokenIs(token.Print) || p.CurtokenIs(token.Println) || p.CurtokenIs(token.Input) {
			p.Peekerror("expected ;")
		}
		p.nextToken()
	}

	return &stmt
}

func (p *Parser) parseReturnstatment() *ast.Retrunstatment {
	stmt := &ast.Retrunstatment{Token: p.curtoken}
	p.nextToken()
	stmt.ReturnValue = p.parseExpression(LOWEST)

	p.nextToken()

	return stmt
}

func (p *Parser) parseFunctionstatment() *ast.Functionstatment {
	stmt := &ast.Functionstatment{Token: p.curtoken}
	if !p.expectedpeek(token.Ident) {
		p.Peekerror("no name for the function")
		return nil
	}
	t := &ast.Identity{Token: p.curtoken, Value: p.curtoken.Lit}
	if !p.expectedpeek(token.LPAREN) {
		p.Peekerror("no () for the function")
		return nil
	}
	stmt.Param = p.parseFunctionParams()

	if !p.PeekTokenIs(token.Intt) && !p.PeekTokenIs(token.COLONEqual) && !p.PeekTokenIs(token.Strt) && !p.PeekTokenIs(token.Floatt) && !p.PeekTokenIs(token.Boolt) && !p.PeekTokenIs(token.LBRACE) {
		p.Peekerror("not a valid data type for return function")
		return nil
	} else if p.PeekTokenIs(token.LBRACE) {
		p.nextToken()
		t.Type = token.Voidt
		stmt.Body = p.parseBlockStatment()

	} else {
		p.nextToken()

		t.Type = p.curtoken.Type
		p.nextToken()
		stmt.Body = p.parseBlockStatment()

	}

	stmt.Name = t
	p.savefunc(*t, *stmt)
	return stmt
}

func (p *Parser) parseFunctionParams() []*ast.Identity {
	ids := []*ast.Identity{}
	if p.PeekTokenIs(token.RPAREN) {
		p.nextToken()
		return ids
	}
	p.nextToken()
	ident := &ast.Identity{Token: p.curtoken, Value: p.curtoken.Lit}

	if !p.PeekTokenIs(token.Intt) && !p.PeekTokenIs(token.COLONEqual) && !p.PeekTokenIs(token.Strt) && !p.PeekTokenIs(token.Floatt) && !p.PeekTokenIs(token.Boolt) {
		p.Peekerror("not a valid data type")
		return nil
	}
	p.nextToken()
	ident.Type = p.curtoken.Type
	ids = append(ids, ident)
	for p.PeekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identity{Token: p.curtoken, Value: p.curtoken.Lit}
		p.nextToken()
		ident.Type = p.curtoken.Type
		ids = append(ids, ident)
	}
	if !p.expectedpeek(token.RPAREN) {
		p.Peekerror("not ) for the func")
		return nil
	}
	return ids
}

func (p *Parser) parseExpressionstatment() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curtoken}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.PeekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixparsefns[p.curtoken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curtoken.Type)
		return nil
	}
	leftexp := prefix()
	for !p.PeekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixparsefns[p.peektoken.Type]
		if infix == nil {
			return leftexp
		}
		p.nextToken()
		leftexp = infix(leftexp)
	}
	return leftexp
}

func (p *Parser) parseprefixexpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curtoken,
		Operator: p.curtoken.Lit,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curtoken,
		Operator: p.curtoken.Lit,
		Left:     left,
	}
	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}

func (p *Parser) parseIncExpression(left ast.Expression) ast.Expression {

	expression := &ast.IncExpression{Token: p.curtoken, Left: left}
	return expression

}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()
	exp := p.parseExpression(LOWEST)

	if !p.expectedpeek(token.RPAREN) {
		return nil
	}
	return exp
}

func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.curtoken}

	if !p.expectedpeek(token.LPAREN) {
		return nil
	}

	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)
	if !p.expectedpeek(token.RPAREN) {
		return nil
	}
	if !p.expectedpeek(token.LBRACE) {
		return nil
	}
	expression.Consequence = p.parseBlockStatment()
	if p.PeekTokenIs(token.Else) {
		p.nextToken()
		if !p.expectedpeek(token.LBRACE) {
			return nil
		}
		expression.Alternative = p.parseBlockStatment()
	}
	return expression
}

func (p *Parser) parseCallexpression(i *ast.Identity, f ast.Functionstatment) ast.Expression {

	exp := &ast.CallExpression{Token: p.curtoken, Function: f}
	exp.Arg = p.parsecallargs()
	return exp
}

func (p *Parser) parsecallargs() []ast.Expression {
	args := []ast.Expression{}
	p.nextToken()
	if p.PeekTokenIs(token.RPAREN) {
		p.nextToken()
		return args
	}

	p.nextToken()

	args = append(args, p.parseExpression(LOWEST))

	for p.PeekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}
	if !p.expectedpeek(token.RPAREN) {
		return nil
	}
	return args

}

func (p *Parser) parseForExpression() ast.Expression {
	f := &ast.ForExpression{Token: p.curtoken}
	p.nextToken()
	if p.CurtokenIs(token.Ident) && (p.PeekTokenIs(token.Intt) || p.PeekTokenIs(token.COLONEqual) || p.PeekTokenIs(token.Strt) || p.PeekTokenIs(token.Floatt) || p.PeekTokenIs(token.Boolt)) {
		f.Var = p.parsedeclarestatment()
	}
	p.nextToken()
	f.Condition = p.parseExpression(LOWEST)
	p.nextToken()
	p.nextToken()
	f.Relation = p.parseExpression(LOWEST)
	if !p.PeekTokenIs(token.LBRACE) {
		p.Peekerror("no body for the for")
		return nil
	}
	p.nextToken()
	f.Body = p.parseBlockStatment()
	return f
}

func (p *Parser) parsePrintStatment() *ast.PrintStatment {
	ps := &ast.PrintStatment{Token: p.curtoken}
	p.nextToken()
	ps.Data = p.parseExpression(LOWEST)
	p.nextToken()
	return ps
}

func (p *Parser) parseInfixcall(t *ast.Identity, f ast.Functionstatment) ast.Expression {
	ex := &ast.InfixExpression{Left: p.parseCallexpression(t, f), Operator: p.peektoken.Lit}
	p.nextToken()
	precedence := p.curPrecedence()
	p.nextToken()
	ex.Right = p.parseExpression(precedence)
	return ex
}
