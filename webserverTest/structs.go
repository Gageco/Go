package main

type Timer struct {
  Length      int           `json:"length"`
  Id          int           `json:"id"`
  Completed   bool          `json"completed"`
}

type Timers []Timer
