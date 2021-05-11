package main

import (
	"code-cadets-2021/homework_1/fizzbuzz"
	"flag"
)


func main() {

	var start, end int
	flag.IntVar(&start, "start", 1, "number to start the game from")
	flag.IntVar(&end, "end", 10, "number to stop the game at")
	flag.Parse()
	fizzbuzz.Fizzbuzz(start, end)



 }