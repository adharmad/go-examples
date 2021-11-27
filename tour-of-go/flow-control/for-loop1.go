package main

import "fmt"

func main() {
    sum := 1
    // another for loop
    // this is basically Go's while loop
    for ; sum < 1000 ; {
        sum += sum
        fmt.Println(sum)
    }

    sum = 1
    // making the "for" really like "while"
    for sum < 1000 {
        sum += sum
        fmt.Println(sum)
    }
}
