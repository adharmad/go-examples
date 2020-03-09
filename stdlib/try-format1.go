package main

import "fmt"

func main() {
	i := 10
	s := "test"

	fmt.Print("Print this\n")
	fmt.Println("Print this too!")
	fmt.Printf("This is a formatted string %s %d\n", s, i)
	s1 := fmt.Sprintf("This is a formatted string %s %d\n", s, i)

	fmt.Printf("Returned string: %s", s1)
}

