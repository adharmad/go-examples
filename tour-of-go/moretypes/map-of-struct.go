package main

import "fmt"

type User struct {
    first string
    last string
    age int
    address string
}

func main() {

    var m map[string]User = make(map[string]User)
    m["amol"] = User{"Loma", "Irakihdamrahd", 45, "Sunnyvale"}
    m["vishal"] = User{"Lahsiv", "Athem", 44, "Cupertino"}
    m["eshwar"] = User{"Rawhse", "Naseradnus", 999, "Bangalore"}

    printMap(m)

}

func printMap(m map[string]User) {
    fmt.Printf("type = %T\n", m)
    for k, _ := range m {
        fmt.Printf("m[%s] ==> %v\n", k, m[k])
    }
}
