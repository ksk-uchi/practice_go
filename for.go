package main

import "fmt"

func main () {
    // normal
    sum := 0
    for i:=0; i<10; i++ {
        sum += 1
    }
    fmt.Println(sum)

    // remove front, back
    sum2 := 1
    for ; sum2 < 1000; {
        sum2 += sum2
    }
    fmt.Println(sum2)

    // like while
    sum3 := 1
    for sum3 < 1000 {
        sum3 += sum3
    }
    fmt.Println(sum3)
}
