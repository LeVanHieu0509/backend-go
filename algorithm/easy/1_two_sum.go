package main

import (
	"log"
	"strings"
)

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
	log.Println(towSumPracticeDay2(num, 6))

	log.Println(twoSumInter([]StructTwoSum{
		{Key: "A", Value: "25"},
		{Key: "B", Value: "26"},
		{Key: "C", Value: "27"},
		{Key: "D", Value: "28"},
		{Key: "E", Value: "25"},
		{Key: "G", Value: "26"},
	}))

}

func towSumPracticeDay2(nums []int, target int) []int {
	//1. Create object empty to save value element to key of object
	var object = map[int]int{}

	//2. Loop the nums
	for i := 0; i < len(nums); i++ {
		//3. Create variable diff to save value of target - nums[i]
		var diff = target - nums[i]

		//4. Check if the diff is exists key in object
		if value, exists := object[diff]; exists {

			//5. If exists return value of key and i of nums
			return []int{value, i}
		} else {
			//6. else save the element of nums to object with element of nums as key and i as value
			object[nums[i]] = i
		}
	}
	//7. error return -1,-1
	return []int{-1, -1}
}

type StructTwoSum struct {
	Key   string
	Value string
}

func twoSumInter(arr []StructTwoSum) []StructTwoSum {
	// Khởi tạo map để lưu trữ các key tương ứng với value
	var newObj = map[string][]string{}
	var newArr = []StructTwoSum{}

	// Duyệt qua các phần tử trong arr chỉ với 1 vòng lặp
	for _, item := range arr {
		// Kiểm tra xem value đã có trong map chưa
		if _, exists := newObj[item.Value]; exists {
			// Nếu đã tồn tại, thêm key vào value tương ứng
			newObj[item.Value] = append(newObj[item.Value], item.Key)
		} else {
			// Nếu chưa có, khởi tạo giá trị mới cho key và value
			newObj[item.Value] = []string{item.Key}
		}
	}

	// Duyệt qua map để chuyển các nhóm thành kết quả cuối cùng
	// (Tạo newArr từ newObj)
	// map trong Go có độ phức tạp trung bình là O(1) cho các thao tác truy xuất và chèn
	for value, keys := range newObj {
		newArr = append(newArr, StructTwoSum{
			Key:   value,
			Value: strings.Join(keys, ","), // Nối các key lại với nhau
		})
	}

	// Trả về mảng các StructTwoSum đã nhóm
	return newArr
}
