package gorasp

type RankSelect interface {
	rankOfIndex(index int) uint
	index_with_rank(rank int) int
}
