package main

import (
    "fmt"
    "strconv"
)

type Person struct {
    first string
    last string
    full string
    age int
}

func (p *Person) toString() string {
    return "First: " + p.first + ", Last: " + p.last + " , Full: " + p.full + ", Age: " + strconv.Itoa(p.age)
}

// A method with a pointer receiver can change the pointer
func (p *Person) computeFullName() {
    p.full = p.first + "-" + p.last
}

func main() {
    p1 := Person{first: "John", last: "Doe", age: 25}
    fmt.Println(p1)
    fmt.Println(p1.toString())

    p1.computeFullName()
    fmt.Println(p1.toString())

}
