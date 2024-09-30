package main

import (
	"fmt"
	"math/big"
)

func main() {

	//var distance1 uint64 = 24e18  //Trường hợp này không thoả mãn do 24,000,000,000,000,000,000 > 9,223,372,036,854,775,807
	//fmt.Println("Alpha Centauri is", distance1, "km away.")

	//The big package provides three types:
	// big.Int is for big integers, when 18 quintillion isn’t enough. //dành cho số nguyên lớn, khi 18 triệu là không đủ
	// lightSpeed1 := big.NewInt(299792)
	// secondsPerDay1 := big.NewInt(86400)

	// distance := new(big.Int)
	// distance.SetString("24000000000000000000", 10)

	// big.Float is for arbitrary-precision floating-point numbers. //dành cho các số dấu phẩy động có độ chính xác tùy ý
	// big.Rat is for fractions like 1⁄3. //dành cho các phân số như 1⁄3.

	distance1 := new(big.Int) //khai báo và khởi tạo biến distance1 là một big.Int
	distance1.SetString("24000000000000000000", 10)
	fmt.Println("Andromeda Galaxy is", distance1, "km away.")

	lightSpeed1 := big.NewInt(299792)   //tốc độ ánh sáng trong kilômét mỗi giây
	secondsPerDay1 := big.NewInt(86400) //số giây trong một ngày

	// tạo secondsPerDay cách 2
	// secondsPerDay := new(big.Int)
	// secondsPerDay.SetString("86400", 10)

	seconds := new(big.Int)
	seconds.Div(distance1, lightSpeed1) //Chúng ta sử dụng biến seconds và days là các big.Int để tính toán thời gian cần thiết
	//để đi từ Trái Đất tới thiên hà Andromeda với tốc độ ánh sáng

	days := new(big.Int)
	days.Div(seconds, secondsPerDay1) //số ngày cần thiết để đi từ Trái Đất tới đó với tốc độ ánh sáng.

	fmt.Println("That is", days, "days of travel at light speed.")

	//-------------------------------------------------
	// Kết quả của phép tính được thực hiện ngay khi chương trình biên dịch. Máy tính sẽ thực hiện phép tính trên các giá trị trực tiếp.

	const distance = 24000000000000000000 //Việc tính toán các hằng số và hằng số được thực hiện trong quá trình biên dịch thay vì trong khi chương trình đang chạy
	fmt.Println("Andromeda Galaxy is", 24000000000000000000/299792/86400, "light days away.")

	// dùng hằng số
	// Đoạn mã thứ hai là cách viết tốt hơn vì nó mang lại tính rõ ràng, dễ bảo trì và tái sử dụng.
	// Mặc dù có vẻ dài hơn, nhưng nó giúp mã trở nên có cấu trúc và dễ hiểu hơn.
	const distance2 = 24000000000000000000
	const lightSpeed = 299792
	const secondsPerDay = 86400

	// Prints Andromeda Galaxy is 926568346 light days away.
	const days2 = distance2 / lightSpeed / secondsPerDay
	fmt.Println("Andromeda Galaxy is", days2, "light days away.")
}

// Andromeda Galaxy is 24000000000000000000 km away.
// That is 926568346 days of travel at light speed.
// Andromeda Galaxy is 926568346 light days away.
// Andromeda Galaxy is 926568346 light days away.
