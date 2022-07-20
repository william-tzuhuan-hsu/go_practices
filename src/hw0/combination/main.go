package main

import("fmt")

func main() {
    fmt.Println(Combination(12, 1))
}

func Combination(n int, k int) int {
    return Permutation(n, k)/Factorial(k)
}


func Factorial (n int) int {
    ans := 1
    for i := 1; i <= n; i++ {
        ans *= i
    }
    return ans
}

func Permutation(n int, k int) int {
    ans := 1
    for i := n; i >= n-k+1; i-- {
        ans = ans*i
    }
    return ans
}
