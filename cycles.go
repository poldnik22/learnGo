package main

import "fmt"

func main() {
	//Цикл for
	/*
		for i := 0; i < 5; i++ {
			fmt.Printf("Hello %d\n", i)
		}
	*/

	//пример работы break
	/*
		for i := 1; i <= 100; i++ {
			if i == 50 {
				break
			}
			fmt.Println(i)
		}

		fmt.Println("After loop...")
	*/

	//цикл Range
	nums := []int{1, 2, 3, 4, 5}

	for index, element := range nums {
		fmt.Printf("Index: %d element: %d\n", index, element)
	}

	for _, element := range nums {
		fmt.Printf("Element: %d\n", element)
	}
}
