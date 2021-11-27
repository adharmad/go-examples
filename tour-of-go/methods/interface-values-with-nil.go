package main

import "fmt"

type Producer interface {
	produce()
}

type Cow struct {
	name string
}

type Chicken struct {
	name string
}

type Tiger struct {
	name string
}

func (c *Cow) produce() {
    fmt.Printf("Cow %s produces milk\n", c.name)
}

func (c *Chicken) produce() {
    fmt.Printf("Chicken %s produces eggs\n", c.name)
}

func (t *Tiger) produce() {
    if t == nil {
        fmt.Println("Nil Tiger cannot produce anything!!")
    } else {
        fmt.Printf("Tiger %s produces ROAR\n", t.name)
    }
}

func main() {
    var c Cow = Cow{"mycow"}
    var ch Chicken = Chicken{"mychicken"}
	var t *Tiger

    c.produce()
    ch.produce()
    t.produce()
}

