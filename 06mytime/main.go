package main

import "fmt"
import "time"

func main() {
	fmt.Println("Time Module");

	presentTime := time.Now();
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"));

}