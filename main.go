package main

import (
  "fmt"
  "encoding/json"
)

   
// declaring a struct
type Human struct{
        
  // defining struct variables
  Name string  `json:"Name"`
  Address string  `json:"Address"`
  Age int  `json:"Age"`
}

 
func main() {
        
    // defining an array instance
    // of struct type
    var human []Human
        
    // JSON array to be decoded
    // to an array
    Data := []byte(`
    [
        {"Name": "sofiane", "Address": "Djborland", "Age": 26},
        {"Name": "mohMost", "Address": "HotVille", "Age": 24},
        {"Name": "Yanis", "Address": "MarseilleBébé", "Age": 25}
    ]`)
        
    // decoding JSON array to 
    // human2 array
    err := json.Unmarshal(Data, &human)
        
        if err != nil {
        
        // if error is not nil
        // print error
            fmt.Println(err)
        }
        
    // printing decoded array 
    // values one by one
    for  i := range human{
        
        fmt.Println( human[i])
    }
 

    Salfils := Human{
      Name: "Salfils",
      Address: "Dans sa chambre",
      Age: 24,
    } 
     mohProut := Human{
      Name: "Moh Prout",
    Address: "Mirabo" ,
    Age: 24,
  } 
 
  rkhwas := []Human{ Salfils,  mohProut}
 
    
  b, err := json.Marshal(rkhwas)
  if err != nil {
      fmt.Printf("Error: %s", err)
      return;
  }
  
  fmt.Println(string(b))



}