package main

import (
    "bufio"
    "fmt"
    //"log"
    "os"
)

func main() {
  inFile, _ := os.Open("./readThisFile.txt")
  defer inFile.Close()
  scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

  scanner.Scan()                  //line 1
  fmt.Println(scanner.Text())

  scanner.Scan()                  //line 2
  text := scanner.Text()
  fmt.Println(text)

  scanner.Scan()                  //line 3
  fmt.Println(scanner.Text())

  for scanner.Scan() {            //all lines, wont work cause reached end of file
    fmt.Println(scanner.Text())
  }
}
