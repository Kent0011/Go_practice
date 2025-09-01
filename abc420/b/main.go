package main

import (
	"fmt"
	"strings"
)

func main() {
	var people_num, vote_num int
	fmt.Scan(&people_num, &vote_num)

	scores := make([]int, people_num)
	votes := make([][]byte, people_num)

	for i := 0; i < people_num; i++ {
		fmt.Scan(&votes[i])
	}

	for i := 0; i < vote_num; i++ {
		vote_0 := 0
		vote_1 := 0
		for j := 0; j < people_num; j++ {
			if votes[j][i] == '0' {
				vote_0++
			} else {
				vote_1++
			}
		}
		if vote_0 == 0 || vote_1 == 0 {
			for j := 0; j < people_num; j++ {
				scores[j]++
			}
		} else if vote_0 > vote_1 {
			for j := 0; j < people_num; j++ {
				if votes[j][i] == '1' {
					scores[j]++
				}
			}
		} else if vote_0 < vote_1 {
			for j := 0; j < people_num; j++ {
				if votes[j][i] == '0' {
					scores[j]++
				}
			}
		}
	}

	max_score := 0
	winners := make([]string, 0, people_num)
	for i, v := range scores {
		if v > max_score {
			max_score = v
			winners = winners[:0]
			winners = append(winners, fmt.Sprint(i+1))
		} else if v == max_score {
			winners = append(winners, fmt.Sprint(i+1))
		}
	}

	fmt.Println(strings.Join(winners, " "))
}
