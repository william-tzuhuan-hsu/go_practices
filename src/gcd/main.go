package main

import (
    // "fmt"
    "time"
    "log"
)

func main() {
    x := 378202873
    y := 273147834

    start := time.Now()
    TrivialGCD(x, y)
    elapsed := time.Since(start)
    log.Printf("TrivialGCD took %s", elapsed)

    start2 := time.Now()
    EuclidGCD(x, y)
    elapsed2 := time.Since(start2)
    log.Printf("EuclidGCD took %s", elapsed2)

}
/*
TrivialGCD(a, b)
  d <- 1
  m <- Min2(a, b)
  for every integer p from 1 to m
    if p is a divisor of a and b
      d <- p
  return d
*/

func TrivialGCD(a, b int) int {
    d := 1
    m := Min2(a, b)
    for p := 1; p <= m; p++ {
        // remainder operation Remainder(n, k) is n%k (e.g. 14%3 = 1)
        // if p is a divisor of a
        if a % p == 0 && b % p == 0{
            // if p is also a divisor of b
            d = p
        }
    }
    return d
}

// there is an "OR" of two statements as well; operator is ||
// e || f is true if one or both of e and f is true. Its' only false if
// both e and f are false.

/*
EuclidGCD(a, b)
    while a not equal to b
        if a > b
            a = a - b
        else // b > a
            b = b - a
    return a // or b
*/

func EuclidGCD(a, b int) int {
    for a != b {
        if a > b {
            a = a - b
        } else{
            b = b - a
        }
    }
    return a
}



func Min2(a, b int) int {
    if a > b {
        return b
    }
    return a
}
