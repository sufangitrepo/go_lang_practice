
package main


import (
	"fmt"
)



func main(){


	//declaring list
	// var collection[3] int ;
	// collection = [3] int{1,2,4}

	var varName int = 22
	fmt.Print(varName)
	 
	

	/*
	* Sum of the number is list 
    *
	*
	*
	*
	*/

	var listOfEnrollment  = [...] int {
		1,
		2,
		3,
		4,
		5,
	}
	var sum int;
	for i:=0; i< len(listOfEnrollment) ; i++{
		sum = listOfEnrollment[i] + sum
	} 	
	fmt.Println(sum)
	
    var list [5] any;
	list = [5] any{"",12,3,2,3}

	for i:=0 ; i< len(list); i++{
      list[i] = i
	} 
	fmt.Print(list)


}




func my(a... any){ 
}