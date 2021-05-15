package progressivetax_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"task2/progressive_tax"
)

func TestCalculateTax(t *testing.T) {

	for index, testCase := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", index, testCase), t, func() {

			actualOutput, actualErr := progressivetax.CalculateTax(testCase.inputIncome, testCase.inputTaxBrackets)

			if testCase.expectingError {
				So(actualErr, ShouldNotBeNil)
			} else {
				So(actualErr, ShouldBeNil)
				So(actualOutput, ShouldResemble, testCase.expectedTax)
			}

		})
	}

}
