package main

import (
	"fmt"
	"sync"
)

func printTheString(s string, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Println(s)
}

func main() {

	var wg sync.WaitGroup

	sliceOfWords := []string{"string1", "string2", "string3", "string4", "string5", "string"}



	for index, value := range sliceOfWords {
		wg.Add(1)	
		go printTheString(fmt.Sprintf("Index : %d and Value : %s", index, value), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printTheString("Program is end there" , &wg)
}
