package main

import "fmt"

func sum(s []string, c chan string) {  
	c <- s[0] + "-" + s[1]
}

func main() {
	s := []string{"This", "is", "A", "Test"}

	c := make(chan string)
    go sum(s[0:2], c)
    go sum(s[2:4], c)
	s1 := <-c // receive from c
	s2 := <-c // receive from c

	fmt.Println(s1)
	fmt.Println(s2)
}

