package main

import (
  "time"
  "fmt"
  "runtime"
)

func timer(t int) {
  runtime.Gosched()
  fmt.Println("starting timer:",t)
  time.Sleep(30 * time.Second)
  fmt.Println("timer over",t)
}
