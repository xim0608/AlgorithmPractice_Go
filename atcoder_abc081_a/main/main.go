package main

import (
	"fmt"
)

func main() {
	var (
		marbles string
		counter = 0
	)
	fmt.Scan(&marbles)
	for _, value := range marbles {
		if value == '1' {
			counter++
		}
	}
	fmt.Println(counter)
}
