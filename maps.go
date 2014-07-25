package main

import "fmt"

func main() {
    m := make(map[string]int)

    m["k1"] = 1
    m["k2"] = 2

    fmt.Println("map:", m)

    v1 := m["k1"]

    fmt.Println("v1: ", v1)

    fmt.Println("LEN:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    value, present := m["k2"]
    fmt.Println("value:", value)
    fmt.Println("present:", present)

    n := map[string]int{"first": 1, "second": 2}
    fmt.Println("n:", n)

}
