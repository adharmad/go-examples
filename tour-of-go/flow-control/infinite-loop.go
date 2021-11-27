package main

import "fmt"

func main() {
    // make an infinite loop using for
    // by omitting the condition
    iter := 0

    for {
        fmt.Println("Iteration number", iter)
        iter++
    }
}
