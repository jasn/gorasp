package gorasp

import (
	"fmt"
)

func printAThing(v RankSelect) {
	val := v.rankOfIndex(2)
	fmt.Println(val)
}

func main() {
	fmt.Println("Hello, rasp!")

	val := NewRankSelectSimple([]int{0, 0, 1, 0, 1, 1, 0})
	printAThing(val)
	fmt.Println(val)
	fmt.Println(val.rankOfIndex(1))
}
