package main

import "fmt"

func main() {

	//демонстрация работы ключегого слова fallthroug. По умолчанию Case заканчивает работать при первом совпадении,
	//но если мы хотим продолжить сравнение со следующими вариантами, то мы можем добавить в условии ключевое слово
	//fallthrough, как показано ниже.
	number := 10

	switch {
	case number > 1:
		fmt.Println("Number is greater than 1")
		fallthrough
	case number < 11:
		fmt.Println("Number < 11")
	}

	/*
		name := "John"
		//Функция switch используется, когда есть много разных вариантов значений
		switch name {
		case "Jordan":
			fmt.Println("Hello Jordan")
		case "Kate":
			fmt.Println("Hello Kate")
		case "John":
			fmt.Println("Hello John")

		default:
			fmt.Println("I dont't know you!")
		}
	*/
}
