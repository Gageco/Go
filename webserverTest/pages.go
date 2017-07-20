package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "io"
  "io/ioutil"
  "math/rand"
  //"github.com/gorilla/mux"
)

var timerID int
var alterID int

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Welcome!")
}

func TimerCreate(w http.ResponseWriter, r *http.Request) {
  var timer Timer

  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))                  //limit the requests that can come in
  checkErr(err)

  err = r.Body.Close()
  checkErr(err)

  if err := json.Unmarshal(body, &timer); err != nil {                          // <- this is a pretty neat layout, its a line that dones something then checks for error, good to remember for future projects
    panic(err)
  }
  // w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  // w.WriteHeader(422)
  //
  //   err := json.NewEncoder(w).Encode(err)
  //   if err != nil {
  //     panic(err)
  //   }
  // }
  timerID +=1
  timer.Id = timerID
  go RepoCreateTimer(timer, w)

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  err = json.NewEncoder(w).Encode(timer)
  checkErr(err)
}

func TimerIndex(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(422)
  err := json.NewEncoder(w).Encode(timers)
  checkErr(err)
}

func RandomText(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  randNum := rand.Intn(10)
  randText := [10]string {"hello", "pirate", "watermelon", "1612", "heart", "ford focus", "peice", "peace", "keyboard", "coffee"}

  err := json.NewEncoder(w).Encode(randText[randNum])
  if err != nil {
    panic(err)
  }
}

func StateChange(w http.ResponseWriter, r *http.Request) {
  var alter AlterStruct

  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
  checkErr(err)

  err = r.Body.Close()
  checkErr(err)

  err = json.Unmarshal(body, &alter)
  checkErr(err)

  // alterID += 1
  // alter.Id = alterID
  alter = RepoStateChange(alter)

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  err = json.NewEncoder(w).Encode(alter)
  checkErr(err)
}

func StateChangeIndex(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(422)
  err := json.NewEncoder(w).Encode(alters)
  checkErr(err)
}
