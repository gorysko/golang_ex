package main

import "fmt"

func main() {

    s := make([]string, 3)
    fmt.Println("empty: ", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"

    fmt.Println("get all:", s)
    fmt.Println("get 3rd:", s[2])

    fmt.Println("length:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")

    fmt.Println("updated:", s)

    c := make([]string, len(s))

    copy(c, s)

    var l []string

    l = s[2:5]

    fmt.Println("slice1:", l)

    l = s[:5]

    fmt.Println("slice2:", l)

    l = s[2:]

    fmt.Println("slice3:", l)

    t := []string{"a", "b", "c"}
    fmt.Println("declared:", t)

    two_d := make([][]int, 3)

    for i := 0; i< 3; i++ {
        inner_len := i + 1
        two_d[i] = make([]int, inner_len)
        for j := 0; j < inner_len; j++ {
            two_d[i][j] = i + j
        }
    }

    fmt.Println("two_d:", two_d)

}
