package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	query, _ := ioutil.ReadFile("main.go")
	r := strings.NewReader(strings.ToLower(string(query)))

	scanner := NewScanner(r)

	lex := scanner.Scan()

	println(lex)

}
