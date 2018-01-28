package gorasp

type RankSelectSimple struct {
	array []int
}

func NewRankSelectSimple(array []int) *RankSelectSimple {
	obj := new(RankSelectSimple)
	obj.array = array
	return obj
}

func (s *RankSelectSimple) rankOfIndex(index int) uint {
	var result = 0
	for _, val := range s.array[:index] {
		result += val
	}
	return uint(result)
}

func (s *RankSelectSimple) index_with_rank(rank int) int {
	var count = 0
	for i, val := range s.array {
		if count == rank {
			return i
		}
		count += val
	}
	return len(s.array)
}
