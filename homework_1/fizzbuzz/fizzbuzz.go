package main

import (
	"flag"
)

func main() {

	var start, end int
	flag.IntVar(&start, "start", 1, "number to start the game from")
	flag.IntVar(&end, "end", 10, "number to stop the game at")

	for i := start; i <= end; i++ {

	}

}

