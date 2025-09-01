package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var n, q int
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	q, _ = strconv.Atoi(scanner.Text())

	a, b := make([]int, n), make([]int, n)
	for i := range a {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}
	for i := range b {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}

	sum_minAB := 0
	for i := 0; i < n; i++ {
		sum_minAB += min(a[i], b[i])
	}

	for i := 0; i < q; i++ {
		var c string
		var x, y int
		scanner.Scan()
		c = scanner.Text()
		scanner.Scan()
		x, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		y, _ = strconv.Atoi(scanner.Text())
		x--
		switch c {
		case "A":
			diff := min(b[x], y) - min(a[x], b[x])
			sum_minAB += diff
			a[x] = y
		case "B":
			diff := min(a[x], y) - min(a[x], b[x])
			sum_minAB += diff
			b[x] = y
		}
		fmt.Println(sum_minAB)
	}
}
