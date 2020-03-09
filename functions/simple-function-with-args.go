package main

import "fmt"

func main() {
    sum := add(10, 20)
    fmt.Printf("sum = %d\n", sum)
}

func add(i int, j int) int {
    return i+j
}
