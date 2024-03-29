package main

import (
    "fmt"
    "math/cmplx"
)


var (
    ToBe bool = false
    MaxInt uint64 = 1<<64-1
    f float64 = 123455556499995.123 / 35325324234.666
    z complex128 = cmplx.Sqrt(-5 + 12i)
    MyStr string = "Hello World!"
)

func main() {
    fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
    fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
    fmt.Printf("Type: %T Value: %v\n", z, z)
    fmt.Printf("Type: %T Value: %v\n", f, f)
    fmt.Printf("Type: %T Value: %v\n", MyStr, MyStr)
}
