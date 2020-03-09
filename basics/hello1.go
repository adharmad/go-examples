package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    fmt.Println("This is a number:", 1)
    fmt.Println("This is another number:", 3.14)

    fmt.Println("Time now is", time.Now())

    fmt.Println("Some random number", rand.Intn(10))
}
