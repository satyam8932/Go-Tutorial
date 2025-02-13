package main

import "fmt"

type gasEngine struct {
    kmpl uint8
	liters uint8
	owner
}

type owner struct {
    name string
    age  uint8
}


func main() {
	// Defined Struct
	var prevEngine gasEngine = gasEngine{25,65,owner{"John Doe",30}}
	fmt.Printf("Previous Engine Details: MPG=%d, Liters=%d, Owner=%s\n", prevEngine.kmpl, prevEngine.liters, prevEngine.owner.name)

	// Anonymous Struct ( but can't be used another place than where it's declared )
	var myEnigne = struct{
		mpg uint8
		gallons uint8
	}{34,65}
	fmt.Println(myEnigne)
}

func (e gasEngine) kmsLeft() uint8 {
	return e.kmpl * e.liters
}

func (o owner) fullName() string {
    return fmt.Sprintf("%s %d", o.name, o.age)
}

func test(s string) string {
	return fmt.Sprintf("Hello, %s!", s)
}

// Interfaces
type engine interface {
    kmsLeft() uint8
}

type ownerInterface interface {
    fullName() string
}