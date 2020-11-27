package main

import "fmt"

func main(){
  var name string

  fmt.Print("What's your name? ")
  fmt.Scanln(&name)
  fmt.Println("Hello ", name)
}
