package main

import "fmt"

func main() {
    // define an array
    var intarray = []int{11, 22, 33}
    fmt.Println("intarray = ", intarray)

    // use function append to append to an array
    var anotherarray = append(intarray, 44, 55, 66)
    fmt.Println("anotherarray = ", anotherarray)


    var arr1 = []int{1, 2, 3}
    var arr2 = []int{4, 5}

    // use copy function to copy an array from destination
    // to source
    copy(arr2, arr1)
    fmt.Println(arr1, arr2)

    // try copy again but with destination and source reversed
    arr1 = []int{1, 2, 3}
    arr2 = []int{4, 5}
    copy(arr1, arr2)
    fmt.Println(arr1, arr2)
}

