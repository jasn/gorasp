package gorasp

import (
	"fmt"
)

type RankSelect interface {
	rank_of_index(index int) int
	index_with_rank(rank int) int
}

func printAThing(v RankSelect) {
	val := v.rank_of_index(2)
	fmt.Println(val)
}

func main() {
	fmt.Println("Hello, rasp!")

	val := NewRankSelectSimple([]int{0, 0, 1, 0, 1, 1, 0})
	printAThing(val)
	fmt.Println(val)
	fmt.Println(val.rank_of_index(1))
}
