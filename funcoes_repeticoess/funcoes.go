package main

import "fmt"

func main () {
	// For loop
	for i := 0; i< 10; i++{
		fmt.Println(i)
	}

	// While
	teste := 0
	for teste <= 10 {
		fmt.Println("MENOR QUE 10")
		teste++
		
	}
	// Do While
	anExpression := false
	for ok := true; ok; ok = anExpression {
		fmt.Println("OOI")
	}

	// For com Range
	temp  := []string{"temp1", "temp2", "temp3"}
	for i, value := range temp {
		println(i, value)
	}
	// Ignorando Indice
	for _, value := range temp{
		println(value)
	}
}