package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 25.64578

	result := math.Ceil(a)        //делает округление к большему значению
	b := math.Floor(a)            //округленеи к меньшему значению
	var c float64 = math.Round(a) //округление по классическим правилам математики

	fmt.Println(result)
	fmt.Println(b)
	fmt.Println(c)

	/*
		var a float64 = 144
		result := math.Sqrt(a)
		fmt.Println(result)
	*/
}
