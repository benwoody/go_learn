package main

import (
  "fmt"
  "math"
)

// circleArea takes a single arguement, r, that is of type float64.
// The second invocation of float64 is the return type of the function.
func circleArea(r float64) (float64) {
  return math.Pi * r * r
}

func main() {
  fmt.Println("The area of a circle with a radius of 3 is:", circleArea(3))
}
