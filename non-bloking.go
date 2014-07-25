package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)

    select {
    case msg := <- messages:
        fmt.Println("receives message: ", msg)
    default:
        fmt.Println("no messages received")
    }

    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message: ", msg)
    default:
        fmt.Println("no messages")
    }

    select {
    case msg := <- messages:
        fmt.Println("received message", msg)
    case sig := <- signals:
        fmt.Println("recived signal", sig)
    default:
        fmt.Println("no activity")
    }
}
