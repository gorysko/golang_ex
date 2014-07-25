package main

import (
    "fmt"
    "time"
)

func main() {

    channel_1 := make(chan string)
    channel_2 := make(chan string)

    go func() {
        time.Sleep(time.Second * 1)
        channel_1 <- "one"
    }()
    go func() {
        time.Sleep(time.Second * 2)
        channel_2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {
            case msg1 := <-channel_1:
                fmt.Println("received", msg1)
            case msg2 := <-channel_2:
                fmt.Println("received", msg2)
        }
    }
}
