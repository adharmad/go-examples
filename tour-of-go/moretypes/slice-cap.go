package main

import "fmt"

func main() {
	var s []int
	oldcap := 0

	for i:=0 ; i<100000 ; i++ {
		s = append(s, 0)
		if oldcap != cap(s) {
			// Print out the capacity and length, but only if it has
			// changed
			oldcap = cap(s)
			fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
		}
	}
}
