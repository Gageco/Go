package main
//import "fmt"
//import "sync"

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
  //t.Completed = false
  timers = append(timers, t)

  c := make(chan bool)
  go timer(&t, c)                   //so you can have multiple timers running at once
  t.Finished = <- c

  return t
  //So there is an issue here and it is that the timer() doesnt alter the variable in the
  //struct so you have no way of knowing when it finishes besides the fmt.Println. Look into
  //channels
  //https://stackoverflow.com/questions/18499352/golang-concurrency-how-to-append-to-the-same-slice-from-different-goroutines

}
