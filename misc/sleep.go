package main

import (
  "time"
  "fmt"
)

func main() {
  fmt.Println("100 Milliseconds Later")
  time.Sleep(100 * time.Millisecond)
  fmt.Println("1 Second Later")
  time.Sleep(1 * time.Second)
  fmt.Println("This Prints Out, Now 1 minutes later")
  time.Sleep(1 * time.Minute)
  fmt.Println("This Prints Out")

  var waitTime int = 10
  time.Sleep(time.Duration(waitTime) * time.Second)

}
