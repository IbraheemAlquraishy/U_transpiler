package lexer

import (
	"fmt"
	"testing"

	token "github.com/IbraheemAlquraishy/U_transpiler/internal/modules/token"
)

func TestNextToken(t *testing.T) {
	input := `
		//another comment
		for(){
		jasim :=0;
		//comment
		jasim ++;
		jasim += 5;
		jasim --;
		jasim>=5;
		print jasim;
	}
	
	`
	tests := []struct {
		expectedtype token.Tokentype
		expectedlit  string
	}{
		{token.For, "for"},

		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.Ident, "jasim"},
		{token.COLONEqual, ":="},
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
		{token.Print, "print"},
		{token.Ident, "jasim"},
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
