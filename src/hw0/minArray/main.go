package main

import("fmt")

func main () {
    a := []int {5, 2, 3, 4, 5, 5}
    fmt.Println(minArray(a))
}

func minArray(a []int) int {
    min := a[0]
    for i := 0; i < len(a); i++{
        if a[i] < min {
            min = a[i]
        }
    }
    return min
}
