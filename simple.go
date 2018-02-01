package gorasp

import (
	"errors"
)

type RankSelectSimple struct {
	array []int
}

func NewRankSelectSimple(array []int) *RankSelectSimple {
	obj := new(RankSelectSimple)
	obj.array = array
	return obj
}

func (s *RankSelectSimple) RankOfIndex(index int) uint {
	var result = 0
	for _, val := range s.array[:index] {
		result += val
	}
	return uint(result)
}

func (s *RankSelectSimple) IndexWithRank(rank int) (int, error) {
	var count = 0
	for i, val := range s.array {
		if count == rank {
			return i, nil
		}
		count += val
	}
	if count == rank {
		return len(s.array), nil
	}
	return -1, errors.New("No element with that rank.")
}
