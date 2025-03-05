package main

import (
	"log"
	"strconv"
)

func main() {
	log.Println(reverseString([]string{"1", "2", "3", "4"}))

	print(isPalindrome, 242)
	print(isPalindromeWay2, 242)
	print(isPalindromeOptimize, 242)
	print(isPalindromeBest, 121)
	print(isPalindromeMoreOptimize, 2332)
	print(isPalindromePractice, 232)

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

// Thời gian: O(d) = O(log₁₀(x))
// Thuật toán lặp qua d chữ số của x, do đó có độ phức tạp O(log₁₀(x)) (vì số có n chữ số thì n ≈ log₁₀(x)).
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

/*
	[a,b,c] => [c,b,a]
*/

func reverseString(s []string) []string {
	// Tạo biến left để lưu giá trị của 0
	left := 0
	// Tạo biến right để lưu giá trị của len(s) - 1
	right := len(s) - 1
	// Vòng lặp để đảo ngược mảng

	// 0 -> 1 -> 2 -> 3

	for left < right {
		// Tạo biến temp để lưu giá trị của s[left]
		temp := s[left]
		// Gán s[left] = s[right]
		s[left] = s[right]
		// Gán s[right] = temp
		s[right] = temp
		// Tăng left lên 1
		left++
		// Giảm right đi 1
		right--
	}

	return s
}

// Two-Pointer Technique
func isPalindromeBest(x int) bool {
	xStr := strconv.Itoa(x)

	j := len(xStr) - 1
	for i := 0; i < len(xStr); i++ {
		if xStr[i] != xStr[j] {
			return false
		}
		j--
	}
	return true
}

// Thời gian: O(log₁₀(x))
func isPalindromeMoreOptimize(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return x == reversed || x == reversed/10
}

func isPalindromePractice(x int) bool {
	reserved := 0

	for x > reserved {
		reserved = reserved*10 + x%10 // 1-12-123
		x /= 10                       // 1-2-3
	}
	log.Printf("reserved: %d, %d", x, reserved)

	/*
		1. x = 121 → reserved = 1
		2. x = 12 → reserved = 1*10 + 2 = 12
		3. x = 1 → return x == reserved/10 -> 1 = 12/10 = 1
	*/

	// x == reversed/10 → Dành cho số có chữ số lẻ, loại bỏ chữ số dư
	return x == reserved || x == reserved/10

}

/*
1. Dùng hai con trỏ (Two-Pointer Technique):
	- Một con trỏ i bắt đầu từ đầu chuỗi (0).
	- Một con trỏ j bắt đầu từ cuối chuỗi (len(xStr) - 1).
	- So sánh các cặp ký tự đối xứng nhau (xStr[i] và xStr[j]).
	- Nếu có một cặp không khớp, trả về false.
	- Nếu duyệt hết mà không có sự khác biệt, trả về true.

2. Độ phức tạp thời gian:
	- O(n) với n là độ dài của chuỗi số (len(xStr)).
	- Vì chỉ duyệt qua một nửa chuỗi nên thời gian thực tế là O(n/2) nhưng vẫn được biểu diễn là O(n).
3. Nhận xét:
	- không tối ưu về mặt bộ nhớ do cần tạo chuỗi xStr

*/
