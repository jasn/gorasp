package gorasp

type RankSelectFast struct {
	packedArray  []int64
	partialRanks []int64
}

func setBit(array []int64, index, value int) {
	wordIndex := index / 64
	bitIndex := index % 64

	word := array[wordIndex]
	mask := int64(1) << uint(bitIndex)
	if value == 0 {
		mask = mask ^ 0
		word = word & mask
	} else {
		word = word | mask
	}
	array[wordIndex] = word
}

func NewRankSelectFast(array []int) *RankSelectFast {
	result := new(RankSelectFast)
	packedArrayLength := computePackedLength(array)
	packedArray := make([]int64, packedArrayLength)
	for index, bitValue := range array {
		setBit(packedArray, index, bitValue)
	}
	result.packedArray = packedArray
	return result
}

func computePackedLength(array []int) int {
	if len(array) == 0 {
		return 0
	}
	return (len(array)-1)/64 + 1
}
