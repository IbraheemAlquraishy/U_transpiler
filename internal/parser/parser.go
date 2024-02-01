package parser

import (
	"fmt"

	"strconv"

	"github.com/IbraheemAlquraishy/U_transpiler/internal/lexer"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/ast"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/token"
)

type (
	prefixParsefunc func() ast.Expression
	infixParsefunc  func(ast.Expression) ast.Expression
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	INFIX
	PREFIX
	CALL
)

var precedences = map[token.Tokentype]int{
	token.Assign:         EQUALS,
	token.Isequal:        EQUALS,
	token.Notequal:       EQUALS,
	token.Lower:          LESSGREATER,
	token.Greater:        LESSGREATER,
	token.Lowerorequal:   LESSGREATER,
	token.Greaterorequal: LESSGREATER,
	token.Plus:           SUM,
	token.Sub:            SUM,
	token.Multi:          PRODUCT,
	token.Div:            PRODUCT,
	token.Power:          PRODUCT,
	token.Plusequal:      SUM,
	token.Subequal:       SUM,
	token.Inc:            INFIX,
	token.Dec:            INFIX,
	token.Multiequal:     PRODUCT,
	token.Divequal:       PRODUCT,
	token.Powerequal:     PRODUCT,
}

type Parser struct {
	l *lexer.Lexer

	errors []string

	curtoken  token.Token
	peektoken token.Token

	prefixparsefns map[token.Tokentype]prefixParsefunc
	infixparsefns  map[token.Tokentype]infixParsefunc
	funcs          map[string]ast.Functionstatment
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()
	p.prefixparsefns = make(map[token.Tokentype]prefixParsefunc)
	p.registerPrefix(token.Ident, p.parseIdentity)
	p.registerPrefix(token.Int, p.parseIntlit)
	p.registerPrefix(token.Str, p.parseStringlit)
	p.registerPrefix(token.Float, p.parseFloatlit)
	p.registerPrefix(token.Sub, p.parseprefixexpression)
	p.registerPrefix(token.Not, p.parseprefixexpression)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.False, p.parseBoolean)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.If, p.parseIfExpression)
	p.registerPrefix(token.For, p.parseForExpression)
	p.infixparsefns = make(map[token.Tokentype]infixParsefunc)
	p.registerInfix(token.Plus, p.parseInfixExpression)
	p.registerInfix(token.Sub, p.parseInfixExpression)
	p.registerInfix(token.Multi, p.parseInfixExpression)
	p.registerInfix(token.Div, p.parseInfixExpression)
	p.registerInfix(token.Power, p.parseInfixExpression)
	p.registerInfix(token.Plusequal, p.parseInfixExpression)
	p.registerInfix(token.Subequal, p.parseInfixExpression)
	p.registerInfix(token.Multiequal, p.parseInfixExpression)
	p.registerInfix(token.Divequal, p.parseInfixExpression)
	p.registerInfix(token.Powerequal, p.parseInfixExpression)
	p.registerInfix(token.Greater, p.parseInfixExpression)
	p.registerInfix(token.Lower, p.parseInfixExpression)
	p.registerInfix(token.Greaterorequal, p.parseInfixExpression)
	p.registerInfix(token.Lowerorequal, p.parseInfixExpression)
	p.registerInfix(token.Assign, p.parseInfixExpression)
	p.registerInfix(token.Notequal, p.parseInfixExpression)
	p.registerInfix(token.Isequal, p.parseInfixExpression)
	p.registerInfix(token.Inc, p.parseIncExpression)
	p.registerInfix(token.Dec, p.parseIncExpression)
	p.funcs = make(map[string]ast.Functionstatment)
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) Peekerror(err string) {
	msg := fmt.Sprintf(err)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curtoken = p.peektoken
	p.peektoken = p.l.NextToken()
}

func (p *Parser) registerPrefix(tokentype token.Tokentype, fn prefixParsefunc) {
	p.prefixparsefns[tokentype] = fn
}

func (p *Parser) registerInfix(tokentype token.Tokentype, fn infixParsefunc) {
	p.infixparsefns[tokentype] = fn
}

func (p *Parser) CurtokenIs(t token.Tokentype) bool {
	return p.curtoken.Type == t
}

func (p *Parser) PeekTokenIs(t token.Tokentype) bool {
	return p.peektoken.Type == t
}

func (p *Parser) expectedpeek(t token.Tokentype) bool {
	if p.PeekTokenIs(t) {
		p.nextToken()
		return true
	} else {

		return false
	}
}

func (p *Parser) parseIdentity() ast.Expression {
	return &ast.Identity{Token: p.curtoken, Value: p.curtoken.Lit, Type: token.UNdefined}
}

func (p *Parser) parseIntlit() ast.Expression {
	lit := &ast.IntegerLit{Token: p.curtoken}
	value, err := strconv.ParseInt(p.curtoken.Lit, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("cout not parse %q as int", p.curtoken.Lit)
		p.Peekerror(msg)
		return nil
	}
	lit.Value = int(value)
	return lit
}

func (p *Parser) parseStringlit() ast.Expression {
	lit := &ast.Stringlit{Token: p.curtoken, Value: p.curtoken.Lit}
	return lit
}

func (p *Parser) parseFloatlit() ast.Expression {
	lit := &ast.Floatlit{Token: p.curtoken}
	v, err := strconv.ParseFloat(p.curtoken.Lit, 64)
	if err != nil {
		p.Peekerror("cant convert to float 64")
		return nil
	}
	lit.Value = v
	return lit
}

func (p *Parser) noPrefixParseFnError(t token.Tokentype) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peektoken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curtoken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curtoken, Value: p.CurtokenIs(token.TRUE)}
}

func (p *Parser) parseBlockStatment() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curtoken}
	block.Statements = []ast.Statement{}
	p.nextToken()
	for !p.CurtokenIs(token.RBRACE) && !p.CurtokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}

func (p *Parser) lookupfuncs(i ast.Identity) (ast.Functionstatment, string) {
	f, ok := p.funcs[i.Value]
	if ok {
		return f, ""
	}
	return f, "no such func"

}

func (p *Parser) savefunc(i ast.Identity, fn ast.Functionstatment) {
	p.funcs[i.Value] = fn
}
