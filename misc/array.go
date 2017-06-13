package main
import "fmt"

func main() {
  var myArray[5]int //not uint8 for some reason
  coinArray := [3]string {"eth", "btc", "ltc"}

  for i:=0; i < len(coinArray); i++ {
    fmt.Printf(coinArray[i], ".")
  }

  for i:= 0; i < 5; i++ {
    myArray[i] = i+5
  }
  for i:=0; i <5; i++ {
    fmt.Println(myArray[i])
  }

}
