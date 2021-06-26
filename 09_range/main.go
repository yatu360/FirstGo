package main

import (
	"fmt"

)

func main(){
	ids := []int{33,55,77,22,99,22}

	for _,id := range ids{
		fmt.Println(id)
	}
}