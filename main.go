package main

import "fmt"

func main() {
	nums := []int{1, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > 0 && nums[1] > 0 {
			return nums[0] + nums[1]
		}
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}
	var sum int
	for i := 0; i < len(nums); i++ {
		if sum+nums[i] < sum {
			s := maxSubArray(nums[i+1:])
			if sum == 0 {
				return s
			}

			// 尝试继续累加
			if s > sum {
				return s
			} else {
				return sum
			}
		}
		//if sum+nums[i] < sum {
		//	if i+1 < len(nums) {
		//		if sum+nums[i]+nums[i+1] < sum {
		//			s := maxSubArray(nums[i+1:])
		//			if sum == 0 {
		//				return s
		//			}
		//			if s > sum {
		//				return s
		//			} else {
		//				return sum
		//			}
		//		}
		//	}
		//}
		sum += nums[i]
	}
	return sum
}
