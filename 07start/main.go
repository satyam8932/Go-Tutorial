package main

import "fmt"

func main() {
	// var num1 int = 10
    // var num2 int = 5

    // sum := num1 + num2
    // fmt.Println("Sum:", sum)

    // diff := num1 - num2
    // fmt.Println("Difference:", diff)

    // product := num1 * num2
    // fmt.Println("Product:", product)

    // quotient := float64(num1) / float64(num2)
    // fmt.Println("Quotient:", quotient)

    // remainder := num1 % num2
    // fmt.Println("Remainder:", remainder)

	var n int = 0;
	for i:=0; i<1000000; i++ {
		n = n + 1;
	}
	fmt.Println(n)
}