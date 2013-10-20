package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main () {
    p := Vertex{1, 2}
    q := p
    r := &p
    q.X = 20
    fmt.Println(p, "p")
    fmt.Println(q, "q")

    r.Y = 30
    fmt.Println(p, "p")
    fmt.Println(r, "r")
}
