package main

import (
	"fmt"
)

const (
	white = byte('.')
	black = byte('#')
)

type grid [][]byte

func (s *grid) isNg(i, j int) bool {

	up := !(i == 0)
	down := !(i == len(*s)-1)
	left := !(j == 0)
	right := !(j == len((*s)[i])-1)

	black_count := 0
	if up && (*s)[i-1][j] == black {
		black_count++
	}
	if down && (*s)[i+1][j] == black {
		black_count++
	}
	if left && (*s)[i][j-1] == black {
		black_count++
	}
	if right && (*s)[i][j+1] == black {
		black_count++
	}
	return black_count != 2 && black_count != 4

}

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	var s grid

	for i := 0; i < h; i++ {
		var tmp []byte
		fmt.Scan(&tmp)
		s = append(s, tmp)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == black {
				if s.isNg(i, j) {
					fmt.Println("No")
					return
				}
			}
		}
	}
	fmt.Println("Yes")
}
