package main

import (
    "fmt"
    "os"
)

func main() {
    file ,err := os.Open("file-stats.go")
    if err != nil {
        // handle error
        return 
    }
    defer file.Close()

    // file size
    stat, err := file.Stat()
    if err != nil {
        return
    }

    fmt.Println("File name: ", stat.Name())
    fmt.Println("File size: ", stat.Size())
    fmt.Println("File mode: ", stat.Mode())

    // read the file
    bs := make([]byte, stat.Size())
    _, err = file.Read(bs)
    if err != nil {
        return
    }

    str := string(bs)
    fmt.Println(str)

}
