package main

import("fmt")

func main() {
    fmt.Println(FactorialArray(18))
    fmt.Println(NonTrivialFactorialArray(18))
}

func Factorial (n int) int {
    ans := 1
    for i := 1; i <= n; i++ {
        ans *= i
    }
    return ans
}

func FactorialArray(n int) []int {
    facArr := make([]int, n+1)
    for i := 0; i <= n; i++ {
        facArr[i] = Factorial(i)
    }
    return facArr
}

func NonTrivialFactorialArray(n int) []int {
    facArr := make([]int, n+1)
    fac := 1
    facArr[0] = fac
    for i := 1; i <= n; i++ {
        fac *= i
        facArr[i] = fac
    }
    return facArr
}
