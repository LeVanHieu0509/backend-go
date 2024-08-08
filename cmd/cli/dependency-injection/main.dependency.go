package main

import "fmt"

/*
Đoạn mã trên tạo một hệ thống mô phỏng hoạt động leo núi đá.
Hệ thống này có các thành phần chính bao gồm một cấu trúc dữ liệu RockClimber, một giao diện SafetyPlacer,
và các loại đối tượng triển khai SafetyPlacer như NOPSafetyPlace. Dưới đây là giải thích chi tiết về từng phần của đoạn mã:
*/

type RockClimber struct {
	rocksClimbed int          //Số lượng đá đã leo.
	kind         int          //Một thuộc tính bổ sung để định loại người leo núi
	sp           SafetyPlacer //Một đối tượng triển khai giao diện SafetyPlacer.
}

type SafetyPlacer interface {
	placeSafeties()
}

type IceSafetyPlacer struct {
	//db
	//data
	//api
}

// Hàm này tạo và trả về một con trỏ đến một đối tượng RockClimber mới, nhận vào một đối tượng triển khai giao diện SafetyPlacer.
func newRockClimber(sp SafetyPlacer) *RockClimber {
	return &RockClimber{
		sp: sp,
	}

}

type NOPSafetyPlace struct{}

func (sp NOPSafetyPlace) placeSafeties() {
	fmt.Println("place my safeties...")
}

func (rc *RockClimber) climbRock() {
	rc.rocksClimbed++

	if rc.rocksClimbed == 10 {
		rc.sp.placeSafeties()
	}
}

// func (sp *IceSafetyPlacer) placeSafeties() {
// 	fmt.Println("place my safeties...")
// 	// switch rc.kind {
// 	// case 1:
// 	// case 2:
// 	// case 3:

// 	// }
// }

func main() {
	rc := newRockClimber(&NOPSafetyPlace{})

	for i := 0; i < 11; i++ {
		rc.climbRock()
	}
}
