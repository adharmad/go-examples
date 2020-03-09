package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    prevprev := 0
    prev := 1

    return func() int {
        if prevprev == 0 && prev == 1 {
            prevprev, prev = prev, prevprev+prev
            return 0
        } else if prevprev == 1 && prev == 1 {
            prevprev, prev = prev, prevprev+prev
            return 1
        } else {
            prevprev, prev = prev, prevprev+prev
            return prevprev
        }
    }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

