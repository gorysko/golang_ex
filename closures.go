package main

import "fmt"

func init_seq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

func main() {
    next_int := init_seq()

    fmt.Println(next_int())
    fmt.Println(next_int())
    fmt.Println(next_int())
    fmt.Println(next_int())

    next_ints := init_seq()

    fmt.Println(next_ints())
}
