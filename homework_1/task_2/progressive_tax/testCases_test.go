package progressive_tax_test

type testCase struct {

	inputIncome string

	expectedTax string
	expectingError bool

}

func getTestCases() []testCase {
	return []testCase {
		{
			inputIncome: "7000 HRK",
			expectedTax: "800 HRK",
			expectingError: false,
		},
		{
			inputIncome: "15000 HRK",
			expectedTax: "2900 HRK",
			expectingError: false,
		},
		{
			inputIncome: "15000",
			expectingError: true,
		},
		{
			inputIncome: "-7000 HRK",
			expectingError: true,
		},
	}
}
