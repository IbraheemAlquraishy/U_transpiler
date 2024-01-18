package src

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var program = []Inctruct{}
var Token []tokens

func Readfile(filename string) {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	content := string(file)
	parsefiletocodes(content)
	parsecodes()
}

func parsefiletocodes(content string) {
	var ins Inctruct
	ins.code = ""
	for _, data := range content {
		ins.code += string(data)
		if string(data) == ";" || string(data) == "}" {
			program = append(program, ins)

			ins.code = ""
		}
	}
}

func parsecodes() {

	for _, ins := range program {
		var t tokens

		newins2 := strings.ReplaceAll(ins.code, `\n`, "")
		newins3 := strings.ReplaceAll(newins2, `\t`, "")

		newins4 := strings.ReplaceAll(newins3, "endl", "<<endl")

		n := strings.Fields(newins4)
		for _, s := range n {

			t.push(s)

		}
		Token = append(Token, t)
	}
	fmt.Println(Token)
}
