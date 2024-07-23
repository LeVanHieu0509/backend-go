### Khi nào nên sài Kafka

1. System truyền thống => Là cách lập trình nối tiếp, đồng bộ

- Nhược điểm:

* Các function có tỉ lệ phụ thuộc rất là cao, nếu lỗi 1 step là sai luôn cả flow
* Không nên để cho user đặt hàng không thành công
* Update cart bị tắc ngẽn or bị trễ thời gian rất lâu.
* Chờ handle đến step cuối cùng thì mới có kết quả
* Hệ thông có nhiều user order thì sẽ có nhược điểm

2. Kafka

- Tách hệ thống, loại bỏ lượng truy cập đồng thời cao

Hình 2:

- Producer là để write

* Là thành phần tạo ra các order và gửi chúng đến hàng đợi khác nhau
* là nguồn dữ liệu của tin nhắn
* Sau khi gửi thì sẽ cung cấp cho consumer

- Offset: Vị trí order trong hàng đợi
- Empty cuối hàng đợi: là phần bù và trỏ tới phần tử cuối cùng

Hình 3:

- Consumer là để read

* Là thành phần liên tục nhận các order => đọc những order đó để xử lý các công việc
* Hàng đợi đóng vài trò trung gian để kết nối giống như git (pull và push)
* 1 hàng đợi sẽ ko thể xử lý được nhiều lượng đồng thời cao => topic ra đời
* 1 topic chia ra nhiều partition - Tạo và sử dụng dữ liệu dựa vào kích thước phân vùng - nhiều partition có thể
  chạy song song mà ko ảnh hưởng lẫn nhau => nhiều chiến lược phân vùng => mở rộng dữ liệu
  Thứ tự partition ko đảm bảo thứ tự nhưng mà thứ tự trong partition sẽ được đảm bảo.

Hình 6:

- Brocker như là 1 đại lý, tất cả các brocker sẽ đồng bộ hoá siêu dữ liệu với nhau
- Quản lý các topic và các dữ liệu con

Hình 7, Hình 8

- Hệ thống đẹp nhất thì 5 parttion và 5 consumer -> 1 consumer còn lại là để hỗ trợ lỗi => khi nào có 1 consumer nào chết thì cái thứ 6 vào để thay thế
- Có thể mở rộng theo chiều ngang

- replication: Bản sao => nếu 1 container này lỗi thì lập tức có bản sao chép lên để chạy thay thế.
- Masters splace: Khi dữ liệu được sao chép bởi nhiều bản sao thì cần được sự đồng bộ hoá giữ các bản sao chính và bản sao phụ

* Lấy dữ liệu mới nhất ở master để đọc => ghi thì ghi ở master.
