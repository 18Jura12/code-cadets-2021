package main

import (
	"code-cadets-2021/homework_1/task_2/progressive_tax"
	"fmt"
)

func main() {

	var value string
	var currency string
	fmt.Print("Input value to calculate from: ")
	fmt.Scanf("%s %s", &value, &currency)

	result, err := progressive_tax.CalculateTax(value + " " + currency)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
