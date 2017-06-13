package main

import "time"

type Todo struct {
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
    Id        int       `json:"id"`
}

type Todos []Todo
