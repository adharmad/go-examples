package main

import (
    "fmt"
    "math"
)

type Shape interface {
    area() float64
}

type Circle struct {
    radius float64
}

type Rectangle struct {
    len float64
    bre float64
}

type Square struct {
    side float64
}

func (c *Circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (r *Rectangle) area() float64 {
    return r.len * r.bre
}

func (s *Square) area() float64 {
    return s.side * s.side
}

func computeAreas(s ...Shape) {
    for _, s := range s {
        fmt.Println("Area = ", s.area())
    }
}

func main() {
    // create a few shapes
    c1 := Circle{radius: 2.3}
    c2 := Circle{radius: 3.4}
    r1 := Rectangle{len: 1.1, bre: 2.2}
    r2 := Rectangle{len: 5.3, bre: 2.1}
    s1 := Square{side: 7.5}
    s2 := Square{side: 6.6}

    computeAreas(&c1, &c2, &r1, &r2, &s1, &s2)

}
