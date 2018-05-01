package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X   float64
  Y   float64
  Z   string
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4, "test"}
	fmt.Println(v.Abs())
  fmt.Println(v.Z)
}
