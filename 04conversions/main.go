package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Conversions")
	fmt.Println("-------------")
	fmt.Println("Please rate the pizza between 1 and 5")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating the pizza as", input)
	
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Rating as number is", numRating + 1)
	}
}