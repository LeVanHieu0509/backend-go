### Khái niệm

1. Go cho phép chúng ta chạy mã đồng thời bằng goroutine.
2. Tuy nhiên, khi các tiến trình đồng thời truy cập cùng một phần dữ liệu, nó có thể dẫn đến tình trạng chạy đua.
3. Mutex là các cấu trúc dữ liệu được cung cấp bởi gói sync. Chúng có thể giúp chúng ta đặt khóa trên các phần dữ liệu khác nhau để chỉ một goroutine có thể truy cập vào nó tại một thời điểm.
