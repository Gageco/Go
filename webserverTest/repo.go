package main
import (
  "net/http"
  "math/rand"
  // "fmt"
)

var timerIDs int
var timers Timers
var alters AlterStructs
var givenAuth validAuth

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

/*
So this whole auth thing isnt exactly perfect or good or anything just kinda a fun test.

It works for the most part which is neat and so ill be able to authorize my IoT applications but
the whole concept needs some work because this is still pretty basic

Like:
- Use actual seeds
- Have a reset so if its wrong you dont have to worry about it
- Make thread safe
- Add more users and junk

*/

func RepoAuthTest(a AuthStruct) int {
  validUser := "gage"
  validPass := "coprivnicar"

  if a.User == validUser && a.Pass == validPass {
    randNum := rand.Intn(1000000000)
    givenAuth.Code = randNum
    return randNum
  }
  givenAuth.Code = 8127349182734918
  return 0
}

func RepoAuthTestTryId(u userAuth) bool {
  // fmt.Println(givenAuth.Code)
  if u.Usercode == givenAuth.Code {
    return true;
  }
  return false
}
