package main

import "fmt"

var powof2 = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

func main() {
    for idx, value := range powof2 {
        fmt.Printf("2^%d = %d\n", idx, value)
    }
}
