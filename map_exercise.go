package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    words := strings.Fields(s)
    res := map[string]int{}
    for _, v := range words {
        res[v]++
    }
    return res
}

func main () {
    wc.Test(WordCount)
}
