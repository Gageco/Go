package main

import (
  "log"
  "net/http"
  "fmt"
)

func main() {
  fmt.Println("Webserver Has Started at localhost:8080")
  router := NewRouter()

  log.Fatal(http.ListenAndServe(":8080", router))
}
