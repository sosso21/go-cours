package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
  "math/rand"
    "time"
    
)

func main() {


  rand.Seed(time.Now().UnixNano())
  min := 10
  max := 30
  random_number := rand.Intn(max - min + 1) + min 


  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Simple Guess Shell")
  fmt.Println("---------------------")

  for {
    fmt.Print("-> ")
    text, _ := reader.ReadString('\n')
    // convert CRLF to LF
    text = strings.Replace(text, "\n", "", -1)
    text = strings.Replace(text, "\r", "", -1)

    if strings.Compare("hi", text) == 0 {
      fmt.Println("hello, Yourself")
      break
    }else {
      
		fmt.Printf("Say Hi \n")
    // continue

	}
  }


  
  for {
    fmt.Print("/number/~/ -> ")
    str_Number, _ := reader.ReadString('\n')

    // convert CRLF to LF
    str_Number = strings.Replace(str_Number, "\n", "", -1)
    str_Number = strings.Replace(str_Number, "\r", "", -1)


    
    
    number, err := strconv.Atoi(str_Number)

	if err != nil {
		fmt.Println("Error during conversion")
		 continue
	}

  if number == random_number {
		fmt.Println("You guess")
    break
  }  else if number < random_number {
		fmt.Println("More")
     
  }  else  {
		fmt.Println("less")
     
  }


    
  }

  fmt.Println("========================== \n  GOOD Bye !")

}
