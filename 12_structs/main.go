package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

func main() {

	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "London", age: 25, gender: "f"}

	person2 := Person{"Samantha", "Smith", "London", "f", 24}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(greet1(person2))
	fmt.Println(person2.greet())
	person2.hasbirthday()
	fmt.Println(person2.greet())
	person2.getMarried("Williams")
	fmt.Println(person2.greet())

}

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " " + strconv.Itoa(p.age)

}

func greet1(p Person) string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " " + strconv.Itoa(p.age)
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "f" {
		p.lastName = spouseLastName
	}

}

func (p *Person) hasbirthday() {
	p.age++

}
