package main

import (
  "time"
  "fmt"
  "runtime"
)

func timer(t *Timer, c chan bool)  {
  runtime.Gosched()
  lengthOfTime := t.Length

  fmt.Println("starting timer:",t.Id)
  time.Sleep(time.Duration(lengthOfTime) * time.Second)
  fmt.Println("timer over:",t.Id)
  c <- true


}
