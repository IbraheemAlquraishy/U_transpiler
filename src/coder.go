package src

import (
	"log"
	"os"
)

func Coding() {
	file, err := os.Create("./test/code.cpp")
	if err != nil {
		log.Fatal("file exist")
	}
	defer file.Close()
	line := `#include <iostream>

	using namespace std;
	
	int main(){
		`
	for _, x := range Token {

		for _, s := range x.stack {

			switch s {
			case "int":
				line += "int "
			case "string":
				line += "string "
			case "float":
				line += "float "
			case "bool":
				line += "bool "
			case `if`:
				line += "if"
			case "for":
				line += "for"
			case "print":
				line += "cout<<"
			case "input":
				line += "cin>>"

			default:

				line += s + " "

			}

		}

		line += `
`

	}
	line += `
}`
	file.WriteString(line)
}
