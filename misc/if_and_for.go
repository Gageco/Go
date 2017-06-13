package main
import "fmt"

func main() {

  numberone := 1

  if numberone == 1 { //so for somereason you have to have th else and else if statements directly after the brackets so just be aware
    fmt.Println("variable numberone is: ", numberone)
  } else if numberone == 2{
    fmt.Println("how did you even get into this loop??")
  } else {
    fmt.Println("variable numberone is not ", numberone)
  }



  //for (initial variable; how to kick out, iterate up)
  for i := 1; i <= 5; i+=1 {
    fmt.Println(i)
  }

//combo if and for
  for i := 1; i <= 10; i+=1 {
    if i % 2 == 0 {
      fmt.Println(i, " Even")
    } else {
      fmt.Println(i, " Odd")
    }
  }

//for statement used as a while statements
for numberone <= 5 {
  fmt.Println(numberone)
  numberone += 1
}

}
