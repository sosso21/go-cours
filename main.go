package main

import (
  "fmt"
  "errors"
  
)
 //        [T float64 | int64] 
 // ?   ~ : can accept float32 

func Divide [T ~float64 | ~int64](a , b T) (T , error) {
  if b == 0 {
    return 0.0 , errors.New("Invalid Name")
  }
  return  a/b  ,nil
}

func main() {
  // result , err := Divide(22.00, 7.00)
  result , err := Divide[float64](22.00, 7.00)
   if  err != nil {
    panic(err)
   }

   fmt.Printf("the result is: %v",result)
}