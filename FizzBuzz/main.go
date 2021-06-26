package main

import "fmt"

//FizzBuzz n = 100
// Multiple of 3 output Fizz
//Multiple of 5 output Buzz
//Multiple of 3 and 5, FizzBuzz
func main(){

	for i:=0; i<=100; i++{
		var a string = ""
		var b string = ""
		if i%3 == 0 {
			a="Fizz"
		}
		if i%5 == 0{
			b="Buzz"
			
		}
		
		fmt.Println("number ",i,": ", a , b )

	}
}