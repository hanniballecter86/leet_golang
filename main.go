package main

import (
	"fmt"
	r "leetcode/traverse"
)

func main() {
	nums := []int{"eat", "tea", "tan", "ate", "nat", "bat"]}
	res := r.GroupAnagrams(nums)
	fmt.Println(res)
}
