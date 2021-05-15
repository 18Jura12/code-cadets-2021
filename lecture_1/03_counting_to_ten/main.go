package main

import (
	"fmt" //dodatne stvari za formatiranje
	"log"  //neki drugi loggeri mogu informirati da se vidi greška na serveru
)

func main() {

	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	i = 0
	for {
		log.Println(i)
		i++
		if i == 10 {
			break
		}
	}
}
