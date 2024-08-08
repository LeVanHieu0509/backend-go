ur := repo.NewUserRepo()
us := service.NewUserService(ur)
userHandleNonDependency := controller.NewUserController(us)

---

### Trong interface: Để xử lý được nhược điểm là phải khai báo tất cả ràng buộc ra thì mới lấy được thằng chính.

Inversion of control: Là nguyên tắc thiết kết giảm sự kết nối giữa các thành phần phù thuộc với nhau
Dependency Injection: Thường áp dụng tính năng bảo trì code khiến cho những thành phần đó không phụ thuộc vào lẫn nhau

## Nếu các class sẽ yêu cầu 1 csdl thì cần phụ thuộc và khai báo nhiều lần.

### Inversion of control

- Là nguyên tắc trong lập trình hướng đối tượng

### Dependency Injection

- Là phương pháp phổ biến trong Inversion of control
