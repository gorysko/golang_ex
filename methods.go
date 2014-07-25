package main

import "fmt"

type rect struct {
    width, height int
}

func (r *rect) area_pointer() int {
    return r.width * r.height
}

func (r *rect) area() int {
    return r.width * r.height
}


func (r rect) perim() int {
    return 2*r.width + 2*r.height
}


func main() {

    r := rect{width: 20, height: 30}

    fmt.Println(r)

    fmt.Println("r.area_pointer:", r.area_pointer())
    fmt.Println("r.area", r.area())
    fmt.Println("r.perim:", r.perim())

    r_pointer := &r

    fmt.Println("r_pointer.area_pointer:", r_pointer.area_pointer())
    fmt.Println("r_pointer.area:", r_pointer.area())
    fmt.Println("r_pointer.perim:", r_pointer.perim())
}
