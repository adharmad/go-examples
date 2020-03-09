package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
    // v is defined and used in the if/else statement only
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
        fmt.Printf("%g >= %g\n", v, lim)
    }

    // v cannot be used here
    //
	return lim
}

func main() {
    i := 1

    if i > 1 {
        fmt.Println("i is greater than 1")
    } else {
        fmt.Println("i is less than or equal to 1")
    }

    fmt.Println(pow(3, 2, 10), pow(3, 3, 10))

}
