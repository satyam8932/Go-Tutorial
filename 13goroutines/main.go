package main

import (
	"fmt"
	"sync" // comment if you don't want to use goroutines
	"time"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"John", "Jane", "Bob", "Alice", "Eve"}
var results = []string{}

func main (){
	t0 := time.Now()
	for i := 0; i<len(dbData); i++ {
		wg.Add(1) // comment if you don't want to use goroutines
		go dbCall(i)
	}
	wg.Wait() // comment if you don't want to use goroutines
	fmt.Printf("nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are: %v\n", results)
}
	
func dbCall(i int) {
	// Simulate DB call delay
	var delay float32 = 2000
	time. Sleep(time. Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done()  // comment if you don't want to use goroutines
}