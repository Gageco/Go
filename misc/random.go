package main

import (
  // "time"
  "fmt"
  "math/rand"
)

func main() {
  randNum := rand.Intn(15) //random number between 0 and 15
  fmt.Println("Random Number between 0 and 15: ", randNum)

  randomFloat := rand.Float64 // random less than 0 float
  fmt.Println(randomFloat)
}
