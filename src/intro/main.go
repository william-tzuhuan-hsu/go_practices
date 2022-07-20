package main

import "fmt"
// Go won't read this line

/*
Everything here won't be read either (multi-line comment)
Today: we'll see four variable types
int: integers
unit: nonnegative integers
bool: true or false Boolean variable
float64: decimal number

next time:
byte: single symbol
string contiguous collections of symbols (words)
*/

func main() {
  fmt.Println("Let's get started!")

  //declarance of variables
  var j int = 14// j is an integer variable
  var x float64 = -2.3
  var yo string = "Hi" // default value for string is ""
  var u uint = 14
  var symbol byte = 'H'
  var statement bool = true // defualt to false

  fmt.Println(j)
  fmt.Println(x)
  fmt.Println(yo)
  fmt.Println(u)
  fmt.Println(symbol)
  fmt.Println(statement)

  // shorthand declaration
  i := -6 //automatically type int (if u want a unit, declare it)
  hi := "Yo " // Go automatically give this type string
  k := 34 //automatically an integer

  // we can do arithmetic on numeric variables

  fmt.Println((i+j)*2*k)
  fmt.Println(2*x - 3.16)

  fmt.Println(hi + yo)

  fmt.Println(j/i) // Go views this as integer divisor  .. throws away the remainder

  // if we want the actual value k/j, we use type conversions
  fmt.Println(float64(k)/float64(j))

  // var ok bool = bool(0) // false? No ... compiler error

  // fmt.Println(ok)

  var p int = -187
  var s uint = uint(p)  // print out a very large number

  fmt.Println(s)

  m := 9223372036854775807 // integer overlfow to the smallest integer
  fmt.Println(m+1)

  fmt.Println(float64(j)*x)

  /* Go demands that input variables have correct type
  w, z := DoubleAndDupicate(m)
  fmt.Println(w, z)
  */

  n := 17
  fmt.Println(AddOne(n))
  fmt.Println(n)

}

func AddOne(n int) int {
  // the variable n gets created as a copy of whatever is given
  n = n+1
  return n  // we are returning a value
  // the new n is now destroyed at the end of function
}

// SumTwoInts takes two integers and returns their sum.
func SumTwoInts(a int, b int) int {
  return a + b
}

func DoubleAndDupicate(x float64) (float64, float64) {
  return 2.0*x, 2.0*x
}

func Pi() float64 { // doesn't take inputs
  return 3.14  // wrong haha
}

// functions can also not return anything
// similar like func main() {}
func PrintHi() {
  fmt.Println("Hi")
}
