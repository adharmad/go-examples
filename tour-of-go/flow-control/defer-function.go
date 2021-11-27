// defer is similar to java's "finally"
// it executes a specific function at the end of the function

package main

import "fmt"

func main() {
    do_something(1)
    do_something(2)
    do_something(100)
}

func do_something(i int) string {
    defer second()

    if i == 1 {
        first()
        fmt.Println("An exit point")
        return "hello"
    } else if i == 2 {
        fmt.Println("test")
        return "world"
    } else {
        third()
        return "foo"
    }

}

func first() {
    fmt.Println("First!")
}

func second() {
    fmt.Println("Do this at the end!")
}

func third() {
    fmt.Println("Third!")
}
