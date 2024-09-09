package main

import "fmt"

func main() {
	num := 3
	if num > 0 {
		fmt.Println("Number is greater then 0")
	} else if num < 0 {
		fmt.Println("Number < 0")
	} else if num == 3 {
		fmt.Println("Number equals 3!!")
	}

	fmt.Scan(&num)

}
