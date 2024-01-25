package lexer

import (
	"fmt"
	"testing"

	token "github.com/IbraheemAlquraishy/U_transpiler/internal/modules"
)

func TestNextToken(t *testing.T) {
	input := `if 5 ==5{
		jasim int=0;
		
		jasim ++;
		jasim += 5;
		jasim --;
		jasim>=5;
	}
	
	`
	tests := []struct {
		expectedtype token.Tokentype
		expectedlit  string
	}{
		{token.If, "if"},
		{token.Int, "5"},
		{token.Isequal, "=="},
		{token.Int, "5"},
		{token.LBRACE, "{"},
		{token.Ident, "jasim"},
		{token.Intt, "int"},
		{token.Assign, "="},
		{token.Int, "0"},
		{token.SEMICOLON, ";"},
		{token.Ident, "jasim"},
		{token.Inc, "++"},
		{token.SEMICOLON, ";"},
		{token.Ident, "jasim"},
		{token.Plusequal, "+="},
		{token.Int, "5"},
		{token.SEMICOLON, ";"},
		{token.Ident, "jasim"},
		{token.Dec, "--"},
		{token.SEMICOLON, ";"},
		{token.Ident, "jasim"},
		{token.Greaterorequal, ">="},
		{token.Int, "5"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		fmt.Println(tok)
		if tok.Type != tt.expectedtype {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedtype, tok.Type)
		}
		if tok.Lit != tt.expectedlit {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedlit, tok.Lit)
		}
	}
}
