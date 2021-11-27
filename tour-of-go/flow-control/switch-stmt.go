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
            // freebsd, openbsd, windows
            fmt.Printf("%s.\n", os)
    }
}
