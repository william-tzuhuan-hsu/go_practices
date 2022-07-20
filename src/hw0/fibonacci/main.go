package main

import("fmt")

func main() {
    fmt.Println(FibonacciArray(10))
}

func FibonacciArray(n int) []int {
    fibArr := make([]int, n+1)
    if n == 0 {
        fibArr[0] = 1
        return fibArr
    }else if n == 1 {
        fibArr[0] = 1
        fibArr[1] = 1
        return fibArr
    }else {
        for i := 2; i <= n; i++ {
            fibArr[0] = 1
            fibArr[1] = 1
            fibArr[i] = fibArr[i-1] + fibArr[i-2]
        }
    }
    return fibArr
}
