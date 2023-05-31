package pkg_test

import (
	"testing"

	"github.com/temiller/person/pkg"
)

func TestAddPerson(t *testing.T) {
	group := GetGroup()

	size := group.Size()

	if size != 6 {
		t.Errorf(" Error Expected size of 1 got %v", size)
	}
	actualPerson, _ := group.GetPersonWithLastName("Deer")
	if "John" != actualPerson.FirstName {
		t.Error("Expected first name is not equal to actual first name")
	}

}

func TestGetFullNameByAgeAndID(t *testing.T) {
	group := GetGroup()
	fullName, err := group.GetFullNameByIDAndAge(1, 31)
	if err != nil {
		t.Error("Failed to find name")
	}
	if fullName != "John Deer" {
		t.Error("Fullname does not match expected full name")
	}
}

func TestSize(t *testing.T) {
	// persons := []string{"John", "Meg", "Peter"}
	group := GetGroup()
	if group.Size() != 6 {
		t.Error("Expected 3 names, did not return 3")
	}
}

func TestGetFirstName(t *testing.T) {

	group := GetGroup()

	firstName, err := group.GetFirstName(1)

	if firstName != "John" {
		t.Errorf("Expected the name John got %s", firstName)
	}

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

}

// GetGroup creates a group and adds people to it
func GetGroup() pkg.Group {
	group := pkg.Group{}
	persons := []pkg.Person{
		pkg.Person{
			FirstName: "John",
			LastName:  "Deer",
			Age:       31,
		},
		{
			FirstName: "Meg",
			LastName:  "Deer",
			Age:       30,
		},
		{
			FirstName: "Peter",
			LastName:  "Deer",
			Age:       35,
		},
		{
			FirstName: "Chris",
			LastName:  "Deer",
			Age:       20,
		},
		{
			FirstName: "Tom",
			LastName:  "Deer",
			Age:       50,
		},
		{FirstName: "Joe",
			LastName: "Deer",
			Age:      33,
		},
	}
	for _, person := range persons {
		group.AddPerson(person)
	}
	return group
}
