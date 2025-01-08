package leetcode

func twoSum(nums []int, target int) []int {
	var result []int
	var flag bool
	for i := 0; i < len(nums)-1; i++ {
		if flag == true {
			break
		}
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				result = append(result, i, j)
				flag = true
				break
			}
		}
	}

	return result
}
