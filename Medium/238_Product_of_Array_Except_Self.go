/*
https://leetcode.com/problems/product-of-array-except-self/

Given an integer array nums, return an array answer such that answer[i]
is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to
fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without
using the division operation.

Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
*/

package Medium

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	pre := 1
	for i := 0; i < n; i++ {
		result[i] = pre
		pre *= nums[i]
	}

	post := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= post
		post *= nums[i]
	}

	return result
}
