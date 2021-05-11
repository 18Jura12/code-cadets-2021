package fizzbuzz

import (
	"github.com/pkg/errors"
	"strconv"
)

func Fizzbuzz(start, end int) (string, error) {

	if start > end {
		return "", errors.New("Start is greater than end!")
	}

	output := ""
	for i := start; i <= end; i++ {
		number := fizz(i) + buzz(i)
		if number == "" {
			output += strconv.Itoa(i) + " "
		} else {
			output += number + " "
		}
	}
	return output, nil

}

func fizz(x int) string {
	if x % 3 == 0 {
		return "Fizz"
	} else {
		return ""
	}
}

func buzz(x int) string {
	if x % 5 == 0 {
		return "Buzz"
	} else {
		return ""
	}
}

