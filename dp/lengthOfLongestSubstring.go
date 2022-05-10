package dp

import "strings"

func LengthOfLongestSubstring(s string) int {
	chars := strings.Split(s, "")
	continueChars, maxLen := []string{}, 0
	for _, v := range chars {
		contain, idx := sliceContainIdx(continueChars, v)
		if contain {
			continueChars = append(continueChars, v)
			continueChars = continueChars[idx+1:]
		} else {
			continueChars = append(continueChars, v)
		}
		if len(continueChars) > maxLen {
			maxLen = len(continueChars)
		}
	}
	return maxLen
}

func sliceContainIdx(a []string, b string) (bool, int) {
	for i, v := range a {
		if v == b {
			return true, i
		}
	}
	return false, 0
}
