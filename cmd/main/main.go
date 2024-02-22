package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/IbraheemAlquraishy/U_transpiler/internal/coder"
)

func main() {
	t := time.Now()
	leng := len(os.Args)

	if leng < 2 {
		log.Fatal("no file to transpiler")
	}
	p := os.Args[1]

	file, err := os.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}
	code := coder.Code(string(file))
	o := strings.Split(p, ".u")[0]
	o += ".cpp"
	if leng > 2 {
		if os.Args[2] == "-o" {
			if leng != 4 {
				log.Fatal("no output file")
			}
			o = os.Args[3]
		}
	}
	f, e := os.Create(o)
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
