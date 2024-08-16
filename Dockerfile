#  Bạn sử dụng một image Golang để xây dựng ứng dụng.
FROM golang:alpine as builder 

# Tạo thư mục /build trong container và đặt nó làm thư mục làm việc hiện tại.
WORKDIR /build

# Copy toàn bộ mã nguồn từ thư mục cục bộ vào thư mục /build trong container.
COPY . .

# ải về tất cả các module Go cần thiết
RUN go mod download

# Build file nhị phân của ứng dụng từ file main.go.
RUN go build -o crm.shopdev.com cmd/server/main.go

# Tạo một container mới từ scratch, nghĩa là không có hệ điều hành cơ bản, chỉ chứa các thành phần cần thiết.
FROM scratch

# Copy thư mục configs từ máy cục bộ vào /config trong container.
COPY configs /configs

COPY --from=builder /build/crm.shopdev.com /

# Chỉ định lệnh để chạy ứng dụng, với tham số là đường dẫn đến file cấu hình.
CMD ["/crm.shopdev.com", "/configs/local.yaml"]