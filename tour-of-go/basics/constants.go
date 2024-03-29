package main

import "fmt"

const Pi = 3.14
const E = 2.78

const True = true
const False = false

func main() {
    const World = "世界"

    fmt.Println("Hello ", World)
    fmt.Println("Happy", Pi, "Day")

    fmt.Println("Euler's constant", E)
    fmt.Println("True and False are", True, False)

    // Try to change value of Pi to 3!
    // This will result in a compile error
    Pi = 3
    fmt.Println("Value of Pi is", Pi)
}
