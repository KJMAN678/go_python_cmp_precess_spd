package main

import (
	"fmt"
	"time"
)

func main() {

	var x int
	fmt.Scanf("%d", &x)

	start := time.Now()

	var answer_list = []int{}

	for i := 1; i <= x; i++ {
		if x%i == 0 {
			answer_list = append(answer_list, i)
		}
	}

	fmt.Println("answer: ", answer_list)

	end := time.Now()
	fmt.Printf("elapsed_time:%f[sec]\n", (end.Sub(start)).Seconds())

}
