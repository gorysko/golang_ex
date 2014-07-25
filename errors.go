package main

import (
    "fmt"
    "errors"
)

func f1(arg int) (int, error) {
    if arg == 42 {
        return -1, errors.New("can't work with 42s")
    }

    return arg + 3, nil
}

type arg_error struct {
    arg int
    prod string
}


func (e *arg_error) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prod)
}

func f2(arg int) (int, error) {
    if arg == 42 {
        return -1, &arg_error{arg, "can't work with it"}
    }

    return arg + 3 , nil
}

func main() {
    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }

    for _, i := range []int{13, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

    _, e := f2(42)

    if ae, ok := e.(*arg_error); ok {
        fmt.Println(ae.prod)
        fmt.Println(ae.arg)
    }

}
