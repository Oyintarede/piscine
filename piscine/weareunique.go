package piscine

import (
	"fmt"
)

func WeAreUnique(str1, str2 string) int{
	if str1 == "" && str2 == "" {
		return -1
	}

	set1 := make(map[rune]bool)
	set2 := make(map[rune]bool)

	for _, r := range str1 {
		set1[r] = true
	}

	for _, r := range str2 {
		set2[r] = true
	}

	count := 0

	for r := range set1 {
		if !set2[r] {
			count++
		}
	}

	for r := range set2 {
		if !set1[r] {
			count++
		}
	}
	return count
}
