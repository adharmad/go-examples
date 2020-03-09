package main

import "fmt"

func main() {
    nested_func := func() {
        fmt.Println("Hello from nested function")
    }

    nested_func()
}
