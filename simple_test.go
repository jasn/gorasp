package gorasp

import (
	"fmt"
	"testing"
)

func TestEmpty(t *testing.T) {
	data := []int{}
	dataStructure := NewRankSelectSimple(data)
	val := dataStructure.RankOfIndex(0)

	if val != 0 {
		t.Fail()
	}

	val2, _ := dataStructure.IndexWithRank(0)

	if val2 != 0 {
		t.Fail()
	}
}

func TestOnlyOnes(t *testing.T) {
	data := []int{1, 1, 1, 1, 1, 1}
	dataStructure := NewRankSelectSimple(data)
	for index, _ := range data {
		query_rank := dataStructure.RankOfIndex(index)
		query_index, _ := dataStructure.IndexWithRank(index)
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

func TestRankAndIndexOfLast(t *testing.T) {
	data := []int{0, 1, 0, 1, 0, 1}
	dataStructure := NewRankSelectSimple(data)
	queryRankSecondToLast := dataStructure.RankOfIndex(5)
	queryRankLast := dataStructure.RankOfIndex(6)
	querySelectLast, err := dataStructure.IndexWithRank(3)

	if queryRankSecondToLast != 2 {
		fmt.Println("[TestRankAndIndexOfLast] Error. second-to-last-rank incorrect. Expected ", 2, " Received", queryRankSecondToLast)
		t.Fail()
	}

	if queryRankLast != 3 {
		fmt.Println("[TestRankAndIndexOfLast] Error. last rank incorrect. Expected ", 3, " Received", queryRankLast)
		t.Fail()
	}

	if querySelectLast != 6 {
		fmt.Println("[TestRankAndIndexOfLast] Error. last select incorrect. Expected ", 6, " Received", querySelectLast)
		t.Fail()
	}

	if err != nil {
		fmt.Println("[TestRankAndIndexOfLast] Error. last select gave unexpected error.")
		t.Fail()
	}

}
