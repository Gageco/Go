package main

import "fmt"
import "math"

func main() {
  //declare its a variable, name of variable, type
  //var integerName uint8
  //var strName string
  //var doubleName float64

  //for integers you can have uint8 - uint32 for different lengths of numbers

  fmt.Println(math.Pow(4, 2)) //4^2
  fmt.Println(math.Pow(4, .5)) //sqrt(4)

  fmt.Println("1 + 1 =", 1+1) // addition
  fmt.Println("length of \"hello_world\": ", len("hello_world"), " chars")
  fmt.Println(true && true) //equals true
  fmt.Println(false || true) //equals true

  //variables
  var intName uint8 //8 bit integer
  var doubleName float64 //float type
  var str string // type necessary because it not decleared straight up
  str = "test string"
  str1 := "test string 2" // type decelaration not necessary because it is declared straight up
  fmt.Println(str)
  fmt.Println(str1)

  //constant strings or variables
  const const_str string = `this is permanant` //this is the little tild ` on the top left of the keyboard it probably has some special meaning but i dont know what it is
  fmt.Println(const_str)

  //multiline stuff
  fmt.Println(`1
    2
3
    4
    5`) //formated just as it looks
}
