package main

import (
    "fmt"
    "math/rand"
    "time"
)

func f(n int) {
    for i := 0; i < 10; i++ {
        fmt.Println(n , ":", i)
        amt := time.Duration(rand.Intn(250))
        time.Sleep(time.Millisecond * amt)
    }
}

func main() {
    for n :=0; n < 10; n++ {
        go f(n)
    }
    var input string
    fmt.Scanln(&input)
}
