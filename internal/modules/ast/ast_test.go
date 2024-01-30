package ast

import (
	"testing"

	"github.com/IbraheemAlquraishy/U_transpiler/internal/modules/token"
)

func TestString(t *testing.T) {
	program := &Program{Statements: []Statement{&Declarestatment{Name: &Identity{
		Token: token.Token{Type: token.Ident, Lit: "x"},
		Type:  token.Intt,
		Value: "x",
	}, Value: nil,
	},
	},
	}
	if program.String() != "x int=;" {
		t.Errorf("program.string is wrong got=%q", program.String())
	}
}
