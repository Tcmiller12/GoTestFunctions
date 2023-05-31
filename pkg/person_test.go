package pkg_test

import ("testing"
    "github.com/temiller/person/person")

func Calculate(x int) int {
	return x + 3
}

// func TestCalculate(t *testing.T) {
// 	if Calculate(2) != 4 {
// 		t.Error("Expected 4")
// 	}
// }




func TestComputeAge_Over25(t *testing.T) {
	description, over25, err := person.ComputeAge(40, "Bryan") 
	if !over25 {
		t.Error("Expected true but got false for over 25")
	}

	if description != "Bryan is 40 years old"{
		t.Error ("Expected description Bryan is 40 years old")
	}

	if err != nil {
		t.Errorf("Expected no error but got error %v", err)
	}

}

func TestComputeAge_Under25(t *testing.T) {
	description, over25, err := person.ComputeAge(20, "Bryan") 
	if over25 {
		t.Error("Expected false but got true for over 25")
	}

	if description != "Bryan is 20 years old"{
		t.Error ("Expected description Bryan is 20 years old")
	}

	if err != nil {
		t.Errorf("Expected no error but got error %v", err)
	}

}

func TestComputeAge_Equals25(t *testing.T) {
	description, over25, err := person.ComputeAge(25, "Bryan") 
	if over25 {
		t.Error("Expected false but got true for over 25")
	}

	if description != "Bryan is 25 years old"{
		t.Error ("Expected description Bryan is 25 years old")
	}

	if err != nil {
		t.Errorf("Expected no error but got error %v", err)
	}

}

func TestComputeAgeCollins_Over25(t *testing.T) {
	description, over25, err := person.ComputeAge(40, "Collins") 
	if !over25 {
		t.Error("Expected true but got false for over 25")
	}

	if description != "Collins is 40 years old"{
		t.Error ("Expected description Collins is 40 years old")
	}

	if err != nil {
		t.Errorf("Expected no error but got error %v", err)
	}

}

func TestComputeAgeCollins_Under25(t *testing.T) {
	description, over25, err := person.ComputeAge(20, "Collins") 
	if over25 {
		t.Error("Expected false but got true for over 25")
	}

	if description != "Collins is 20 years old"{
		t.Error ("Expected description Collins is 20 years old")
	}

	if err != nil {
		t.Errorf("Expected no error but got error %v", err)
	}

}

func TestComputeAgeCollins_Equals25(t *testing.T) {
	description, over25, err := person.ComputeAge(25, "Collins") 
	if over25 {
		t.Error("Expected false but got true for over 25")
	}

	if description != "Collins is 25 years old"{
		t.Error ("Expected description Collins is 25 years old")
	}

	if err != nil {
		t.Errorf("Expected no error but got error %v", err)
	}

}


func TestComputeBobAge_Over25(t *testing.T) {
	description, over25, err := person.ComputeAge(50, "Bob") 
	if !over25 {
		t.Error("Expected true but got false for over 25")
	}

	if description != "Bob is 50 years old"{
		t.Error ("Expected description Collins is 40 years old")
	}

	if err != nil {
		t.Errorf("Expected no error but got error %v", err)
	}

}

func TestComputeBobAge_Under25(t *testing.T) {
	description, over25, err := person.ComputeAge(20, "Bob") 
	if over25 {
		t.Error("Expected false but got true for over 25")
	}

	if description != "Bob is 20 years old"{
		t.Error ("Expected description Collins is 20 years old")
	}

	if err != nil {
		t.Errorf("Expected no error but got error %v", err)
	}

}

func TestComputeBobAge_Equals25(t *testing.T) {
	description, over25, err := person.ComputeAge(25, "Bob") 
	if over25 {
		t.Error("Expected false but got true for over 25")
	}

	if description != "Bob is 25 years old"{
		t.Error ("Expected description Collins is 25 years old")
	}

	if err != nil {
		t.Errorf("Expected no error but got error %v", err)
	}

}

// replace previous 3 names & age for Collins & test
