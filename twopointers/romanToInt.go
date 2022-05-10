package twopointers

import (
	"errors"
	"strings"
)

/*
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/

func RomanToInt(s string) int {
	chars := strings.Split(s, "")
	totalNum := 0
	sizeMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000}
	validPairs := [][]string{{"I", "V"}, {"I", "X"}, {"X", "L"}, {"X", "C"}, {"C", "D"}, {"C", "M"}}
	for i, j := 0, 1; i < len(chars); i, j = i+1, j+1 {
		if i == len(chars)-1 {
			totalNum += sizeMap[chars[i]]
			return totalNum
		} else {
			pre := chars[i]
			cur := chars[j]
			preNum := sizeMap[pre]
			curNum := sizeMap[cur]
			if curNum > preNum {
				if checkValid([]string{pre, cur}, validPairs) {
					totalNum += preNum * -1
				} else {
					errors.New("It's not a valid Roman number")
				}
			} else {
				totalNum += preNum
			}
		}
	}
	return totalNum
}

func checkValid(a []string, b [][]string) bool {
	for _, v := range b {
		if checkEqual(a, v) {
			return true
		}
	}
	return false
}

func checkEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
