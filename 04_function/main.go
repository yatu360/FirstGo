package main

import (
	"fmt"
)


func greeting(name string) string{
	return "Hello" + name
}


func main() {

	fruitslice := []string{"Apple", "Orange"}
	fmt.Println(fruitslice)

	fruitslice = append(fruitslice, "banana")
	fmt.Println(fruitslice)

	

}
