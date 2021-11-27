package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	q1 := q[2:5]
	fmt.Println(q1)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	r1 := r[1:4]
	fmt.Println(r1)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	s1 := s[2:6]
	fmt.Println(s1)
}
