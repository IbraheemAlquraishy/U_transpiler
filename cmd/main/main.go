package main

import (
	"log"
	"os"

	"github.com/IbraheemAlquraishy/U_transpiler/internal/lexer"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/parser"
	"github.com/IbraheemAlquraishy/U_transpiler/internal/typechecker"
)

func main() {
	file, err := os.ReadFile("./tests/test1.u")
	if err != nil {
		log.Fatal("cant open the file")
	}
	l := lexer.New(string(file))
	p := parser.New(l)
	var c typechecker.Checker
	c = typechecker.New(p)
	c.Info()
	//f, e := os.Create("./tests/test1")
	// if e != nil {
	// 	log.Fatal(e)
	// }
	// le, er := f.WriteString(program.String())
	// if er != nil {
	// 	log.Fatal(er)
	// }
	// fmt.Print(le)
	// defer f.Close()
}
