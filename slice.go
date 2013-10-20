package main

import "fmt"

func main () {
    p := []int{2, 3, 5, 7, 11, 13}
    q := [] float64 {1.5, 2.4, 3.3, 9.23}
    fmt.Println("p ==", p)
    fmt.Println("q ==", q)

    for i :=0; i < len(p); i++ {
        fmt.Printf("p[%d] == %d\n",
            i, p[i])
    }

    for j :=0; j < len(q); j++ {
        fmt.Printf("q[%d] == %f\n",
            j, q[j])
    }

    fmt.Println(p[3:5])


    var z []int
    fmt.Println(z, len(z), cap(z))
    if z == nil {
        fmt.Println("nil!")
    }
}
