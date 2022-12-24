package main

import (
	"errors"
	"fmt"
)

func SayHello(name string ) (string , error)  {
	if name== "" {
		return "" , errors.New("You should enter name")
		 
	}
	result := fmt.Sprintf("hello %v",name)
	  
	return result , nil
}