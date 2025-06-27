28. GO FRESHER (28): Tái thiết cấu trúc CODE từ MVC qua DDD Project (Source MVC - FEAT)
    link: https://drive.google.com/file/d/1OLH8J_h1AiLInBgr4eNI9rCOi_9PkMMM/view?usp=sharing

- Mô hình MVC
- Mô hình Feature - Layer (giống nestJs)
  1. Kiến trúc nhiều tính năng
  2. Chia dịch vụ, service
  3. Làm việc tầm 20 người, maintaining, scalable với module độc lập khác nhau
     Nhược điểm:
  4. Dễ dependency chéo
  5. Chưa hoàn hoàn tách biệt về domain logic
  6. Chưa rõ ràng về domain và application
- Mô hình DDD

  1. Là 1 trong những kiến trúc sau khi microvices được phát hành
  2. Phù hợp với team từ 50 người trở lên

  3. Domain: Business logic thuần
  4. Application: xử lý use case
  5. Infa: Thao tác I/O, cache, config, repo
  6. Controller: Entry point, http, grpc

- Mỗi folder chứa 4 (Domain, Application ,Infa, Controller)

1. Project là quán nhậu
