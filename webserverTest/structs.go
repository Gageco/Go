package main

type Timer struct {
  Length      int           `json:"length"`
  Id          int           `json:"id"`
  Finished   bool           `json:"finished"`
}

type Timers []Timer

type AlterStruct struct {
  Id            int       `json:"id"`
  State         bool      `json:"state"`
  Description   string    `json:"description"`
}

type AlterStructs []AlterStruct
