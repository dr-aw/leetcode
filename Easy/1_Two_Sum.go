/*
https://leetcode.com/problems/two-sum/

Given an array of integers nums and an integer target, return
indices of the two numbers such that they add up to target.

Input: nums = [2,7,11,15], target = 9
Output: [0,1]

Input: nums = [3,2,4], target = 6
Output: [1,2]
*/

package Easy

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, element := range nums {
		m[element] = i
	}

	for i, num := range nums {
		diff := target - num
		if j, ok := m[diff]; ok && j != i {
			return []int{i, j}
		}
	}
	return nil
}
