package main

import "fmt"

func main() {
	/*
		a := 1
		if a > 0 {
			fmt.Println("a >0 ")
		}
	*/
	a := 3
	b := 10

	// && - И
	if a > 1 && b > 10 {
		fmt.Println("TRUE!1")
	}

	// || - или
	if a > 1 || b > 9 {
		fmt.Println("TRUE!2")
	}

	// != не
	if a != 1 {
		fmt.Println("DONE!")
	}

}
