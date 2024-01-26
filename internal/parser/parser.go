package parser

import (
	"fmt"

	"github.com/IbraheemAlquraishy/U_transpiler/internal/lexer"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/ast"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/token"
)

type Parser struct {
	l *lexer.Lexer

	errors []string

	curtoken  token.Token
	peektoken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekerror(err string) {
	msg := fmt.Sprintf(err)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curtoken = p.peektoken
	p.peektoken = p.l.NextToken()
}

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

		return p.parsedeclarestatment()
	case token.Return:
		return p.parseReturnstatment()
	default:
		return nil
	}
}

func (p *Parser) parsedeclarestatment() *ast.Declarestatment {
	var it *ast.Identity
	if !p.peekTokenIs(token.Intt) && !p.peekTokenIs(token.COLONEqual) && !p.peekTokenIs(token.Strt) && !p.peekTokenIs(token.Floatt) && !p.peekTokenIs(token.Boolt) {
		p.peekerror("declarestatement syntax error")
		return nil
	} else if p.peekTokenIs(token.COLONEqual) {
		it = &ast.Identity{Token: p.curtoken, Value: p.peektoken.Lit}
		p.nextToken()
		if p.peekTokenIs(token.Str) {
			it.Type = token.Strt
		} else if p.peekTokenIs(token.Float) {
			it.Type = token.Floatt
		} else if p.peekTokenIs(token.Bool) {
			it.Type = token.Boolt
		} else {
			it.Type = token.Intt
		}

	} else {
		it = &ast.Identity{Token: p.curtoken, Value: p.curtoken.Lit, Type: p.peektoken.Type}
	}
	var stmt ast.Declarestatment
	stmt.New(it)

	for !p.curtokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return &stmt
}

func (p *Parser) curtokenIs(t token.Tokentype) bool {
	return p.curtoken.Type == t
}

func (p *Parser) peekTokenIs(t token.Tokentype) bool {
	return p.peektoken.Type == t
}

// no use for this func yet TODO delete it when the project is finished
func (p *Parser) expectedpeek(t token.Tokentype) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {

		return false
	}
}

func (p *Parser) parseReturnstatment() *ast.Retrunstatment {
	stmt := &ast.Retrunstatment{Token: p.curtoken}
	p.nextToken()
	for !p.curtokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
