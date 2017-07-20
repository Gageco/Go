package main
import (
  "net/http"
  // "fmt"
)

var timerIDs int
var timers Timers
var alters AlterStructs

func RepoFindTimer(id int) Timer {
  for _, t := range timers {
    if t.Id == id {
      return t
    }
  }
  return Timer{}
}

func RepoCreateTimer(t Timer, w http.ResponseWriter) Timer {
  timerIDs += 1;
  // t.Id = timerIDs
  timers = append(timers, t)
  //fmt.Println(timers[0])
  channel := make(chan bool)
  go timer(t, w, &timers[timerIDs-1], channel)                   //so you can have multiple timers running at once
  timers[timerIDs-1].Finished = <- channel
  return t
}

func RepoStateChange(a AlterStruct) AlterStruct {

  for i := 0; i < len(alters); i++ {
    if a.Id == alters[i].Id {
      alters[i].State = a.State
      // fmt.Println("Test Pass")
      return a
    } else {
      // fmt.Println("Test Fail")
    }
  }

  alterID += 1;
  a.Id = alterID
  alters = append(alters, a)
  // fmt.Println("At End")

  return a
}
