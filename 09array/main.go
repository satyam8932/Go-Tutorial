package main

import "fmt"

func main() {
	var intArr [3]int32;
	fmt.Println(intArr[0]) // Index prints 0
	fmt.Println(intArr[1:3]) // Slice of array

	// Address
	fmt.Println(&intArr[0]);
	fmt.Println(&intArr[1]);
	fmt.Println(&intArr[2]);

	arr := []int32{2,5,2};
	fmt.Println(arr); // Array prints [2 5 2]
	fmt.Printf("lenght is %v with capacity %v\n", len(arr), cap(arr))
	arr = append(arr, 7) // Corrected append usage
	fmt.Println(arr) // Array prints [2 5 2 7]
	fmt.Printf("lenght is %v with capacity %v\n", len(arr), cap(arr))

	arr2 := make([]int32, 4, 5)  // map use to create a slice with initial length and capacity
	fmt.Println(arr2); // Array prints [0 0 0 0]

}