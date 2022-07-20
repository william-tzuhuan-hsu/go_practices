package main

import("fmt")

func main() {
    fmt.Println(permutation(12, 1))
    fmt.Println(permutation(6, 6))
    fmt.Println(permutation(10, 4))
    fmt.Println(permutation(10, 0))
}

func permutation(n int, k int) int {
    ans := 1
    for i := n; i >= n-k+1; i-- {
        ans = ans*i
    }
    return ans
}
