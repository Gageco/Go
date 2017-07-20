package main

import (
  "time"
  "fmt"
  "runtime"
  "net/http"
  //"encoding/json"
)

func timer(t Timer, w http.ResponseWriter, theTimer *Timer, channel chan bool)  {
  runtime.Gosched()
  lengthOfTime := t.Length

  // fmt.Println(theTimer)

  // fmt.Println("starting timer:",t.Id)
  time.Sleep(time.Duration(lengthOfTime) * time.Second)
  // fmt.Println("timer over:",t.Id)
  theTimer.Finished = true
  //timers = append(timers, t)
  // fmt.Println(theTimer)
  channel <- true

}
