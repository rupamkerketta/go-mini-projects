package main

import "fmt"

type shape interface {
    getArea() float64
}

type triangle struct {
    base float64
    height float64
}

func (t triangle) getArea() float64 {
    return (0.5 * t.base * t.height)
}

type square struct {
    sideLength float64
}

func (s square) getArea() float64 {
    return (s.sideLength * s.sideLength)
}

func printArea(s shape) {
    fmt.Printf("%v (sq. units)\n", s.getArea())
}

func main() {
    tri := triangle{height: 11.19, base: 19.77}
    sq := square{sideLength: 10}

    fmt.Printf("Area of the triangle with height: %v and base: %v is -> ", tri.height, tri.base)
    printArea(tri)

    fmt.Printf("Area of the square with sideLength: %v is -> ", sq.sideLength)
    printArea(sq)
}
