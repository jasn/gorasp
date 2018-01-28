package gorasp

type RankSelect interface {
	rank_of_index(index int) int
	index_with_rank(rank int) int
}
