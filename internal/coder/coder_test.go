package coder

import (
	"log"
	"os"
	"testing"
)

func TestCode(t *testing.T) {
	file, err := os.ReadFile("../../tests/test1.u")
	if err != nil {
		log.Fatal("cant open the file")
	}
	Code(string(file))

}
