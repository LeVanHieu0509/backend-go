package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

// Trong Go, buffer được sử dụng để lưu trữ tạm thời dữ liệu trước khi xử lý hoặc truyền đi.
// Golang cung cấp nhiều cách để làm việc với buffer

func main() {
	// 1. bytes.Buffer là một kiểu dữ liệu giúp xử lý các chuỗi byte một cách hiệu quả,
	// tránh việc tạo nhiều chuỗi mới (string) gây tốn bộ nhớ.
	// Dễ dàng ghi dữ liệu nhiều lần mà không cần tạo chuỗi mới.
	bytesBuffer()

	//2. bufio.Writer và bufio.Reader
	bufioFunc()

	// 3. bytes.NewBuffer()
	bytesNewBuffer()

	// 4. Đọc file với bufio
	exampleReadLineFileTxt()

	//5. Đọc file theo từng chunk
	exampleReadChuckFileTxt()

	//6. Đọc file theo từng block
	exampleReadBlockFileTxt()

	//7. Đọc file lớn và ghi ra file CSV
	exampleReadFileExportToCSV()
}

func bytesBuffer() {
	// var buffer bytes.Buffer tạo một buffer rỗng.
	var b bytes.Buffer // A Buffer needs no initialization.
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "world!")
	b.WriteTo(os.Stdout)

	//Creating buffer variable to hold and manage the string data -> Ghi dữ liệu vào buffer
	var strBuffer bytes.Buffer
	strBuffer.WriteString("Ranjan")
	strBuffer.WriteString("Kumar")
	fmt.Println("The string buffer output is \n", strBuffer.String()) // Đọc dữ liệu từ buffer

	//Creating buffer variable to hold and manage the string data
	var byteString bytes.Buffer
	byteString.Write([]byte("Hello "))
	fmt.Fprintf(&byteString, "Hello friends how are you \n")
	byteString.WriteTo(os.Stdout)

	//Creating buffer variable to hold and manage the string data
	var strByyte bytes.Buffer
	strByyte.Grow(64)
	strByytestrByyte := strByyte.Bytes()
	strByyte.Write([]byte("It is a 64 byte"))
	fmt.Printf("%b \n", strByytestrByyte[:strByyte.Len()])
}

/*
- Giúp giảm số lần ghi vào file, tăng hiệu suất.
- Hữu ích khi làm việc với I/O.
*/
func bufioFunc() {
	// bufio cung cấp các công cụ để đọc và ghi dữ liệu với hiệu suất cao hơn,
	// đặc biệt khi làm việc với file hoặc dữ liệu lớn.

	file, _ := os.Create("example.txt")
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("Hello, Buffered Writer!\n")

	// Đẩy dữ liệu từ buffer xuống file
	writer.Flush()

	// bufio.Reader: Đọc dữ liệu từ buffer
	// Đọc dữ liệu từng dòng hoặc từng phần nhỏ, tránh chiếm quá nhiều bộ nhớ.
	data := "Hello, Golang!\nWelcome to buffering."
	reader := bufio.NewReader(strings.NewReader(data))

	line, _ := reader.ReadString('\n')
	fmt.Print(line) // Output: Hello, Golang!

}

func bytesNewBuffer() {
	// bytes.NewBuffer(data) tạo buffer từ slice []byte ban đầu.
	data := []byte("Hello, Buffer!")
	buffer := bytes.NewBuffer(data)

	fmt.Println(buffer.String()) // Output: Hello, Buffer!
}

func exampleReadLineFileTxt() {
	file, err := os.Open("example.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	// Tạo scan để đọc theo từng dòng
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Xử lý dữ liệu
		if count < 5 {
			fmt.Printf("line: %d:%s\n", count, line)
		}

		count++
	}

	// Kiểm tra lỗi khi đọc file
	if err := scanner.Err(); err != nil {
		log.Fatalf("Lỗi khi đọc file: %v", err)
	}

	fmt.Printf("Tổng số dòng đã đọc: %d\n", count)
}

