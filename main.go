package main

import (
	"fmt"
	"github.com/temiller/person/person"
)

func Calculate(x int) (result int){
result = x + 2
return result
}

func main() {
	// fmt.Println("Go Testing Tutorial")
	// result := Calculate(2)
	// fmt.Println(result) 
	// Use underscore to ignore a variable, ex: computeAge returns 3 but we only want D
	d, _, _ := person.ComputeAge(40, "Bryan") 
	fmt.Println(d) 
	d, _, _ = person.ComputeAge(19, "Collins")
	fmt.Println(d)
	d, _, _ = person.ComputeAge(55, "Bob")
	fmt.Println(d)
}
