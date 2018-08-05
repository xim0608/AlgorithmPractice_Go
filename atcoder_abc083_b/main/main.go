package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, a, b, k, total int
	var s string

	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)

	total = 0

	for i := 1; i < n+1; i++ {
		sum := 0
		s = fmt.Sprint(i)
		for i := 0; i < len(s); i++ {
			k, _ = strconv.Atoi(s[i:i+1])
			sum += k
		}
		if a <= sum && sum <= b {
			total += i
		}

	}
	fmt.Println(total)
}
