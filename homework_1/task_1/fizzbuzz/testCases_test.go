package fizzbuzz_test

type testCase struct {

	inputStart int
	inputEnd int

	expectedOutput string
	expectingError bool

}

func getTestCases() []testCase {
	return []testCase{
		{

			inputStart: 2,
			inputEnd: 3,

			expectedOutput: "2 Fizz ",
		},
		{

			inputStart: 3,
			inputEnd: 1,

			expectedOutput: "",
			expectingError: true,
		},
		{

			inputStart: -3,
			inputEnd: -1,

			expectedOutput: "Fizz -2 -1 ",
		},
	}
}
