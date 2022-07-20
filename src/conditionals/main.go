package main

func main() {

}

func SimpleFunction(x, y int) int {
	if x == y { // == means testing whether two expressions are equal
		return 0
	} else if x > y {
		return 1
	} else {
		return -1
	}
}

func SameSign(x, y int) bool {
	if x*y >= 0 { // >= is "greater than or equal to"
		return true
	}
	// if we make it here, we know that x * y < 0
	return false
}

// when we declare a variable, it has a "scope" (lifetime) that exists inside
// of the "block" in which it lives.

func PositiveDifference(a, b int) int {
  var c int // declare c first
  if a == b {
    c = 0 // than we can assign the value, which doesn't require ":"
  } else if a > b {
    c = a-b
  } else {
    c = b-a
  }
  return c
}

// Other conditions
/*
>, <, >=, <=. ==

"Not equal to" !=
"not" !x (this evaluates to true when x is false )
*/
