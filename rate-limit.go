package main

import (
    "time"
    "fmt"
)

func main(){

    requests := make(chan int, 5)
    for i:= 1;  i <= 5; i++ {
        requests <- i
    }

    close(requests)

    limiter := time.Tick(time.Millisecond * 200)

    for re := range requests {
        <- limiter
        fmt.Println("request", re, time.Now())
    }

    bursty_limiter := make(chan time.Time, 3)

    go func() {
        for i := 1; i < 3; i++ {
            for t := range time.Tick(time.Millisecond * 200) {
                bursty_limiter <- t
            }
        }
    }()

    bursty_request := make(chan int, 5)
    for i := 1; i <= 5; i++{
        bursty_request <- i
    }

    close(bursty_request)
    for req := range bursty_request {
        <- bursty_limiter
        fmt.Println("requests", req, time.Now())
    }
}
