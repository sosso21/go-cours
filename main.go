package main
import "fmt"




func main() {
	name := "Go Developers"
    fmt.Println("starting..." )
    fmt.Printf("hello%s" , name)
    fmt.Printf("\nvalue:%#v" , name)
	
	// ?  refs 
	
	name2 :=  &name
	*name2   ="Sofiane"
	fmt.Printf("\nyour name now is ... %#v , %#v" , name, name2 )
	
	// ? tables copy
	t1 := [3]int{1,2,3}
	t2  := t1 
	t2[1] =4
  fmt.Printf("\n table 1 & table 2 %#v %#v" ,t1 ,t2) 
  t3 :=  t1[1:] // start from 1 element
  fmt.Printf(" \n t3 %#v" , t3)
  
  t4 := []int{68, 69, 70}
  t5 := append(t4, 1,)
  
   fmt.Printf("\n   table 4 %#v %#v : length : %#v" ,t4 ,t5 ,len(t5) ) 
   
   // ?  les nil 
   
   var n string
   fmt.Printf(" \n %#v" , n)
   
   // ? objet map 
   
   obj := map[string]string{
	   "title": "hello wold",
	   "content" : "how are u",
	}
	objContent := obj["title"]
	fmt.Println("the content is :" , objContent )
	
	
	// ? objet struct 
	
	myTodo := Todo{ id : 1 , title : "buy milk" , description: "candia viva0" }
	setToggle( &myTodo)
	fmt.Printf("witch milk i should buy ? %#v  done :%#v" , myTodo.description , myTodo.done )
 
}


type Todo struct {
	id int
	title  string 
	description string 
	done  bool 
	createdAt int 
}

func (t   *Todo ) ToggleDoneTodo() {
	t.done = !t.done
}
 type CanToggle interface {
	ToggleDoneTodo()
 }
 func setToggle(t CanToggle) {
t.ToggleDoneTodo()
 }