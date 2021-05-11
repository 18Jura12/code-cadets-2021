package fizzbuzz_test

import (
	"code-cadets-2021/homework_1/fizzbuzz"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func init() {

}

func TestFizzbuzz(t *testing.T) {

	for idx, tc := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", idx, tc), t, func() {

			actualOutput, actualErr := fizzbuzz.Fizzbuzz(tc.inputStart, tc.inputEnd)

			if tc.expectingError {
				So(actualErr, ShouldNotBeNil)
			} else {
				So(actualErr, ShouldBeNil)
				So(actualOutput, ShouldResemble, tc.expectedOutput)
			}

		})
	}

}
