package main

import "fmt"

func main() {
    // define an array
    var intarray = [5]int{11, 22, 33, 44, 55}
    sum := 0

    fmt.Printf("Length of array = %d\n", len(intarray))

    for i := 0 ; i < len(intarray) ; i++ {
        v := intarray[i]
        fmt.Printf("array element[%d] = %d\n", i, v)
        sum += v
    }

    fmt.Printf("Sum of array elements = %d\n", sum)

    // another way of creating an array using make function
    var anotherarray = make([]float64, 10)
    fmt.Println("Another array = ", anotherarray)

    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)
}
