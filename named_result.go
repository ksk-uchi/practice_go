package main

import "fmt"

func split(sum int) (x, y int) {
    // 小数点は切り捨て
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(split(17))
}
