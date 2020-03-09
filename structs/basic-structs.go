package main

import "fmt"

type Point struct {
    x int
    y int
}

func main() {
    p1 := Point{1, 2} // defines a point (1,2)
    p2 := &Point{3, 4} // pointer to a point (3,4)
    p3 := Point{x: 5, y: 6} // another way of defining a point
    p4 := Point{x: 7} // y=0 is implicit
    p5 := Point{} // x=0, y=0


    fmt.Println(p1, p2, p3, p4, p5)
}
