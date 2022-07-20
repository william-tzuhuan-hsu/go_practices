package main

import("fmt"; "math")

func main(){
    fmt.Println("More on arrays and slices")
    a := make([]int, 3) // slice with length 3
    var b [3]int //array with length 3

    ChangeFirst1Slice(a)
    ChangeFirst1Array(b) // In the case of array, the first value won't be change

    fmt.Println(a)
    fmt.Println(b)

    a = append(a, 5) // this adds element 5 to the right side of the slice
    fmt.Println(a)

    n := 31
    primes := ListPrimes(n)
    fmt.Println(primes)
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

func IsPrime(p int) bool {
    for k := 2; float64(k) <= math.Sqrt(float64(p)); k++ {
        if p%k == 0 { // p is not prime
            return false
        }
    }
    // if we survive testing all these factors, p is prime
    return true
}


func Max(list []int) int {
    if len(list) == 0 {
        panic("Error: Empty list passed to Max().")
    }
    var m int  // default value = 0

    // range ove list, updating m if we find bigger value
    for i, val := range list { // equivalent to for i := 0; i < len(list); i++
        if i == 0 || val > m {
            m = val
        }
    }
    return m
}


// Sum takes a slice of integers and returns the sum of
// values in the slice.
func Sum(a []int) int {
    var s int

    for _, value := range a { // _ says "I don't need this variable"
        s += value
    }

    return s
}

// passing a slice as a parameter is also possible

func ChangeFirst1Slice(list []int) {
    list[0] = 1 // we can edit a slice by passing it into a function
}

func ChangeFirst1Array(list [3]int) {
    list[0] = 1 //list is a copy of whatever we call function on
}

// ListPrimes takes an integer n and returns a list of all prime numbers up to and including n.
func ListPrimes(n int) []int {
    primes := make([]int, 0) // great if you don't know eventual length of primes
    primeBooleans := SieveOfEratosthenes(n) // gives me a slice of boolean variables
    // whose p-th elemt is true is p is prime
    for p := range primeBooleans {
        // when I encounter a prime, append it to primes
        if primeBooleans[p] == true {
            primes  = append(primes, p)
        }
    }
    return primes
}
