package main

import (
	"fmt"
	"strconv"
)

func revert(a int) int {
	str := fmt.Sprint(a)
	reversed_str := make([]byte, len(str))
	for i, v := range str {
		reversed_str[len(str)-i-1] = byte(v)
	}
	ret_int, err := strconv.Atoi(string(reversed_str))
	if err != nil {
		return 0
	}
	return ret_int
}

func main() {
	a := make([]int, 10)
	fmt.Scan(&a[0], &a[1])
	for i := 2; i < 10; i++ {
		a[i] = revert(a[i-1] + a[i-2])
	}
	fmt.Println(a[9])
}
