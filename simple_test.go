package gorasp

import (
	"fmt"
	"testing"
)

func TestEmpty(t *testing.T) {
	data := []int{}
	dataStructure := NewRankSelectSimple(data)
	val := dataStructure.rankOfIndex(0)

	if val != 0 {
		t.Fail()
	}

	val2 := dataStructure.index_with_rank(0)

	if val2 != 0 {
		t.Fail()
	}
}

func TestOnlyOnes(t *testing.T) {
	data := []int{1, 1, 1, 1, 1, 1}
	dataStructure := NewRankSelectSimple(data)
	for index, _ := range data {
		query_rank := dataStructure.rankOfIndex(index)
		query_index := dataStructure.index_with_rank(index)
		if query_rank != uint(index) {
			t.Fail()
			fmt.Printf("Query Rank: expected %d received %d\n", index, query_rank)
		}
		if query_index != index {
			t.Fail()
			fmt.Printf("Query Index: expected %d received %d\n", index, query_index)
		}
	}
}
