package main

import (
    "time"
    "fmt"
)


func main() {
    channel_1 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        }()
    select{
    case result := <- channel_1:
        fmt.Println(result)
    case <- time.After(time.Second * 1):
        fmt.Println("timeout 1")
    }

    channel_2 := make(chan string, 1)
    go func() {
        time.Slepp(time.Second * 2)
    }()

    select {
    case result := <- channel_2:
        fmt.Println(result)
    case <- time.After(time.Second * 3):
        fmt.Println("timeout 2")
    }
}
