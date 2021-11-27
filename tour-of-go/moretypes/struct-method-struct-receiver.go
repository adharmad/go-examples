package main

import (
    "fmt"
    "strconv"
)

type Person struct {
    name string
    age int
}

func (p Person) toString() string {
    return "Name: " + p.name + ", Age: " + strconv.Itoa(p.age)
}

func main() {
    p1 := Person{name: "John", age: 25}
    fmt.Println(p1)
    fmt.Println(p1)
    fmt.Println(p1.toString())

}
