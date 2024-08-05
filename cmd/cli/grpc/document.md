comment:

- @aken2802
  Mình cũng đang triển khai cả 2 thằng này nếu có đủ nguồn lực và thời gian thì triển khai đúng kiến trúc như anh trình bày thì quá ngon rồi nhưng 1 vài trường hợp thời gian triển khai trên gRPC nó mất nhiều thời gian vì phải triển khai trển cả 2 đầu ( còn rest thì ko cần ) nên trường hợp nếu như ko yêu cầu tốc độ hoặc lượng data gọi qua lại ko quá lớn thì triển khai rest trong local server vẫn xài được bù lại tiết kiệm thời gian code cũng như chi phí nguồn lực coder
- Có lời khuyên là khi nào mà hệ thống có lượng traffic cao và đặc biệt là phát sinh bottleneck ở việc encode/decode request data giữa các services thì lúc đó hẳn tính đến việc xài GRPC hay không? Chứ bình thường thì xài REST vẫn tiện và ngon chán.
- Nên là cần phải có hệ thống monitoring, tracing đủ tốt để phát hiện ra chỗ nào cần tối ưu. Thấy có nhiều ông lên mạng đọc blog thấy tụi Big Tech như Google, Uber xài GRPC nghe nói nhanh thì cũng phải apply vào hệ thống cho bằng được trong khi chỗ làm chậm hệ thống có khi nằm ở Database hoặc code viết chưa tối ưu, rốt cuộc là chả cải thiện hệ thống lên bao nhiêu mà bắt cả team học thêm cái tech mới cùng với đống convention, best practice,... Chưa kể team DevOps/System cũng phải setup, tìm hiểu, cài thêm plugin để có thể monitor cái tech mới đó.

- gRPC phổ biến trong kiến trúc micro-service
  mà micro-service được triển khai để độc lập hóa giữa các module (cũng như là chia team ra quản lý luôn),
  mình thấy micro-service là scale về mặt nhân sự hơn là thay vì performance ( cũng có nhưng mà ko đáng kể )
