package main

import("fmt")

func main() {
    fmt.Println(FastCombination(1000, 999))
}

func FastCombination(n, k int) int {
    nom := 1
    if k == 0 || n == k {
        return nom
    }
    if k >= n-k {
        for i := n; i >= k+1; i-- {
            nom *= i
        }
        return nom/Factorial(n-k)
    }else{
        for i := n; i >= n-k+1; i--{
            nom *= i
        }
        return nom/Factorial(k)
    }
}

func Factorial (n int) int {
    ans := 1
    for i := 1; i <= n; i++ {
        ans *= i
    }
    return ans
}
