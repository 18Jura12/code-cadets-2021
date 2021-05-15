package main

import (
	"flag"
	"fmt"

	"task1/fizzbuzz"
)


func main() {

	var start, end int
	flag.IntVar(&start, "start", 1, "number to start the game from")
	flag.IntVar(&end, "end", 10, "number to stop the game at (inclusive)")
	flag.Parse()

	output, err := fizzbuzz.Fizzbuzz(start, end)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(output)

 }
