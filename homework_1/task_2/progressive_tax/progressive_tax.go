package progressivetax

import (
	"github.com/pkg/errors"
	"math"
)

func CalculateTax(income float64, taxBrackets []TaxBracket) (float64, error) {

	if !checkTaxClasses(taxBrackets) {
		return 0, errors.New("wrong tax rates")
	}

	if income < 0 {
		return 0, errors.New("income cannot be negative")
	}

	result := 0.0
	for index, taxBracket := range taxBrackets {
		if taxBracket.IncomeLowerBound > income {
			break
		}
		if index != len(taxBrackets)-1 {
			result += (math.Min(taxBrackets[index+1].IncomeLowerBound, income) - taxBracket.IncomeLowerBound) * taxBracket.TaxRate
		} else {
			result += (income - taxBracket.IncomeLowerBound) * taxBracket.TaxRate
		}
	}

	return result, nil

}

func checkTaxClasses(taxBrackets []TaxBracket) bool {

	for index, taxBracket := range taxBrackets {
		if index != len(taxBrackets)-1 &&
			(taxBracket.IncomeLowerBound >= taxBrackets[index+1].IncomeLowerBound ||
				taxBracket.TaxRate >= taxBrackets[index+1].TaxRate) {
			return false
		}
	}
	return true

}
