package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    bs, err := ioutil.ReadFile("read-file-ioutil.go")
    if err != nil {
        return
    }

    //str := string(bs)
    //fmt.Println(str)

    for _, line := range(bs) {
        fmt.Println(string(line))
    }
}
