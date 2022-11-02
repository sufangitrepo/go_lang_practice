package main

import (
	"fmt"
	"reflect"
	
)

func main() {
  
/*
There is just three data types in go language those data types are 
   
  string    (just store text data) 
  numeric    
            int   (8,16,32)
			float same 
  bolean
           jsut store true or false means boolean value

    
*/	



 //1 Strict typing 
  var name string = "sufyan"
  var rollNo int = 123
  var isMale bool = true
  fmt.Print(isMale, name, rollNo)

  //2 it is like a type inference no define type but adopt type 
  //acording the data that is stored in it either 1. using just var/ := this 

  var name2 = "sufi"
  rollNo2 := 123
  isMale2 := true
  fmt.Print(name2, rollNo2, isMale2)
  fmt.Print(reflect.TypeOf(name2))
  


}



func sum(a int, b int) (int, string) {

	return 23, "hello"
}