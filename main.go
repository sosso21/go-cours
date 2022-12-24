package main

import (
  "fmt"
)



func main() {
  
hello := func () (string) {
  return  "hello"
}
fmt.Printf("%v \n", hello() )

  months := map[string]int{
    "jan": 31 ,
    "feb": 28,
    "mar": 31,
    "apr":30,
    "may": 31,
    "jun":30,
    "jul": 31,
    "aug":31,
    "sep":30,
    "oct": 31,
    "nov":30,
    "dev": 31,
  }

  for  month, days := range months {
    fmt.Printf("the month %s have %v days \n",  month , days )
  }

}
