- là 1 tool comment để quản lý csdl được viết bằng go
- Quản lý phiên bản cơ sở dữ liệu
- Ví dụ cần phải sửa database và cần phải làm nhiều môi trường => quy chuẩn về database

## help

up: Di chuyển DB sang phiên bản mới nhất có sẵn
up-by-one: Di chuyển DB lên 1
up-to VERSION: Di chuyển DB sang một PHIÊN BẢN cụ thể
down: Quay lại phiên bản 1
down-to VERSION: Quay lại một PHIÊN BẢN cụ thể
redo: Chạy lại lần di chuyển mới nhất
reset: Quay lại tất cả các lần di chuyển
status Dump: trạng thái di chuyển cho DB hiện tại
version In: phiên bản hiện tại của cơ sở dữ liệu
create NAME: [sql|go] Tạo tệp di chuyển mới với dấu thời gian hiện tại
fix: Áp dụng thứ tự tuần tự cho các lần di chuyển
validate: Kiểm tra các tệp di chuyển mà không chạy chúng

---

goose create order sql: tên tệp được tự động tạo ra và tạo ra phiên bản cơ sở dữ liệu
SELECT 'up SQL query'; - Biểu hiện cho sự di chuyển đồng bộ dữ liệu
SELECT 'down SQL query'; - Khôi phục

cmd: goose mysql "levanhieu:levanhieu1234@tcp(127.0.0.1:33060)/shopdevgo" up
2024/08/22 16:56:33 OK 20240822094441_order.sql (23.82ms)
2024/08/22 16:56:33 goose: successfully migrated database to version: 20240822094441

### execute database

use shopdevgo;

show tables;
DESC orders;
select \* from goose_db_version;

### Quản lý version giúp version giúp lặp lại
