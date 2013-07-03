package circle

import "testing"

func Test(t *testing.T) {
  var tests = []struct {
    radius, area float64
  }{
    {2, 12.566370614359172},
    {3, 28.274333882308138},
  }
  for _, c := range tests {
    got := Area(c.radius)
    if got != c.area {
      t.Errorf("Area(%f) == %f, wanted %f", c.radius, got, c.area)
    }
  }
}
