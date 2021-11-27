package main

import "fmt"

func product(x, y int) (prod int) {
    prod = x * y
    return
}

func main() {
	fmt.Println("Product is", product(33, 47))
}

