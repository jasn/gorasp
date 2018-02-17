package gorasp

import (
	"errors"
	"math/bits"
)

// struct that will implement the RankSelect interface.
type RankSelectFast struct {
	packedArray    []uint64
	partialRanks   []uint
	partialSelects []uint32
	n              int
}

func (self *RankSelectFast) At(index int) int {
	if index >= self.n {
		return 0 // return an error or just pretend everything beyond n is 0?
	}
	return int(getBit(self.packedArray, index))
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

// find the first bit in word with rank ones before it.
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

// helper method to build the lookup table used when answer select(i) queries.
// This function comptues the table select(i) for i = 64k, k = 0..n/64
func (self *RankSelectFast) computePartialSelects() {
	allSelects := self.computeAllSelects()
	self.partialSelects = make([]uint32, computePackedLength(self.n))
	for i := 0; i < self.n; i += 64 {
		self.partialSelects[i/64] = allSelects[i]
	}
}

// helper method to build the lookup table used when answer select(i) queries.
// this function computes the table select(i) for 0 <= i <= n.
func (self *RankSelectFast) computeAllSelects() []uint32 {
	result := make([]uint32, self.n+1)
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
	if index >= self.n {
		lastPartial := self.partialRanks[len(self.partialRanks)-1]
		lastWord := self.packedArray[len(self.packedArray)-1]
		return lastPartial + uint(bits.OnesCount64(lastWord))
	}
	partialRank := self.partialRanks[index/64]

	if index%64 == 0 {
		return partialRank
	}

	word := self.packedArray[index/64]

	mask := uint64(1) << (uint(index % 64))
	mask = mask - 1
	return partialRank + uint(bits.OnesCount64(word&mask))
}

// helper method for building the lookup tables used when answering rank(i) queries.
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

// helper method for accessing individual bits in a 'packed array'
// index refers to the bit-index, i.e. the first word stores bits for indices 0-63.
func getBit(array []uint64, index int) uint {
	wordIndex := index / 64
	bitIndex := index % 64

	word := array[wordIndex]
	mask := uint64(1) << uint(bitIndex)

	val := uint((word & mask) >> uint(bitIndex))
	return val
}

// helper method for setting bits in a packed array.
// index refers to the bit-index, i.e. the first word stores bits for indices 0-63.
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

// constructor for RankSelectFast.
// this function computes all the necessary tables, so the structure is ready for querying when this returns.
// array is not a packed array, i.e. every bit is in its own int.
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
