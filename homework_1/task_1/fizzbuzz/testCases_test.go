package fizzbuzz_test

type testCase struct {

	inputStart int
	inputEnd int

	expectedOutput []string
	expectingError bool

}

func getTestCases() []testCase {
	return []testCase{
		{

			inputStart: 2,
			inputEnd: 3,

			expectedOutput: []string{ "2", "Fizz" },
		},
		{

			inputStart: 3,
			inputEnd: 1,

			expectedOutput: nil,
			expectingError: true,
		},
		{

			inputStart: -3,
			inputEnd: -1,

			expectedOutput: []string{ "Fizz", "-2", "-1" },
		},
		{

			inputStart: -3,
			inputEnd: 3,

			expectedOutput: []string{ "Fizz", "-2", "-1", "FizzBuzz", "1", "2", "Fizz" },
		},
		{

			inputStart: 1,
			inputEnd: 1,

			expectedOutput: []string{ "1" },
		},
		{

			inputStart: 4,
			inputEnd: 5,

			expectedOutput: []string{ "4", "Buzz" },
		},
		{

			inputStart: 14,
			inputEnd: 15,

			expectedOutput: []string{ "14", "FizzBuzz" },
		},
	}
}
