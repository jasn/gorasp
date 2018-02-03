package gorasp

import (
	"errors"
	"math/bits"
)

type RankSelectFast struct {
	packedArray    []uint64
	partialRanks   []uint
	partialSelects []uint32
	n              int
}

func (self *RankSelectFast) getWordIndexOfWordWithKthOneBit(k int) (rank, wordIdx int) {
	bitIdx := self.partialSelects[k/64]
	tmpWordIdx := bitIdx / 64
	rankSoFar := int(self.partialRanks[tmpWordIdx])
	for wordIdx := bitIdx / 64; wordIdx < uint32(len(self.packedArray)); wordIdx++ {
		word := self.packedArray[wordIdx]
		onesCount := bits.OnesCount64(word)
		if rankSoFar+onesCount >= k {
			return rankSoFar, int(wordIdx)
		}
		rankSoFar += onesCount
	}
	rank = -1
	wordIdx = -1
	return
}

func (self *RankSelectFast) IndexWithRank(rank int) (int, error) {
	bitIdx := self.partialSelects[rank/64]
	if bitIdx == uint32(self.n+1) {
		return -1, errors.New("No element with thank rank.")
	}
	rankOfIdx := (rank / 64) * 64
	if rankOfIdx == rank {
		return int(bitIdx), nil
	}

	// find the word that has the 'rank'th 1-bit.
	rankExcludingWordIdx, wordIdx := self.getWordIndexOfWordWithKthOneBit(rank)
	if wordIdx == -1 {
		return -1, errors.New("No element with that rank.")
	}

	bitOffset := selectInWord(self.packedArray[wordIdx], rank-rankExcludingWordIdx)
	return wordIdx*64 + bitOffset, nil
}

func selectInWord(word uint64, rank int) int {
	rankCurr := int(0)
	for i := uint(0); i < 64; i += 1 {
		if rankCurr == rank {
			return int(i)
		}
		bitVal := (word & (1 << i)) >> i
		rankCurr += int(bitVal)
	}
	return 64
}

func (self *RankSelectFast) computePartialSelects() {
	allSelects := self.computeAllSelects()
	self.partialSelects = make([]uint32, computePackedLength(self.n))
	for i := 0; i < len(allSelects); i += 64 {
		self.partialSelects[i/64] = allSelects[i]
	}
}

func (self *RankSelectFast) computeAllSelects() []uint32 {
	result := make([]uint32, self.n)
	result[0] = 0
	next := int(1)
	for i, _ := range self.packedArray {
		for j := 0; j < 64; j += 1 {
			bitValue := getBit(self.packedArray, i*64+j)
			if bitValue == 1 {
				result[next] = uint32(i*64 + j + 1)
				next = next + 1
			}
		}
	}
	for i := next; i < self.n; i += 1 {
		result[i] = uint32(self.n + 1)
	}
	return result
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
	result.computePartialSelects()
	return result
}

func computePackedLength(n int) int {
	if n == 0 {
		return 0
	}
	return (n-1)/64 + 1
}
