package main

import "fmt"

func main() {
    // map with nested array
    var m = map[string][5]int {
        "one" : [5]int{1, 2, 3, 4, 5},
        "two" : [5]int{6, 7, 8, 9, 10},
        "three" : [5]int{11, 12, 13, 14, 15},
    }
    fmt.Println(m)

    // map with nested map
    var m1 = map[string]map[string]string {
        "CA" : map[string]string {
            "name" : "California",
            "capital" : "Sacremento",
        },
        "NY" : map[string]string {
            "name" : "New York",
            "capital" : "Albany",
        },
    }
    fmt.Println(m1)
}
