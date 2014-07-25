package main

import "fmt"

func main() {
     
     var a [5]int

     fmt.Println("empty array: ", a)

     a[4] = 100
     fmt.Println("set: ", a)
     fmt.Println("get by id: " , a[4])

     fmt.Println("len: ", len(a))
     
     b := [5]int{1, 2, 3, 4, 5}
     fmt.Println("other way: ", b)

     var two_d [2][3]int
     for i := 0; i < 2; i++ {
     	 for j := 0; j < 3; j++ {
	     two_d[i][j] = i + j
	 }
     } 
     
     fmt.Println("2d: ", two_d)


}