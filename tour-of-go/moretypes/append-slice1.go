package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	for i:=0 ; i<10000 ; i++ {
		s = append(s, i)
		printSlice(s)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
