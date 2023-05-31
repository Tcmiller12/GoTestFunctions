package pkg

import "fmt"

func ComputeAge(age int, name string) (string, bool, error) {
	description := fmt.Sprintf("%v is %d years old", name, age)
	//anything comparing returns boolean
	over25 := age > 25
	return description, over25, nil
}