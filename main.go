package main

import (
	"fmt"
	p "leetcode/twopointers"
)

func main() {
	nums := []int{-2, 0, 1, 1, 2}
	res := p.ThreeSum(nums)
	fmt.Println(res)
}
