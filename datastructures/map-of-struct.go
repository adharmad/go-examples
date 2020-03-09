package main

import "fmt"

type User struct {
    first string
    last string
    age int
}

func main() {

    var m map[string]User = make(map[string]User)
    m["amol"] = User{"Amol", "Dharmadhikari", 45}
    m["vishal"] = User{"Vishal", "Mehta", 44}
    m["eshwar"] = User{"Eshwar", "Sundaresan", 999}

    printMap(m)

}

func printMap(m map[string]User) {
    fmt.Printf("type = %T\n", m)
    for k, _ := range m {
        fmt.Printf("m[%s] ==> %v\n", k, m[k])
    }
}
