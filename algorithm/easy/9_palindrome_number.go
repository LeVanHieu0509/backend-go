package main

import "strconv"

func main() {
	println(isPalindrome(242))
	println(isPalindromeWay2(242))

}

/*
Given an integer x, return true if x is a palindrome, and false otherwise.

Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/

// Check xuôi ngược đều giống nhau
// 7ms
func isPalindrome(x int) bool {
	//1. Find rule palindrome and logic handle it
	//2. If x < 0 return false
	if x < 0 || x%10 == 0 {
		return false
	}
	//3. If x == 0 return true
	if x == 0 {
		return true
	}
	//5. Create variable reverse to save value of x
	reserve := x
	//7. Create variable result to save value of 0
	result := 0
	//8. Loop the reverse
	for reserve != 0 {
		//9. Create variable remainder to save value of reverse % 10
		var remainder = reserve % 10
		//10. Create variable result to save value of result * 10 + remainder
		result = result*10 + remainder
		//11. Create variable reverse to save value of reverse / 10
		x /= 10
	}

	//13. Return false
	return x == result
}

// 2ms
func isPalindromeWay2(x int) bool {
	xStr := strconv.Itoa(x)

	length := len(xStr) - 1
	for i := 0; i < length-i; i++ {
		if xStr[i] != xStr[length-i] {
			return false
		}
	}
	return true
}
