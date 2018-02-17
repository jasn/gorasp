package gorasp

// interface for rank select structure on bit arrays.
type RankSelect interface {
	// return the number of 1-bits to the left of the index'th bit.
	RankOfIndex(index int) uint

	// return the lowest index j where RankOfIndex(j) = rank.
	IndexWithRank(rank int) (int, error)

	// return the bit at index.
	At(index int) int
}
