package main

import (
    "fmt"
    "time"
)


func main() {
    timer_1 := time.NewTimer(time.Second * 2)

    <- timer_1.C
    fmt.Println("timer one expired")
    timer_2 := time.NewTimer(time.Second)
    go func() {
        <- timer_2.C
        fmt.Println("timer two expired")
    }()

    stop_2 := timer_2.Stop()
    if stop_2 {
        fmt.Println("timer two stopped")
    }
}
