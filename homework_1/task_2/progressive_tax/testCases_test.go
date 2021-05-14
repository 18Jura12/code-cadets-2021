package progressivetax_test

import progressivetax "task2/progressive_tax"

type testCase struct {
	inputIncome      float64
	inputTaxBrackets []progressivetax.TaxBracket

	expectedTax    float64
	expectingError bool
}

func getTestCases() []testCase {
	return []testCase{
		{
			inputIncome: 7000,
			inputTaxBrackets: []progressivetax.TaxBracket{
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
			},

			expectedTax:    800,
			expectingError: false,
		},
		{
			inputIncome: 15000,
			inputTaxBrackets: []progressivetax.TaxBracket{
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
			},

			expectedTax:    2900,
			expectingError: false,
		},
		{
			inputIncome: -7000,
			inputTaxBrackets: []progressivetax.TaxBracket{
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
			},

			expectingError: true,
		},
		{
			inputIncome: 7000,
			inputTaxBrackets: []progressivetax.TaxBracket{
				{
					IncomeLowerBound: 0,
					TaxRate:          0,
				},
				{
					IncomeLowerBound: 1000,
					TaxRate:          0.3,
				},
				{
					IncomeLowerBound: 5000,
					TaxRate:          0.2,
				},
				{
					IncomeLowerBound: 10000,
					TaxRate:          0.3,
				},
			},

			expectingError: true,
		},
		{
			inputIncome: 7000,
			inputTaxBrackets: []progressivetax.TaxBracket{
				{
					IncomeLowerBound: 0,
					TaxRate:          0,
				},
				{
					IncomeLowerBound: 1000,
					TaxRate:          0.1,
				},
				{
					IncomeLowerBound: 500,
					TaxRate:          0.2,
				},
				{
					IncomeLowerBound: 10000,
					TaxRate:          0.3,
				},
			},

			expectingError: true,
		},
	}
}
