package main

import (
  "fmt"
  "errors"
  
)

func Divide(a , b float32) (float32 , error) {
  if b == 0 {
    return 0.0 , errors.New("Invalid Name")
  }
  return  a/b  ,nil
}

func main() {
   result , err := Divide(22.00, 7.00)
   if  err != nil {
    panic(err)
   }

   fmt.Printf("the result is: %v",result)
}