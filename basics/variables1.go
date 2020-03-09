package main

import "fmt"

func main() {
    // variables in a function can use implicit assignment :=
    i := 10 // no type needed
    k := 20
    c, python, java := true, false, "no!"

	fmt.Println(i, c, python, java, k)
}
