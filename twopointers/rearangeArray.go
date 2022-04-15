package twopointers

func RearrangeArray(nums []int) []int {
	pos_idx, neg_idx := 0, 0
	res := []int{}
	for idx := 0; idx+idx < len(nums); idx++ {
		for ; nums[pos_idx] < 0; pos_idx++ {
		}
		res = append(res, nums[pos_idx])
		pos_idx++
		for ; nums[neg_idx] > 0; neg_idx++ {
		}
		res = append(res, nums[neg_idx])
		neg_idx++
	}
	return res
}
