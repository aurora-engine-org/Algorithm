package main

import "fmt"

func main() {
	//nums := []int{1, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4}
	fmt.Println(isBalanced(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 4,
				},
			},
		},
	}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(deep(root.Left)-deep(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func deep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := deep(root.Left)
	r := deep(root.Right)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}
func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
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
