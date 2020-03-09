// try using the map data structure

package main

import "fmt"

func main() {
    var m = map[string]int {"one" : 1, "two" : 2, "three" : 3}
    printMap(m)

    // map seems to be pass by reference 
    modifyMap(m)
    printMap(m)

    // create a map dynamically
    var m1 = make(map[string]int)
    m1["test"] = 11
    m1["best"] = 22
    fmt.Println(m1)
}

func printMap(m map[string]int) {
    fmt.Printf("type = %T\n", m)
    for k, _ := range m {
        fmt.Printf("m[%s] ==> %d\n", k, m[k])
    }
}

// modify the map adding two keys and remove one
func modifyMap(m map[string]int) {
    m["four"] = 4
    m["five"] = 5
    delete(m, "one")
}
