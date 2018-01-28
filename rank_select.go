package gorasp

type RankSelect struct {
	packed_array  []int64
	partial_ranks []int64
}

func computePackedLength(array []int) {
	if len(array) == 0 {
		return 0
	}
	return (len(array)-1)/64 + 1
}

func setBit(array []int, index, value int) {
	wordIndex := index / 64
	bitIndex := index % 64

	word := array[wordIndex]
	mask := 1 << bitIndex
	if value == 0 {
		mask = mask ^ 0
		word = word & mask
	} else {
		word := word | mask
	}
	array[wordIndex] = word
}

func NewRankSelect(array []int) *RankSelectSimple {
	result := new(RankSelect)
	packed_array_length := computePackedLength(array)
	packed_array := make([]int64, packed_array_length)
	for index, bit_value := range array {
		setBit(packed_array, index, bit_value)

	}
	obj := new(RankSelectSimple)
	obj.array = array
	return obj
}
