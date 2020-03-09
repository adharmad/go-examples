package main

import "fmt"

func testDefer() {
    defer fmt.Println("after executing testDefer function")
    s := "World"
    fmt.Println("Hello", s)
}

func stackingDefers() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
    testDefer()
    stackingDefers()
}
