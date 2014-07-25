package main

import (
    "fmt"
    "math"
)

type geometry interface {
    area() float64
    perim() float64
}


type square struct {
    width, heigth float64
}

type circle struct {
    radius float64
}

func (s square) area() float64 {
    return s.width * s.heigth
}

func (s square) perim() float64 {
    return 2 * s.width + 2 * s.heigth
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    s := square{width: 15, heigth: 31}
    c := circle{radius: 31}

    measure(c)
    measure(s)
}
