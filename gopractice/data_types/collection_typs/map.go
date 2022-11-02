package main 

import (

	"fmt"
)

func main(){

	/*
	Declaring  map 
	*/

	var mapOfData map[string]any 

	//this make function will be use to create map object it is best technique

	mapOfData = make(map[string]any)

    mapOfData["key1"] = "123"
    
	//fmt.Print(mapOfData["key1"])

	for key, value := range mapOfData{
		fmt.Print(key,value)
	}

	
}