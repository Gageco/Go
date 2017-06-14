package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "io"
  "io/ioutil"
  //"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Welcome!")
}

func TimerCreate(w http.ResponseWriter, r *http.Request) {
  var timer Timer
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
  if err != nil {
    panic(err)
  }
  if err := r.Body.Close(); err != nil {
    panic(err)
  }
  if err := json.Unmarshal(body, &timer); err != nil {
    panic(err)
  }
  // w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  // w.WriteHeader(422)

    // err := json.NewEncoder(w).Encode(err)
    // if err != nil {
    //   panic(err)
    // }
  // }

  go RepoCreateTimer(timer, w)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  // err = json.NewEncoder(w).Encode(t)
  // if err != nil {
  //   panic(err)
  // }
}

func TimerIndex(w http.ResponseWriter, r *http.Request) {
  // w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  // w.WriteHeader(http.StatusOK)
  err := json.NewEncoder(w).Encode(timers)
  if err != nil {
    panic(err)
  }
}
