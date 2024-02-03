package typechecker

import (
	"log"
	"os"
	"testing"

	"github.com/IbraheemAlquraishy/U_transpiler/internal/lexer"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/parser"
)

func TestCheckernew(t *testing.T) {
	file, err := os.ReadFile("../../tests/test1.u")
	if err != nil {
		log.Fatal(err)
	}
	l := lexer.New(string(file))
	p := parser.New(l)
	program := p.ParseProgram()
	c := New(program, p.Errors())
	c.Checkall()

}
