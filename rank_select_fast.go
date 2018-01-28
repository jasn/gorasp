package gorasp

import (
	"math/bits"
)

type RankSelectFast struct {
	packedArray  []uint64
	partialRanks []uint
	n            int
}

func (self *RankSelectFast) IndexWithRank(rank int) int {
	return 65
}

func (self *RankSelectFast) RankOfIndex(index int) uint {
	partialRank := self.partialRanks[index/64]

	if index%64 == 0 {
		return partialRank
	}

	word := self.packedArray[index/64]

	mask := uint64(1) << (uint(index % 64))
	mask = mask - 1
	return partialRank + uint(bits.OnesCount64(word&mask))
}

func (self *RankSelectFast) computePartialRanks() {
	self.partialRanks = make([]uint, computePackedLength(self.n))
	sum := uint(0)
	for i := 0; i < self.n; i += 1 {
		if i%64 == 0 {
			self.partialRanks[i/64] = sum
		}
		bitValue := getBit(self.packedArray, i)
		sum += bitValue
	}
}

func getBit(array []uint64, index int) uint {
	wordIndex := index / 64
	bitIndex := index % 64

	word := array[wordIndex]
	mask := uint64(1) << uint(bitIndex)

	val := uint((word & mask) >> uint(bitIndex))
	return val
}

func setBit(array []uint64, index, value int) {
	wordIndex := index / 64
	bitIndex := index % 64

	word := array[wordIndex]
	mask := uint64(1) << uint(bitIndex)
	if value == 0 {
		mask = ^uint64(0) ^ mask
		word = word & mask
	} else {
		word = word | mask
	}
	array[wordIndex] = word
}

func NewRankSelectFast(array []int) *RankSelectFast {
	result := new(RankSelectFast)
	packedArrayLength := computePackedLength(len(array))
	packedArray := make([]uint64, packedArrayLength)
	for index, bitValue := range array {
		setBit(packedArray, index, bitValue)
	}
	result.packedArray = packedArray
	result.n = len(array)
	result.computePartialRanks()
	return result
}

func computePackedLength(n int) int {
	if n == 0 {
		return 0
	}
	return (n-1)/64 + 1
}
