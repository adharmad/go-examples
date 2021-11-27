// panic and recover are like throw and catch exceptions
// they need to be paired with "defer"
package main

import "fmt"

func main() {
    defer printError()
    panic("hello i am panicking!")

}

func printError() {
    str := recover()
    fmt.Println(str)
}
