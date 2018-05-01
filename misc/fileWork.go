package main

import (
    "bufio"
    "fmt"
    //"log"
    "os"
)

func main() {
  //READ FROM FILE
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

  //APPEND TO FILE
  f, err := os.OpenFile(csvLocation, os.O_APPEND|os.O_WRONLY, 0644)
  csvString := pulledData.Date + "," + pulledData.Temperature + "," + pulledData.Humidity + "\n"
  n, err = f.WriteString(csvString)
  fmt.Println(n)
  if err != nil {
    fmt.Println(err)
  }

  f.Close()

  //OVERWRITE FILE
  csvHeaders := []byte("Date,Temperature,Humidity\n")
  err := ioutil.WriteFile(location, csvHeaders, 0644)
  if err != nil {
    fmt.Println(err)


}
