package gorasp

type RankSelect interface {
	RankOfIndex(index int) uint
	IndexWithRank(rank int) (int, error)
	At(index int) int
}
