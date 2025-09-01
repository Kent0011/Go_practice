package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	var x int
	var y string

	fmt.Scan(&x, &y)

	if s[x-1] == y {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
