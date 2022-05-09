package main

import "fmt"

type Person struct {
	Id        int
	Name      string
	Address   // fields of address will embed in Person
	FirstName string
	Address2  Address // just structure property
}

type Address struct {
	Country string
	City    string
	Home    string
}

func main() {
	checkThisOut()
}

func checkThisOut() {
	person := Person{
		Name:      "Vichka",
		FirstName: "Brichka",
		Address: Address{
			Country: "Russia",
			City:    "Tver",
		},
		Address2: Address{Country: "Kykraina"},
	}

	fmt.Println(person)
	//Embedded vs Property
	fmt.Println(person.Country)
	fmt.Println(person.Address2.Country)

	person.updateName("Sergey")
	fmt.Println(person) // nothing will happen
	person.setName("Ura")
	fmt.Println(person)
	person.Speak()

	objectCast(&person)
	objectCast(&Animal{})
}

// This function useless , because we must do update on object point, but in this case we do on object copy
func (p Person) updateName(name string) {
	// doesn't work any way (&p).Name = name
	p.Name = name
}

// this how it should be
func (p *Person) setName(name string) {
	p.Name = name
}

type Speakable interface {
	Speak()
	Cry()
}

type Animal struct {
	Name string
}

func (p *Person) Speak() {
	fmt.Println("Speaking", p, "Hello")
}
func (p *Person) Cry() {
	fmt.Println(p, "is crying")
}

func (p *Animal) Speak() {
	fmt.Println("I ", p, "watching anime")
}
func (p *Animal) Cry() {
	fmt.Println(p, "is crying")
}

func objectCast(s Speakable) {

	switch s.(type) {
	case *Animal:
		s.Speak()

	case *Person:

		chelovek, ok := s.(*Person) // cast
		if ok {
			fmt.Println("i am ", chelovek.Name)
		}
		s.Speak()

	default:
		fmt.Println("Unknown type")
	}

}
