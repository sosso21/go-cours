package main

import (
  "fmt"
  "math/rand"
 "time"
    
)





 

func randomNumberChanel(numChan chan  int)   {
  min := 20
  max := 100

  rand.Seed(time.Now().UnixNano())
  random := rand.Intn(max - min + 1) + min 


  numChan  <- random
}

func main() {
  number := make(chan int )
  defer  close(number) // defer == fo it at last
  go randomNumberChanel( number)
  num  := <- number
 
  fmt.Println( num)
}