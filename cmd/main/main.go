package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IbraheemAlquraishy/U_transpiler/internal/coder"
)

func main() {
	t := time.Now()
	file, err := os.ReadFile("./tests/test1.u")
	if err != nil {
		log.Fatal(err)
	}
	code := coder.Code(string(file))
	f, e := os.Create("./tests/test1.cpp")
	if e != nil {
		log.Fatal(e)
	}
	le, er := f.WriteString(code)
	if er != nil {
		log.Fatal(er)
	}
	fmt.Print(le)
	defer func() {
		fmt.Println(time.Since(t))
		f.Close()
	}()
}
