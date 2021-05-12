package progressive_tax

import (
	"fmt"
	"github.com/pkg/errors"
	"math"
	"strconv"
	"strings"
)

func CalculateTax(value string) (string, error) {

	taxRates := []TaxRate{
		{
			incomeLowerBound: 0,
			incomeUpperBound: 1000,
			taxBracket:       0,
		},
		{
			incomeLowerBound: 1000,
			incomeUpperBound: 5000,
			taxBracket:       0.1,
		},
		{
			incomeLowerBound: 5000,
			incomeUpperBound: 10000,
			taxBracket:       0.2,
		},
		{
			incomeLowerBound: 10000,
			incomeUpperBound: math.Inf(1),
			taxBracket:       0.3,
		},
	}

	if !checkTaxClasses(taxRates) {
		return "", errors.New("Wrong tax rates!")
	}

	input := strings.Fields(value)

	if len(input) != 2 {
		fmt.Println(len(input))
		return "", errors.New("Wrong value input format!")
	}

	income, err := strconv.Atoi(input[0])
	if err != nil {
		return "", err
	}
	if income < 0 {
		return "", errors.New("Income cannot be negative!")
	}

	result := 0.0
	for _, taxRate := range taxRates {
		result += (math.Min(taxRate.incomeUpperBound, float64(income)) - taxRate.incomeLowerBound) * taxRate.taxBracket
		if taxRate.incomeUpperBound > float64(income) { break }
	}

	return fmt.Sprintf("%d %s", int(result), input[1]), nil

}

func checkTaxClasses(taxRates []TaxRate) bool {

	for index, taxRate := range taxRates {
		if index != 0 && taxRate.incomeLowerBound != taxRates[index - 1].incomeUpperBound {
			return false
		}
		if index != len(taxRates) - 1 && taxRate.incomeUpperBound != taxRates[index + 1].incomeLowerBound {
			return false
		}
	}
	return true

}
