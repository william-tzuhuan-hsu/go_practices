package main

import (
  "fmt"
)

func main() {
  fmt.Println("Loops.")

  n := 5
  m := Factorial(n)

  fmt.Println(m)

  // trigger the panic
  // fmt.Println(Factorial(-100))

  fmt.Println(Fibonacci(10))

  fmt.Println(SumEven(10))

  // Since uint would never reach 0, it would create an infinite loop
  // Starting from the largest integer
  // var i uint = 10
  //
  // for ; i >= 0; i-- {
  //   fmt.Println(i)
  // }
}

//first, a factorial functions

func Factorial(n int) int {
  // let's handle negative input
  if n < 0 {
    // panic() will print an error message and immidiately end the program
    panic("Error: Negative input given to Factorial()!")
  }

  p := 1
  i := 1
  // Go doesn't have a while keyword. They use "for"
  // while i <= n
  for i <= n{
    p = p*i
    i = i+1
  }
  return p
}

// Exercise: write a function in Go using a while loop that takes an integer n
// and returns the sum of the first n postitive integers/
func Fibonacci(n int) int {
  if n < 0 {
    panic("Error: Negative input given to Fibonacci()!")
  }
  sum := 0
  i := 1

  for i <= n{
    sum = sum+i
    i ++ // i = i++, there is also i-- which means i = i-1
  }
  // live still lives here and we could use it
  return sum
}


func AnothreFactorial(n int) int {
  if n < 0 {
    // panic() will print an error message and immidiately end the program
    panic("Error: Negative input given to Factorial()!")
  }
  p := 1
  // for every integer i between 1 and n, p = p*i
  for i := 1; i <= n; i++ {
      p *= i
      // the scope of i would not exceed this for loop
  }
  // i doesn't live here
  // it's good that we don't have zombie variable this way
  return p
}

func AnotherSum(n int) int {
  if n < 0 {
    // panic() will print an error message and immidiately end the program
    panic("Error: Negative input given to Factorial()!")
  }
  sum := 0

  for k := 1; k <= n; k++{
    sum += k
  }

  return sum
}

// Exercise: write a function SumEven that sums all even numbers up to and possibly including n.
func SumEven (n int) int {
  sum := 0
  for i := 0; i <= n; i += 2{
      sum += i
  }
  return sum
}
