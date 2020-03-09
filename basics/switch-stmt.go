package main

import (
    "fmt"
    "time"
    "runtime"
)

func main() {
	fmt.Print("Go runs on ")
    switch os := runtime.GOOS ; os {
        case "darwin":
            fmt.Println("OS X.")
        case "linux":
            fmt.Println("Tux.")
        default:
            // freebsd, openbsd
            // windows
            fmt.Printf("%s.\n", os)
    }

    // switch evaluation order
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}    
}
