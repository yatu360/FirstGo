package main

import "fmt"

type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int

}




func main(){

	person1 := Person{firstName:"Samantha", lastName: "Smith", city:"London", age: 25, gender: "f"}

	person2 := Person{"Samantha", "Smith", "London", "f", 24}

	fmt.Println(person1)
	fmt.Println(person2)


}