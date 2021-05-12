package fizzbuzz_test

import (
	"code-cadets-2021/homework_1/task_1/fizzbuzz"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func init() {

}

func TestFizzbuzz(t *testing.T) {

	for index, testCase := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", index, testCase), t, func() {

			actualOutput, actualErr := fizzbuzz.Fizzbuzz(testCase.inputStart, testCase.inputEnd)

			if testCase.expectingError {
				So(actualErr, ShouldNotBeNil)
			} else {
				So(actualErr, ShouldBeNil)
				So(actualOutput, ShouldResemble, testCase.expectedOutput)
			}

		})
	}

}
