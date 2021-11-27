package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    prev := 0
    curr := 1

    return func() int {
        var retval = 0
        if prev == 0 && curr == 1 {
            retval = 0
            prev, curr = curr, curr+prev
        } else if prev == 1 && curr == 1 {
            prev, curr = curr, curr+prev
            retval = 1
        } else {
            retval = prev
            prev, curr = curr, curr+prev
        }

        return retval
    }
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
