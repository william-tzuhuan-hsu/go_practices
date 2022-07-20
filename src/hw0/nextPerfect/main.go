package main

import("fmt")

func main() {
    fmt.Println(NextPerfectNumber(8128))
}

func NextPerfectNumber(n int) int {
    i := n+1
    for i > 0 {
        if IsPerfect(i) {
            return i
        }
    i++
    }
    return n
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
