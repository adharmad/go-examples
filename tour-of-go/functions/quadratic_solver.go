package main

import (
    "os"
    "fmt"
    "math"
    "strconv"
)

func main() {
    var a,_ = strconv.ParseFloat(os.Args[1], 64)
    var b,_ = strconv.ParseFloat(os.Args[2], 64)
    var c,_ = strconv.ParseFloat(os.Args[3], 64)

    r1, r2 := quadratic_solution(a, b, c)

    fmt.Println("Root1 = ", r1)
    fmt.Println("Root2 = ", r2)
}

func quadratic_solution(a float64, b float64, c float64) (complex128, complex128) {
    var d float64 = math.Pow(b, 2) - 4*a*c
    var r1, r2 complex128

    if d < 0 {
        d = math.Abs(d)
        r1 = complex(-b/2*a, math.Sqrt(d)/2*a)
        r2 = complex(-b/2*a, -math.Sqrt(d)/2*a)
    } else {
        r1 = complex((-b+math.Sqrt(d))/2*a, 0)
        r2 = complex((-b-math.Sqrt(d))/2*a, 0)
    }

    return r1, r2
}
