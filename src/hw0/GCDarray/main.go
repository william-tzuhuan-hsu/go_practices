package main

import("fmt")

func main() {
     a := []int {12, 15, 9, 6, 24}
     fmt.Println(minArray(a))
     fmt.Println(GCDArray(a))
}

func GCDArray(a []int) int {
    min := minArray(a)
    GCD := 1
    for i := 1; i <= min; i++ {
        if DividesAll(a, i){
            GCD = i
        }
    }
    return GCD
}

func DividesAll(a []int, i int) bool {
    for j := 0; j < len(a); j++ {
        if a[j] % i != 0 {
            return false
        }
    }
    return true
}

func minArray(a []int) int {
    min := a[0]
    for i := 1; i < len(a); i++ {
        if a[i] < min {
            min = a[i]
        }
    }
    return min
}
