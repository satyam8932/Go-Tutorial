package main

import "fmt"
import "bufio"
import "os"

func main() {
	welcome := "Welcome to the user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the pizza!")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
}
