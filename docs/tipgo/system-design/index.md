1. Lý thuyết về một số câu hỏi của design parttern

- 15 sự đánh đổi

### Scalability

- Khả năng mở rộng: về tương lai, xác định mở rộng như thế nào sau một thời gian, để xử lý lượng dữ liệu tăng lên theo thời gian hay không, sử dụng cluster redis

- redis nốt đơn -> cluster redis
- Tính đơn giản -> phức tạp -> làm giảm hiệu suất
- Hy sinh khả năng mở rộng để làm tăng hiệu suất: Thương mại điện tử, event,

### Performance

- Hiệu suất đề cập đến tốc độ phản hồi của hệ thống, 1 request => system trả lời hết bao nhiêu giây
- Hy sinh hiệu suất để làm tăng khả năng mở rộng: Giao dịch chứng khoán

- Mở rộng theo chiều dọc: -> 2gb -> 4gb -> 6gb
  Ưu điểm: đơn giản
  Nhược điểm: Là core duy nhất, nếu server chết là die hết

- Mở rộng theo chiều ngang: Cluster (3 server - 2gb)
  Ưu điểm: Khả năng mở rộng vô đối
  Nhược điểm: Dữ liệu bị phân tán, làm sao để quản lý được dữ liệu nhất quán

  -> Mới startup -> chiều dọc (fullstack)
  -> Khi data người dùng nhiều -> chiều ngang (devopps)

### Latency

- Độ trễ của server gửi về cho client, được tính theo thời gian (ms) kể từ khi yêu cầu đến khi phản hồi
- Khi latency càng thấp thì tốc độ phản hồi càng nhanh
- Realtime: 0ms (chơi game)

### Throughput

- Khi client gửi request tới server: số lượng request gọi vào server trong 1 khoảng thời gian nhất định (req/s)
- Phản ánh năng lực xử lý chung
- Nếu Throughput càng nhiều -> Latency càng cao (vì chiếm CPU và Băng thông, ram cao)
- Youtobe: Đảm bảo tất cả người dùng xem cùng 1 lúc

### Relational (Mysql) - Phát triển theo chiều dọc

- Cơ sở dữ liệu quan hệ

### No SQL (MongoDb) - Phát triển theo chiều ngang

- Cơ sở dữ liệu không có quan hệ, phi quan hệ, mở rộng rất dễ dàng
- Social network (Bình luận, like, ...)

### CAP

- Ít có người phụ nữ nào vừa đẹp mà vừa giỏi, chọn giữ tính nhất quán và tính khả dụng

- Consistency: Mỗi lần đọc dữ liệu đều nhận được dữ liệu mới nhất
- Availability: Mỗi yêu cầu (đọc hoặc ghi) sẽ nhận được một phản hồi, dù là thành công hay thất bại, ngay cả khi một số nút không khả dụng.
- Partition Tolerance: Hệ thống có thể tiếp tục hoạt động ngay cả khi có sự phân vùng mạng (mất kết nối giữa các nút).

* CA: Ngân hàng, A chuyển tiền cho B -> B phải nhận được tiền
* AP: Social Network, Ecommerce, Data có thể sai nhưng sẽ ko ảnh hưởng nghiêm trọng, có thể xử lý lại được
* CP: Cơ sở dữ liệu phân tán – Một hệ thống cơ sở dữ liệu phân tán có thể ưu tiên nhất quán trong dữ liệu ngay cả khi có sự cố phân vùng mạng, nhưng đôi khi yêu cầu có thể không được xử lý trong tình huống mạng bị chia cắt.

### Strong consistency (nhất quán mạnh)

- Sử dụng trong hệ thống phân tán, khi 1 record được cập nhật , thì sự thay đổi này phải có hiệu lực ngay lập tức và đồng bộ trên tất cả các server
- Ứng dụng: Được sử dụng trong các hệ thống yêu cầu tính chính xác tuyệt đối trong mỗi giao dịch, như hệ thống tài chính hay ngân hàng.

### Week consistency (nhất quán yếu)

- Không yêu cầu cập nhật bản ghi ngay lập tức trên tất cả các server. Dữ liệu có thể bị mất đồng bộ trong một khoảng thời gian ngắn nhưng vẫn đảm bảo hệ thống sẽ "sửa chữa" khi có thể.
- Ứng dụng: Các ứng dụng như mạng xã hội hoặc thương mại điện tử, nơi sự đồng bộ dữ liệu không phải là yếu tố quan trọng ngay lập tức, và các lỗi tạm thời (như thông tin bị lỗi trong một vài giây) có thể được xử lý sau.

### Cache

- Chiến lược đọc cache:
  Đọc trực tiếp từ cache: Khi một client yêu cầu dữ liệu, hệ thống kiểm tra cache trước. Nếu dữ liệu có sẵn trong cache, thì sẽ trả về dữ liệu nhanh chóng mà không cần phải truy vấn vào cơ sở dữ liệu (MySQL). Nếu không có trong cache, dữ liệu sẽ được lấy từ cơ sở dữ liệu và lưu vào cache để sử dụng lần sau.
- Chiến lược ghi cache: client -> server -> cache -> mysql -> ack -> cache -> ack -> server -> client
  - Lưu trữ và đồng bộ cache: Quá trình ghi vào cache bao gồm các bước:
    1. Client gửi yêu cầu ghi dữ liệu vào server.
    2. Server ghi vào cache và cơ sở dữ liệu (MySQL).
    3. Sau khi ghi vào cơ sở dữ liệu, một phản hồi thành công (ack) sẽ được gửi về cache, sau đó phản hồi về client.
    4. Cache và cơ sở dữ liệu cần phải đồng bộ, đảm bảo rằng mọi thay đổi được ghi chính xác.
  - Với 1 luồng thì chiến lược này có thể ổn. Tuy nhiên, khi có nhiều luồng xử lý cùng lúc, sự đồng bộ dữ liệu giữa các luồng có thể gây ra tình trạng "Nhất quán yếu", bởi vì nhiều yêu cầu ghi có thể dẫn đến việc dữ liệu không đồng bộ hoặc bị xung đột trong quá trình đồng bộ giữa các server và cache.
