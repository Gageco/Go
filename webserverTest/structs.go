package main

type Timer struct {
  Length      int           `json:"length"`
  Id          int           `json:"id"`
  Finished   bool           `json:"finished"`
}

type Timers []Timer
