package main

import "fmt"

func main() {
    a1 := map[string]int{
        "0": 100,
        "1": 200,
    }
    _=a1
    fmt.Println(a1["0"]+a1["0"])
}