package parser

import (
	"fmt"
	"testing"

	"github.com/IbraheemAlquraishy/U_transpiler/internal/lexer"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/ast"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/token"
)

func TestDeclareStatments(t *testing.T) {
	input := `
	for i int=0;i<5;i++{

	}
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("parseprgram() returned nil")
	}
	// if len(program.Statements) != 2 {
	// 	t.Fatalf("program.statments does not contain 2 statment.got=%d", len(program.Statements))
	// }
	for _, k := range program.Statements {
		t.Log(k)
	}

}

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
	input := `return 5<6;`
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
		t.Log(returnstmt.String())
	}
}

func TestFunctionstatments(t *testing.T) {
	input := `func jasim()`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	for _, stmt := range program.Statements {
		funcstmt, ok := stmt.(*ast.Functionstatment)
		if !ok {
			t.Errorf("stmt not *ast.functionstatment got=%T", stmt)
			continue
		}
		if funcstmt.Tokenliteral() != "func" {
			t.Errorf("functionstmt.tokenliteral not 'func', got %q", funcstmt.Tokenliteral())
		}
		t.Log(funcstmt.String())
	}
}

func TestIdentityExpression(t *testing.T) {
	input := "foo;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.statment[0] is not ast.expressoinstatment got=%T", program.Statements[0])
	}
	ident, ok := stmt.Expression.(*ast.Identity)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foo" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar", ident.Tokenliteral())
	}
}

func TestIntLitExpression(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}
	lit, ok := stmt.Expression.(*ast.IntegerLit)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	if lit.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, lit.Value)
	}
	if lit.Tokenliteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5",
			lit.Tokenliteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}
	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
		fmt.Printf(stmt.String())
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int) bool {
	integ, ok := il.(*ast.IntegerLit)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}
	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}
	if integ.Tokenliteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value,
			integ.Tokenliteral())
		return false
	}
	return true

}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int
		operator   string
		rightValue int
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}
	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("exp is not ast.InfixExpression. got=%T", stmt.Expression)
		}
		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			return
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// {
		// 	"-a * b",
		// 	"((-a) * b)"},
		// {
		// 	"!-a",
		// 	"(!(-a))",
		// },
		// {
		// 	"a + b + c",
		// 	"((a + b) + c)",
		// },
		// {
		// 	"a + b - c",
		// 	"((a + b) - c)",
		// },
		// {
		// 	"a * b * c",
		// 	"((a * b) * c)",
		// },
		// {
		// 	"a * b / c",
		// 	"((a * b) / c)",
		// },
		// {
		// 	"a + b / c",
		// 	"(a + (b / c))",
		// },
		// {
		// 	"a + b * c + d / e - f",
		// 	"(((a + (b * c)) + (d / e)) - f)",
		// },
		// {
		// 	"3 + 4; -5 * 5",
		// 	"(3 + 4)((-5) * 5)",
		// },
		// {
		// 	"5 > 4 == 3 < 4",
		// 	"((5 > 4) == (3 < 4))",
		// },
		// {
		// 	"5 < 4 != 3 > 4",
		// 	"((5 < 4) != (3 > 4))",
		// },
		// {
		// 	"3 + 4 * 5 == 3 * 1 + 4 * 5",
		// 	"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		// },
		{
			"x>=(y==false)",
			"(x >= (y == false))",
		},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}
