package main

import "fmt"

func main() {
   do_something() 
}

func do_something() {
    fmt.Println("Hello from function")

    fmt.Println("Sum of 42 and 37 is", add(42, 37))
}

func add(x int, y int) int {
    return x + y
}
