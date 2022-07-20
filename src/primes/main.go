package main

import (
    "fmt"
    "math"
    "log"
    "time"
)

func main() {
    fmt.Println("Primes and arrays.")

    // arrays in Go have a fixed, constant size.
    var list [6]int //gives 6 defualt values
    // var patterns [28]string // gives 28 "" (empty strings)
    list[0] = -8

    i := 3
    list[2*i-4] = 17
    list[len(list)-1] = 43
    // compiler error for stuff like this
    // list[len(list)] = 5
    // list[-4] = 0
    fmt.Println(list)

    n := 1000000

    start := time.Now()
    TrivialPrimeFinder(n)
    elapsed := time.Since(start)
    log.Printf("TrivialPrimeFinder took %s", elapsed)

    start2 := time.Now()
    SieveOfEratosthenes(n)
    elapsed2 := time.Since(start2)
    log.Printf("SieveOfEratosthenes took %s", elapsed2)



}

// remember that arrays have constant size in Go.
// we use something called a "slice" to represent variable sizes in Go.
// even if you're plugging in a variable into the length of the array and will never change it.

// TrivialPrimeFinder takes an integer n and produces a boolean array of length
// n+1 wehre PrimeArray[p] is true if p is prime (and false otherwise).
func TrivialPrimeFinder(n int) []bool {
    // var PrimeArray [n+1]bool can't compile
    var primeArray []bool // type is "slice of booleans", default to false
    // but we still nedd an initial length.
    // slices start as nil and need to be made
    primeArray = make([]bool, n+1)
    // in geneeral, you'll just use primeArray := make([bool, n+1])

    for p := 2; p <= n; p++ {
        if IsPrime(p) == true {
            primeArray[p] = true
        }
    }
    return primeArray
}

func IsPrime(p int) bool {
    for k := 2; float64(k) <= math.Sqrt(float64(p)); k++ {
        if p%k == 0 { // p is not prime
            return false
        }
    }
    // if we survive testing all these factors, p is prime
    return true
}


// Sieveof Eraosthenes takes an integer n and returns a slice of n+1 booleans
// primeArray where primeArray[p] is true if p is prime and false otherwise.
func SieveOfEratosthenes(n int) []bool {
    primeArray := make([]bool, n+1)

    // set everything to prime other than 0 and 1
    for k := 2; k <= n; k++ {
        primeArray[k] = true
    }

    // now, range over primeArray, and cross off multiples of first priem we see and iterate this process
    for p := 2; float64(p) <= math.Sqrt(float64(n)); p++ {
        if primeArray[p] == true {
            primeArray = CrossOffMultiples(primeArray, p)
        }

    }
    return primeArray
}

// CrossOffMultiples takes a boolean slice and an integer p. It crosses off
// multiples of p, meaning turning these multiples to false in the slice.
// It returns the slice obtained as a result.

func CrossOffMultiples(primeArray []bool, p int) []bool {
    n := len(primeArray) - 1
    for k := 2*p; k <= n; k += p {
        // all these multiples should be made composite
        primeArray[k] = false
    }
    return primeArray
}
