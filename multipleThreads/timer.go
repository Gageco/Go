package main

import (
  "time"
  "fmt"
  "runtime"
)

func timer(num int, len int) {
  runtime.Gosched()
  fmt.Println("starting timer:",num)
  time.Sleep(time.Duration(len) * time.Second)
  fmt.Println("timer over",num)
}
