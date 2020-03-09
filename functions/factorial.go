// example of recursive function using factorial

package main

import "fmt"

func main() {
    var i uint64 = 1
    for ; i < 16 ; i++ {
        fact := factorial(i)
        fmt.Printf("%d! = %d\n", i, fact)
    }

}

func factorial(n uint64) uint64 {
    if (n == 0) {
        return 1
    } else {
        return n * factorial(n-1)
    }
}
