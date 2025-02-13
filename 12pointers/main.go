package main

import (
    "fmt"
)

func main(){
	var p *int = new(int)
	var i int
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", *p) // prints: 0
	fmt.Printf("%v\n", p) // prints: address
	*p = 10
	fmt.Printf("%v\n", *p) // prints: 10
	fmt.Printf("%v\n", p) // prints: address

	p = &i
	fmt.Printf("%v\n", i) // prints: 0
	fmt.Printf("%v\n", &i) // prints: 0
	*p = 5
	fmt.Printf("%v\n", i) // prints: 5
	fmt.Printf("%v\n", &i) // prints: 0
}