/*
https://leetcode.com/problems/3sum/

Given an integer array nums, return all the triplets
[nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]

Input: nums = [0,1,1]
Output: []
*/

package Medium

import "sort"

func threeSum(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)
	if n < 3 {
		return result
	}

	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		// Skip duplicates for nums[i]
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1

		// Using 2 pointers
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				// Append triplet to result
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates for left and right
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// Move pointers
				left++
				right--
			} else if sum < 0 {
				left++ // Increase the sum
			} else {
				right-- // Decrease the sum
			}
		}
	}
	return result
}
