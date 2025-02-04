/*
https://leetcode.com/problems/trapping-rain-water/

Given n non-negative integers representing an elevation map where the
width of each bar is 1, compute how much water it can trap after raining.

Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6

Input: height = [4,2,0,3,2,5]
Output: 9
*/

package Hard

func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	// Solvation through the pointers
	maxLeft := make([]int, n)
	maxRight := make([]int, n)

	maxLeft[0] = height[0]
	for i := 1; i < n; i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i])
	}

	maxRight[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i])
	}

	count := 0
	for i := 0; i < n; i++ {
		water := min(maxLeft[i], maxRight[i])
		if water > height[i] {
			count += water - height[i]
		}
	}
	return count
}
