package main

import (
	"fmt"
	"rcfi/repository"
	"rcfi/service/delivery"
)

func main() {

	fmt.Println("Choose What You Want")
	fmt.Println("1. Sorting (Soal no.1)")
	fmt.Println("2. Client Server (Soal no.2)")
	fmt.Println("3. Fix Broken Thing (Soal no.3)")

	fmt.Print("Enter answer: ")
	var value string
	fmt.Scanln(&value)

	if value == "1" {
		inputvalue := []int{4, 9, 7, 5, 8, 9, 3}
		repository.Swap(inputvalue)
	} else if value == "2" {
		delivery.Router()
	} else if value == "3" {
		collection := []int{3, 5, 2, -4, 8, 11}
		fmt.Println(repository.Foo(7, collection))
	} else {
		fmt.Println("Watch Carefully")
	}
}
