package main

import "fmt"

func main() {
    i := 42
    f := 3.141
    g := 0.123 + 0.456i

    fmt.Printf("Type of i is %T\n", i)
    fmt.Printf("Type of f is %T\n", f)
    fmt.Printf("Type of g is %T\n", g)
}
