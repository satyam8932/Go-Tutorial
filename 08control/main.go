package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("Hello World!")

	num1 := 10
	num2 := 0
	var result, remainder, err = intDivision(num1, num2)
	if err != nil {
        fmt.Println(err)
    }
	fmt.Printf("The result of the integer division is %v with remainder %v \n", result, remainder);
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error;

	if denominator == 0 {
        err = errors.New("division by zero is not allowed")
		return 0, 0, err;
    }
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, nil;
}