package main

import "fmt"

type User struct {
    name string
    age int
}

func main() {
    users := []User {
        {"Amol", 45},
        {"Vishal", 44},
        {"Eshwar", 999},
    }

    fmt.Println(users)

    // the struct can be anonymous
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
}
