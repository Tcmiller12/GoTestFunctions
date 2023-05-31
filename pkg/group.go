package pkg

import (
	"fmt"
)

type Group struct {
	individuals []Person
	IdGenerator int
}

func (p *Group) AddPerson(person Person) {
	p.IdGenerator++
	person.id = p.IdGenerator
	p.individuals = append(p.individuals, person)
}

func (p *Group) Size() int {
	return len(p.individuals)
}

func (p *Group) GetPersonWithLastName(lastName string) (Person, error) {
	for _, person := range p.individuals {
		if person.LastName == lastName {
			return person, nil
		}
	}
	return Person{}, fmt.Errorf("does not exist")
}

func (p *Group) GetFirstName(id int) (string, error) {
	fmt.Println("Hello World")
	fmt.Println(len(p.individuals))
	for _, person := range p.individuals {
		fmt.Printf("ID: %v Firstname:%v LastName:%v \n", person.id, person.FirstName, person.LastName,)
		if id == person.id {
			return person.FirstName, nil
		}
	}
	return "", fmt.Errorf("Can't find the person with id: %d",id)
}

func (p *Group) GetFullNameByIDAndAge(id int, age int) (string, error) {
	for _, person := range p.individuals {
		if id == person.id && age == person.Age {
			return fmt.Sprintf("%s %s", person.FirstName, person.LastName), nil
		}
	}
	return "", fmt.Errorf("Can't find the person with Age or Id")
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
	id        int
}
