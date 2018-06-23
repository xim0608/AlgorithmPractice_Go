package main

import (
	"fmt"
)

func main() {
	var (
		a, b, c, x int
		counter    = 0
		total      int
	)

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&x)

	for i := 0; i <= a; i++ {
		if i*500 > x {
			break
		}
		for j := 0; j <= b; j++ {
			if i*500+j*100 > x {
				break
			}
			for k := 0; k <= c; k++ {
				total = i*500 + j*100 + k*50
				if i*500+j*100+k*50 > x {
					break
				}
				if total == x {
					counter++
				}
			}
		}
	}
	fmt.Println(counter)
}
