### 2. Cache breakdown

- Khi mà có voucher bán iphone 0 đồng
- B1: Đưa dữ liệu vào redis trước
- B2: Nếu có số lượng lớn request vào để mua iphone đó thì phải thông qua redis để loại bỏ bớt request xuống database
- B3: Redis chỉ nhận 5000 request nhưng mà user quá lớn nên request tận 6000 thì database sẽ bị ảnh hưởng.

=> 3 cách:

1. Để expired
2. Gia hạn cache theo cronjob vào lúc 12g đêm
3. Thêm khóa mutex (chỉ 1 thằng lọt vô để lấy csdl để lưu data vào cache => 5999 request còn lại chờ để lấy cache) - đặt cặp key và value

### GO SENIOR (03) : REDIS Solutions: Standalone vs Sentinel vs Cluster

- Mysql chỉ chịu đựng được 3k request/s => redis sẽ gánh được request -> lưu trên memory
- Nếu sử dụng redis sentinal thì trên 3 con redis sẽ chỉ làm 1 nhiệm vụ duy nhất ĐỌC hoặc GHI
-

sử dụng: wrk -t12 -c1000 -d30s http://localhost:8001/v1/2024/ticket/item/1
