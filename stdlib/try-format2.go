package main

import "fmt"

func main() {
	var (
		s1, s2, s3 string
		//i1, i2 int
	)

	fmt.Sscanf("one two three", "%s %s %s", &s1, &s2, &s3)
	fmt.Printf("You entered: %s - %s - %s\n", s1, s2, s3)
}

