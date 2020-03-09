package main

import "fmt"

func main() {
	sum, diff := sum_and_diff(456, 178)
	fmt.Printf("sum = %d\ndifference = %d\n", sum, diff)
}

func sum_and_diff(x int, y int) (int, int) {
	return x + y, x - y
}

