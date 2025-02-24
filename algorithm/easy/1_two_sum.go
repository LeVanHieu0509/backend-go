package main

import "log"

/*
 1. Given an array of integers nums and an integer target,
    return indices of the two numbers such that they add up to target.
 2. You may assume that each input would have exactly one solution,
    and you may not use the same element twice.
 3. You can return the answerin any order.
*/

/*
Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

---------------------------------------
Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

---------------------------------------
Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]
*/

func twoSum(nums []int, target int) []int {
	var obj = map[int]int{}

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		value, exists := obj[diff]

		if exists {
			return []int{value, i}
		} else {
			obj[nums[i]] = i
		}
	}
	return []int{-1, -1}
}

func main() {
	var (
		num = []int{3, 2, 4}
	)

	log.Println(twoSum(num, 6))
}
