package main

import "fmt"

func main() {
    fmt.Println(concatStrings("Hello", ",", "World", "!"))
    fmt.Println(concatStrings("This", "is", "a", "test"))
}

func concatStrings(args ...string) string {
    var str string
    for _,v := range args {
        str += v
        str += " "
    }
    return str
}
