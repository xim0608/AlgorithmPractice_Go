package main

import "fmt"

func diviedbyTwo(a int) int {
	// return 0 when % 2 == 1
	if a%2 == 0 {
		return a / 2
	} else {
		return 0
	}
}

func main() {
	var (
		n int
		counter = 0
	)
	fmt.Scan(&n)
	blackboard := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&blackboard[i])
	}
	condition := true
	for condition {
		for i := 0; i < n; i++ {
			returnN := diviedbyTwo(blackboard[i])
			blackboard[i] = returnN
			if returnN == 0 {
				condition = false
			}
		}
		//fmt.Println(blackboard)
		counter++
	}
	fmt.Println(counter-1)
}
