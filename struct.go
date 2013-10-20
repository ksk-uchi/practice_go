package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{Y: 1, X: 2}
    q := &Vertex{22, 34}
    v.X = 4
    fmt.Println(v)
    fmt.Println(q)
}
