package main

import (
	"fmt"
	"strconv"
	"strings"

	"task2/progressive_tax"
)

func main() {

	var value string
	var currency string
	fmt.Print("Input value to calculate from: ")
	fmt.Scanf("%s %s", &value, &currency)

	input := strings.Fields(value)

	income, err := strconv.ParseFloat(input[0], 64)
	if err != nil {
		fmt.Println(err)
	}

	taxBrackets := []progressivetax.TaxBracket{
		{
			IncomeLowerBound: 0,
			TaxRate:          0,
		},
		{
			IncomeLowerBound: 1000,
			TaxRate:          0.1,
		},
		{
			IncomeLowerBound: 5000,
			TaxRate:          0.2,
		},
		{
			IncomeLowerBound: 10000,
			TaxRate:          0.3,
		},
	}

	result, err := progressivetax.CalculateTax(income, taxBrackets)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("%v %s", result, input[1]))
}
