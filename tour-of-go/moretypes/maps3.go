package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// Define the variable as a map literal
// with inline struct literals
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
