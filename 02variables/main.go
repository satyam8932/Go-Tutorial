package main

import "fmt"

func main() {
	var username string = "John Doe";
	fmt.Println("Hello, " + username + "!");
	fmt.Printf("Variable of Type: %T \n", username);
	
	var isLoggedIn bool = true;
	fmt.Println("Is logged in: ", isLoggedIn);
	fmt.Printf("Variable of Type: %T \n", isLoggedIn);

	var smallVal int  = 255;
	fmt.Println("Is logged in: ", smallVal);
	fmt.Printf("Variable of Type: %T \n", smallVal);

	testNum := 10;
	fmt.Println("Test Number: ", testNum);

	const loginToken string = "sadfjklasdjflkjasdf";
	fmt.Printf("Variable of Type: %T \n", loginToken);

	testNum = 20;
}
