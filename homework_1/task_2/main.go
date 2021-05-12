package main

import (
	"fmt"
	"task2/progressive_tax"
)

func main() {

	var value string
	var currency string
	fmt.Print("Input value to calculate from: ")
	fmt.Scanf("%s %s", &value, &currency)

	result, err := progressivetax.CalculateTax(value + " " + currency)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
