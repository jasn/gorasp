package gorasp

import (
	"fmt"
	"testing"
)

func TestSetBit(t *testing.T) {
	array := []int{1, 0, 1, 0, 0, 1}
	expected := []uint64{1, 1, 5, 5, 5, 37}
	structure := []uint64{0}
	for i, v := range array {
		setBit(structure, i, v)
		if structure[0] != expected[i] {
			fmt.Printf("[TestSetBit] Expected %d received %d", expected[i], structure[0])
			fmt.Println()
			t.Fail()
		}
	}

}

func TestRankOfIndexSimpleFastAgreeSingleWord(t *testing.T) {
	array := []int{1, 0, 1, 0, 0, 1}
	fast := NewRankSelectFast(array)
	simple := NewRankSelectSimple(array)

	for i, _ := range array {
		rankFast := fast.rankOfIndex(i)
		rankSimple := simple.rankOfIndex(i)
		if rankFast != rankSimple {
			fmt.Println("Fast differs from Simple")
			fmt.Printf("fast.rankOfIndex(%d) = %d while slow.rankOfindex(%d) = %d",
				i, rankFast, i, rankSimple)
			fmt.Println()
			t.Fail()
		}
	}
}

func TestRankOfIndexSimpleFastAgreeMultipleWords(t *testing.T) {
	array := make([]int, 137)
	for i, _ := range array {
		array[i] = i % 2
	}

	fast := NewRankSelectFast(array)
	simple := NewRankSelectSimple(array)

	for i, _ := range array {
		rankFast := fast.rankOfIndex(i)
		rankSimple := simple.rankOfIndex(i)
		if rankFast != rankSimple {
			fmt.Println("Fast differs from Simple")
			fmt.Printf("fast.rankOfIndex(%d) = %d while slow.rankOfindex(%d) = %d",
				i, rankFast, i, rankSimple)
			fmt.Println()
			t.Fail()
		}
	}
}

func TestLengthComputer(t *testing.T) {
	lengths := []int{1, 63, 64, 65, 127, 128, 129}
	outputs := []int{1, 1, 1, 2, 2, 2, 3}

	for i, length := range lengths {
		expected := outputs[i]
		actual := computePackedLength(length)
		if expected != actual {
			fmt.Printf("Expected %d but received %d", expected, actual)
			t.Fail()
		}
	}
}
