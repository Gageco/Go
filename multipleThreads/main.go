package main

import (
  "time"
)

func main() {
  go timer(1)
  time.Sleep(5 * time.Second)
  go timer(2)
  time.Sleep(1 * time.Minute) // i have this so it doesnt quit before the program finishes, it  addstime ot the end of the program
}
