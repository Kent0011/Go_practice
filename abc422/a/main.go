package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	var world, stage int

	world, _ = strconv.Atoi(string(s[0]))
	stage, _ = strconv.Atoi(string(s[2]))

	if stage == 8 {
		world += 1
		stage = 1
	} else {
		stage += 1
	}

	fmt.Printf("%d-%d\n", world, stage)
}
