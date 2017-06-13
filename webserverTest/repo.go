package main
//import "fmt"

var currentId int

var timers Timers

func RepoFindTimer(id int) Timer {
  for _, t := range timers {
    if t.Id == id {
      return t
    }
  }
  return Timer{}
}

func RepoCreateTimer(t Timer) Timer {
  currentId += 1;
  t.Id = currentId
  t.Completed = false
  timers = append(timers, t)
  go timer(currentId, currentId+10)
  return t
}
