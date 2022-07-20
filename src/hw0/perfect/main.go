package main

import("fmt")

func main() {
    fmt.Println(IsPerfect(1))
}

func IsPerfect(n int) bool {
    perfect := SumDivisors(n)
    if n == perfect {
        return true
    }
    return false
}

func SumDivisors(n int) int {
    sum := 0
    for i := 1; i < n; i ++ {
        if n % i == 0 {
            sum += i
        }
    }
    return sum
}
