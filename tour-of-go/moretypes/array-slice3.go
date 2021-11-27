package main

import "fmt"

func main() {
    // define an array
    var intarray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println("Array is ", intarray)

    // slice the array 
    var slice1 = intarray[0:5]
    var slice2 = intarray[2:8]
    fmt.Println("Slice1 is ", slice1)
    fmt.Println("Slice2 is ", slice2)

    // changing the slice changes the array
    strarray := [5]string{"This", "is", "a", "big", "test"}
    fmt.Println("String array is", strarray)

    s1 := strarray[0:1]
    s2 := strarray[3:4]
    fmt.Println("s1 and s2 are", s1, s2)

    s1[0] = "That"
    s2[0] = "small"
    fmt.Println("String array is", strarray)

    s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

    s = s[:0]
    printSlice(s)

    // Extend the slice
    s = append(s, 9, 10)
    fmt.Println(s)
    printSlice(s)
}

// print slice and capacity
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

