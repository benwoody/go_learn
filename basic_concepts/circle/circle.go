package circle

import "math"

// Area takes a single arguement, r, that is of type float64.
// The second invocation of float64 is the return type of the function.
func Area(r float64) (float64) {
  return math.Pi * r * r
}

