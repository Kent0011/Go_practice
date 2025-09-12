package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Takahashi struct {
	aNum int
	bNum int
	cNum int
}

func (t *Takahashi) calcMinContestNum() int {
	if t.aNum <= t.cNum {
		return t.aNum
	} else {
		return t.cNum
	}
}

func (t *Takahashi) calcContestNum() int {
	minContestNum := t.calcMinContestNum()
	contestNum := (t.aNum + t.bNum + t.cNum) / 3

	if contestNum < minContestNum {
		return contestNum
	} else {
		return minContestNum
	}
}

func Solve(aNum, bNum, cNum int) int {
	t := Takahashi{aNum, bNum, cNum}
	return t.calcContestNum()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var testNum int
	scanner.Scan()
	testNum, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < testNum; i++ {
		scanner.Scan()
		aNum, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		bNum, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		cNum, _ := strconv.Atoi(scanner.Text())
		fmt.Printf("%d\n", Solve(aNum, bNum, cNum))
	}
}
