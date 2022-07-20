package main

import("fmt"; "math")

func main() {
    fmt.Println(NextTwinPrimes(616))
    // fmt.Println(IsPrime(4))
}

func NextTwinPrimes(n int) (int, int) {
    i := n+1
    for i >= n {
        if IsPrime(i) {
            if IsPrime(i+2) {
                return i, i+2
            }
        }
        i++
    }
    return i, i+2
}

func IsPrime(n int) bool {
    if n == 1 {
        return false
    }
    for i:= 2; float64(i) <= math.Sqrt(float64(n)); i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}
