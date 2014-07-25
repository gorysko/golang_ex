package main

import "fmt"

type person struct {
    name string
    age int
}

func main() {

    fmt.Println(person{"bob", 20})

    fmt.Println(person{name: "alice", age: 34})

    fmt.Println(person{name: "shit"})

    fmt.Println(&person{name: "shot", age: 65})

    s := person{name: "Sean", age: 45}

    fmt.Println(s.name)

    s_pointer := &s

    fmt.Println(s_pointer.age)

    s_pointer.age = 34

    fmt.Println(s_pointer.age)
}
