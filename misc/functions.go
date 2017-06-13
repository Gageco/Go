package main

import "fmt"

func returnString() string {
  return "This is a returned string."
}

func returnInt() uint8 {
  return 1
}

func returnToSender(num uint8) uint8 {
  return num-1
}

func main() {
  var str string
  var myInt uint8

  str = returnString()
  myInt = returnInt()

  fmt.Println("Returned String: ", str)
  fmt.Println("Returned Int: ", myInt)
  fmt.Println("Return 10-1: ", returnToSender(10))
}
