package algorithm

import (
	"sort"
	"strings"
)

func IsAnagram(wordOne, wordTwo string) bool {

	lowerWordOne := strings.ToLower(wordOne)
	lowerWordTwo := strings.ToLower(wordTwo)
	if lowerWordOne == lowerWordTwo {
		return false
	}
	s1 := strings.Split(wordOne, "")
	sort.Strings(s1)
	sortedS1 := strings.Join(s1, "")

	s2 := strings.Split(wordTwo, "")
	sort.Strings(s2)

	sortedS2 := strings.Join(s2, "")

	return sortedS1 == sortedS2
}
