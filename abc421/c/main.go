package main

import (
	"fmt"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]byte, 2*n)
	fmt.Scan(&s)

	b_index := make([]int, 0, n)

	for i, v := range s {
		if v == 'B' {
			b_index = append(b_index, i)
		}
	}

	total_move_024 := 0
	total_move_135 := 0

	for i := 0; i < n; i++ {
		total_move_024 += abs(2*i - b_index[i])
		total_move_135 += abs(2*i + 1 - b_index[i])
	}
	if total_move_024 < total_move_135 {
		fmt.Println(total_move_024)
	} else {
		fmt.Println(total_move_135)
	}
}
