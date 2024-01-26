package parser

import (
	"testing"

	"github.com/IbraheemAlquraishy/U_transpiler/internal/lexer"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/ast"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/token"
)

// func TestDeclareStatments(t *testing.T) {
// 	input := `x int=5;
// 	x :="jasim";`
// 	l := lexer.New(input)
// 	p := New(l)

// 	program := p.ParseProgram()
// 	checkParserErrors(t, p)
// 	if program == nil {
// 		t.Fatalf("parseprgram() returned nil")
// 	}
// 	// if len(program.Statements) != 2 {
// 	// 	t.Fatalf("program.statments does not contain 2 statment.got=%d", len(program.Statements))
// 	// }
// 	tests := []struct {
// 		expectedIdentifier string
// 		expectedtype       token.Tokentype
// 	}{
// 		{"x", token.Intt},
// 		{"x", token.Strt},
// 	}
// 	for i, tt := range tests {
// 		stmt := program.Statements[i]
// 		if !testDeclareStatment(t, stmt, tt.expectedIdentifier, tt.expectedtype) {
// 			return
// 		}
// 	}

// }

func testDeclareStatment(t *testing.T, s ast.Statement, name string, ty token.Tokentype) bool {
	if s.Tokentype() != ty {
		t.Errorf("s.tokentype not '%q' got=%q", ty, s.Tokentype())
	}
	letstmt, ok := s.(*ast.Declarestatment)
	if !ok {
		t.Errorf("s not *ast.declarestatment. got=%T", s)
		return false
	}
	if letstmt.Name.Tokenliteral() != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letstmt.Name.Tokenliteral())
		return false
	}
	if letstmt.Name.Tokenliteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s", name, letstmt.Name.Tokenliteral())
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.errors
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestReturnstatment(t *testing.T) {
	input := `return 5;`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	for _, stmt := range program.Statements {
		returnstmt, ok := stmt.(*ast.Retrunstatment)
		if !ok {
			t.Errorf("stmt not * ast.returnstatment.got=%T", stmt)
			continue
		}
		if returnstmt.Tokenliteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				returnstmt.Tokenliteral())
		}

	}
}