func exampleReadChuckFileTxt() {
	// Mở file TXT
	file, err := os.Open("large_file_20M.txt")
	if err != nil {
		log.Fatalf("Không thể mở file: %v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	count := 0

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break // Thoát khi đọc hết file
		}

		// In ra 5 dòng đầu tiên
		if count < 5 {
			fmt.Println(string(line))
		}

		count++
	}

	fmt.Printf("Tổng số dòng đã đọc: %d\n", count)
}

func exampleReadBlockFileTxt() {
	// Mở file TXT
	file, err := os.Open("large_file_20M.txt")
	if err != nil {
		log.Fatalf("Không thể mở file: %v", err)
	}
	defer file.Close()

	// Tạo buffer đọc 64KB mỗi lần
	buf := make([]byte, 64*1024) // 64KB buffer
	totalBytes := 0

	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatalf("Lỗi khi đọc file: %v", err)
		}
		if n == 0 {
			break // Kết thúc khi đọc hết file
		}

		// Xử lý dữ liệu đọc được
		totalBytes += n
	}

	fmt.Printf("Tổng số bytes đã đọc: %d\n", totalBytes)
}

/*
- Đọc file lớn bằng Go routines và lưu vào CSV
*/
func exampleReadFileExportToCSV() {
	// Định nghĩa số lượng Goroutines chạy đồng thời
	const numWorkers = 4

	// Giới hạn số lượng CPU sử dụng để tối ưu hiệu suất
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Mở file đầu vào
	inputFile, err := os.Open("large_file_20M.txt")
	if err != nil {
		log.Fatalf("Không thể mở file: %v", err)
	}
	defer inputFile.Close()

	// Mở file CSV đầu ra
	outputFile, err := os.Create("output.csv")
	if err != nil {
		log.Fatalf("Không thể tạo file CSV: %v", err)
	}
	defer outputFile.Close()
	// Dùng csv.Writer để ghi dữ liệu ra file CSV theo luồng
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Tạo channel để truyền dữ liệu từ file sang worker
	lines := make(chan string, 100)
	var wg sync.WaitGroup

	// Khởi chạy worker goroutines để xử lý dữ liệu
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		// Dùng Goroutines + Channels để đọc file nhanh hơn.
		go processLines(lines, writer, &wg)
	}

	// Đọc file từng dòng và gửi vào channel
	// ✅ Dùng bufio.Scanner để đọc từng dòng, tránh load toàn bộ file vào RAM.
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		lines <- scanner.Text() // Gửi dữ liệu vào channel
	}
	close(lines) // Đóng channel sau khi đọc xong file

	// Đợi tất cả worker hoàn thành
	wg.Wait()

	// Kiểm tra lỗi nếu có
	if err := scanner.Err(); err != nil {
		log.Fatalf("Lỗi khi đọc file: %v", err)
	}

	fmt.Println("Hoàn tất xử lý và ghi file CSV!")
}

// Xử lý từng dòng và ghi vào file CSV
func processLines(lines chan string, writer *csv.Writer, wg *sync.WaitGroup) {
	defer wg.Done()

	// Dùng sync.WaitGroup để chờ tất cả worker hoàn thành.
	for line := range lines {
		// Chuyển đổi dữ liệu thành dạng CSV
		record := []string{line}

		// Ghi vào file CSV (cần mutex nếu nhiều goroutines cùng ghi)
		writer.Write(record)
		writer.Flush() // Flush để đảm bảo dữ liệu được ghi ngay
	}
}

/*
Khi nào sử dụng Buffer?
✅ Khi cần tối ưu hiệu suất đọc/ghi dữ liệu lớn.
✅ Khi cần tránh tạo nhiều chuỗi mới để giảm chi phí bộ nhớ.
✅ Khi làm việc với file, network, hoặc stream dữ liệu.

Bạn đang muốn dùng buffer trong trường hợp cụ thể nào? 🚀
*/
