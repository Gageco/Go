package main

import (
  "time"
  "fmt"
  "runtime"
)

func timer(num int, len int) {
  runtime.Gosched()
  t := RepoFindTimer(num)
  //t.Completed = false
  //t.Id = num
  t.Length = len

  fmt.Println("starting timer:",num)
  time.Sleep(time.Duration(len) * time.Second)
  fmt.Println("timer over:",num)

  t.Completed = true

}
