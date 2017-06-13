package main

import "fmt"

func main() {
  var i int

  defer fmt.Println("two")
	defer fmt.Println("world") // The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
  defer fmt.Println("one")
	fmt.Println("hello ")

  for i < 6 {
    defer fmt.Println(i)
    i+=1
  }
  fmt.Println("Done")

defer fmt.Println(".")
defer fmt.Printf("everyone")
defer fmt.Printf("there ")
defer fmt.Printf("hello ")

}
// ok so defer statements are super new to me this is interesting, its a good way to run things backwards easily so the logic may take a while to get but hopefully you will be able to get it down. if not look here for some help (https://tour.golang.org/flowcontrol/13)

// but yeah so like whats happening in this section is that it kicks out 'hello' because it is before the 'for' statment and then it kicks out 'Done' because it is before the for statement evaluates then it goes back and runs all the other things. pretty cool

//something to check is if the logic is backwards as well in that it defines variables backwards or what?
