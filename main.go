//Nicholas Larsen
//4/25/2020
//asks for 2 number and outputs the before and after the swapping of the 2 numbers
package main

import "fmt"

func integers() {
  var num1 int
  var num2 int
  fmt.Println("enter a single digit number")
  fmt.Scanln(&num1)
  fmt.Println("enter another single digit number")
  fmt.Scanln(&num2)
  //gets the 2 numbers
  fmt.Println("before the swap, The first number is",num1,"and the second is",num2)
  //before the swap
  fmt.Println("after the swap, The first number is",num2,"and the second is",num1)
  //after the swap
}

func main() {
  integers()
}