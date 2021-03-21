package __shu_zu_zhong_zhong_fu_de_shu_zi_lcof

/*
  找出数组中重复的数字。

  在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
  数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
  请找出数组中任意一个重复的数字。

  思路：
     1. map

*/
func findRepeatNumber(nums []int) int {
	ans := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := ans[nums[i]]; ok {
			return nums[i]
		} else {
			ans[nums[i]] = 1
		}
	}
	return 0
}
