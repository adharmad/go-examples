// An empty interface may hold values of any type. (Every type implements
// at least zero methods.)
//
// Empty interfaces are used by code that handles values of unknown type.
// For example, fmt.Print takes any number of arguments of type interface{}.
package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

// polymorphic methods that can take any type
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
