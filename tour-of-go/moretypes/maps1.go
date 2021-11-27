package main

import "fmt"

var m map[string]string

func main() {
	m = make(map[string]string)
	m["test1"] = "Alan Turing"
	m["test2"] = "Kurt Godel"
	m["test3"] = "Evariste Galois"

	fmt.Println(m)
}
