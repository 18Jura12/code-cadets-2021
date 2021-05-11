package fizzbuzz

import (
	"fmt"
)

func Fizzbuzz(start, end int) {

	for i := start; i <= end; i++ {
		output := fizz(i) + buzz(i)
		if output == "" {
			fmt.Printf("%v ", i)
		} else {
			fmt.Printf("%v ", output)
		}
	}

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

