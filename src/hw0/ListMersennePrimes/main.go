package main

import("fmt"; "math")

func main() {
    fmt.Println(IsPrime(60))
    fmt.Println(ListMersennePrimes(60))
}

func ListMersennePrimes(n int) []int {
    PrimeList := make([]int, 0)
    for p := 2; p <= n; p++ {
        if IsPrime(Power(2, p)-1){
            // fmt.Println(Power(2, p)-1)
            PrimeList = append(PrimeList, Power(2, p)-1)
        }
    }
    return PrimeList
}

func Power(n int, k int) int {
    ans := 1
    for i := 1; i <= k; i++ {
        ans *= n
    }
    return ans
}

func IsPrime(n int) bool {
    for i:= 2; float64(i) < math.Sqrt(float64(n)); i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}
