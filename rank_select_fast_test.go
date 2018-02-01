package gorasp

import (
	"fmt"
	"testing"
)

func TestIndexOfRankFastSimpleAgree(t *testing.T) {
	array := make([]int, 137)
	for i, _ := range array {
		array[i] = i % 2
	}
	maxRank := len(array) / 2

	fast := NewRankSelectFast(array)
	simple := NewRankSelectSimple(array)

	for i := 0; i < maxRank+10; i++ {
		idxFast, errFast := fast.IndexWithRank(i)
		idxSimple, errSimple := simple.IndexWithRank(i)
		if errFast == nil && errSimple == nil {
			if idxSimple != idxFast {
				fmt.Println("[TestIndexOfRankFastSimpleAgree Error] ",
					"idxSimple: ", idxSimple, " idxFast: ", idxFast)

				t.Fail()
			}
		} else if errFast == nil && errSimple != nil {
			fmt.Println("[TestIndexOfRankFastSimpleAgree Error] ",
				"Simple gave error, while fast did not")
			t.Fail()
		} else if errFast != nil && errSimple == nil {
			fmt.Println("[TestIndexOfRankFastSimpleAgree Error] ",
				"Fast gave error, while simple did not")
			t.Fail()
		}
	}
}

func TestSelectInWord(t *testing.T) {
	wordsOneSetBit := []uint64{1, 2, 4, 8, 16, 1 << 63}
	expectedOne := []int{1, 2, 3, 4, 5, 64}

	wordsTwoSetBits := []uint64{1 + 2, 1 + 4, 2 + 8, 4 + 16}
	expectedTwo := []int{2, 3, 4, 5}

	wordsThreeSetBits := []uint64{1 + 2 + 4, 2 + 4 + 8, 2 + 8 + 32}
	expectedThree := []int{3, 4, 6}

	for i, word := range wordsOneSetBit {
		idx := selectInWord(word, 1)
		if idx != expectedOne[i] {
			fmt.Println("Expected ", expectedOne[i], " received ", idx)
			t.Fail()
		}
	}

	for i, word := range wordsTwoSetBits {
		idx := selectInWord(word, 2)
		if idx != expectedTwo[i] {
			fmt.Println("wordsTwoSetBits failed\n", "Expected ", expectedTwo[i], " received ", idx)
			t.Fail()
		}
	}

	for i, word := range wordsThreeSetBits {
		idx := selectInWord(word, 3)
		if idx != expectedThree[i] {
			fmt.Println("wordsThreeSetBits failed\n", "Expected ", expectedThree[i], " received ", idx)
			t.Fail()
		}
	}
}

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
		rankFast := fast.RankOfIndex(i)
		rankSimple := simple.RankOfIndex(i)
		if rankFast != rankSimple {
			fmt.Println("Fast differs from Simple")
			fmt.Printf("fast.RankOfIndex(%d) = %d while slow.RankOfIndex(%d) = %d",
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
		rankFast := fast.RankOfIndex(i)
		rankSimple := simple.RankOfIndex(i)
		if rankFast != rankSimple {
			fmt.Println("Fast differs from Simple")
			fmt.Printf("fast.RankOfIndex(%d) = %d while slow.RankOfIndex(%d) = %d",
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
