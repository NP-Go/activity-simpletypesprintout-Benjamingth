package main

import "fmt"

//Insert variables declarations here based on activity

func tellMeTypes() {
	//insert code here to print out types of variables
	age := 20
	weight := 73.0
	firstname := "Luke"
	lastname := "Skywalker"

	fmt.Printf("%T %T %T %T\n", age, weight, firstname, lastname)
}

func main() {
	tellMeTypes()
}
