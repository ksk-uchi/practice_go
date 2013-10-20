package main

import (
    "fmt"
    "math/rand"
    "math"
)

func main() {
    fmt.Println("My favorite number is", rand.Intn(10))
    fmt.Println(math.Pi)

    res := make([][]int, 5, 5)
    for i:=0; i<len(res); i++ {
        res[i] = make([]int, i+1, i+1)
    }
    fmt.Println(res)
}
