package main 

import (
	"testing"
  )

  var tests  = []struct {
	name   string 
	a       float32 
	b       float32
	want    float32 
	isErr   bool
  } {
	{"valid" ,10.00 , 2.00 , 5.00 , false},
	{"invalid" ,10.00 , 0.00, 0.00 , true},
	{"invalid2" ,10.00 , 2.00, 7.00 , false},
  }

  func TestDivide(t *testing.T ){

for _, tt := range tests {
	want , err := Divide(tt.a , tt.b)

	if (err != nil) != tt.isErr {
		 t.Errorf("%s : return %v , we need %v"  ,tt.name ,err ,tt.isErr  )
	}
	if   want != tt.want {
		
		t.Errorf("%s : return %v , we need %v"  ,tt.name ,want ,tt.want )
	}


}


  }

