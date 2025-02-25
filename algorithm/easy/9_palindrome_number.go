package main

import "strconv"

func main() {
	print(isPalindrome, 242)
	print(isPalindromeWay2, 242)
	print(isPalindromeOptimize, 242)
}

type Func func(int) bool

func print(fun Func, x int) {
	println(fun(x))
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
func isPalindrome(x int) bool {
	// Nếu x là số âm hoặc kết thúc bằng 0 (nhưng không phải là 0), thì không thể là số palindrome
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	// Nếu x == 0 trả về true
	if x == 0 {
		return true
	}
	// Tạo biến reserve để lưu giá trị của x
	var reserve = x
	// Tạo biến temp để lưu giá trị của x
	var temp = x
	// Tạo biến result để lưu giá trị của 0
	var result = 0
	// Vòng lặp để đảo ngược số
	for reserve != 0 {
		// Tạo biến remainder để lưu giá trị của reserve % 10
		var remainder = reserve % 10
		println("remainder", remainder)
		// Tạo biến result để lưu giá trị của result * 10 + remainder
		result = result*10 + remainder
		println("result", result)
		// Tạo biến reserve để lưu giá trị của reserve / 10
		reserve = reserve / 10
		println("reserve", reserve)
	}
	// Kiểm tra nếu temp == result trả về true
	if temp == result {
		return true
	}
	// Trả về false
	return false
}

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

func isPalindromeOptimize(x int) bool {
	// Nếu x là số âm hoặc kết thúc bằng 0 (nhưng không phải là 0), thì không thể là số palindrome
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reversed := 0
	original := x
	// Đảo ngược số
	for x > 0 {
		reversed = reversed*10 + x%10 // Toán tử chia lấy phần dư. Chia toán hạng đầu tiên cho toán hạng thứ hai và tạo ra phần dư. => lấy phần dư
		x /= 10                       // Toán tử chia. Chia toán hạng đầu tiên cho toán hạng thứ hai và tạo ra thương. => lấy thương
	}
	// Kiểm tra xem số gốc có bằng với số đã đảo ngược hay không
	return original == reversed
}
