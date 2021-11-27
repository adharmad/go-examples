package main

import "fmt"

type Person struct {
    name string
    age int
}

func main() {
    p1 := Person{name: "John", age: 25}
    fmt.Println(p1)

    // structs are pass-by-value
    // changing the value of a struct inside a function
    // does not change the original struct
    p2 := Person{name: "Harry", age: 67}
    fmt.Println("Before changing: ", p2)
    increaseAge(p2)
    fmt.Println("After changing: ", p2)

    // now pass the same struct using a pointer
    increaseAgePtr(&p2)
    fmt.Println("After changing: ", p2)
}

func increaseAge(p Person) {
    p.age++
}


func increaseAgePtr(p *Person) {
    p.age++
}

