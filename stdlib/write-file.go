package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("test.txt")

    if err != nil {
        fmt.Println("Cannot create file")
        // handle error here
        return
    }
    defer file.Close()

    file.WriteString("To be or not to be\n")
    file.WriteString("This is a test file\n")
    file.WriteString("Hello World!\n")
}
