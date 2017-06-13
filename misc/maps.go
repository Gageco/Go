package main

import (
  "fmt"
)
func main() {

  var m map[string]int                                                          //after you declare the map you have to make it
  m = make(map[string]int)                                                      //how you make the map

  m["test"] = 5                                                                 //assigning a variable to the map

  fmt.Println(m["test"])
}
