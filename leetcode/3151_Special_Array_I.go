package leetcode

// func isArraySpecial(nums []int) bool {
// 	if len(nums) == 1 {
// 		return true
// 	}

// 	for i := 0; i < len(nums)-1; i++ {
// 		if ((nums[i]%2 == 0) && (nums[i+1]%2 == 0)) || ((nums[i]%2 != 0) && (nums[i+1]%2 != 0)) {
// 			return false
// 		}
// 		continue
// 	}

// 	return true
// }

func isArraySpecial(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]%2 == nums[i+1]%2 {
			return false
		}
	}

	return true
}
