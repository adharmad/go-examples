// sum of array passing it to a function
package main

import "fmt"

func main() {
    var intarr = [5]int{100, 200, 300, 400, 500}
    var sum = sumOfArray(intarr)

    fmt.Printf("length of array = %d\n", len(intarr))
    fmt.Printf("sum of array = %d\n", sum)
}

func sumOfArray(x [5]int) int {
    s := 0
    for i := 0 ; i < len(x) ; i++ {
        s += x[i]
    }

    return s
}
