package main

import (
	"fmt"
	"math/rand"
)

// This function will return a slice of length dy,
// each element of which is a slice of dx 8-bit unsigned integers.
// xpos and ypos are the beginning of the slice
// No bounds checks are performed
func twod_slice(xpos, ypos, dx, dy int, arr [][]int) [][]int {
	// This is throwing runtime error
	return arr[xpos:xpos+dx][ypos:ypos+dy]
}

func main() {
	arr2d := make([][]int, 10)
	for i:=0 ; i<10 ; i++ {
		arr2d[i] = make([]int, 10)

		for j:=0 ; j<10 ; j++ {
			arr2d[i][j] = rand.Intn(100)
		}
	}

	print2dArr(arr2d)

	arr2d_slice := twod_slice(1, 1, 3, 3, arr2d)
	print2dArr(arr2d_slice)
}

func print2dArr(arr [][]int) {
	for i:=0 ; i<10 ; i++ {
		for j:=0 ; j<10 ; j++ {
			fmt.Printf("%2d ", arr[i][j])
		}
		fmt.Println()
	}
}
