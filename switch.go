package main

import "fmt"

func main() {
	name := "John"
	//Функция switch используется, когда есть много разных вариантов значений
	switch name {
	case "Jordan":
		fmt.Println("Hello Jordan")
	case "Kate":
		fmt.Println("Hello Kate")
	case "John":
		fmt.Println("Hello John")
	}

}
