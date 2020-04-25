package main

import "fmt"

func integers() {
  var num1 int
  var num2 int
  fmt.Println("enter a single digit number")
  fmt.Scanln(&num1)
  fmt.Println("enter another single digit number")
  fmt.Scanln(&num2)
  fmt.Println("before the swap, The first number is",num1,"and the second is",num2)
  fmt.Println("after the swap, The first number is",num2,"and the second is",num1)
}

func main() {
  integers()
}