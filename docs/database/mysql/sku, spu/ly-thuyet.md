### Vấn đề

- 1 sản phẩm có nhiều sản phẩm con chỉ thực hiện được trên 1 bảng => sau này số lượng người dùng phát triển lên thì mới phát sinh ra 2 khái niệm này.

### SKU (standard product unit) là gì?

- SPU là 1 quy chuẩn của 1 đơn vị sản xuất, đề cập đến trong kho của chung ta có bao nhiêu sản phẩm mặc định

1. Iphone -> SPU

- SKU: color: red, blue, green
- SKU: Gb (64, 128, 256)

2. Giày Adidas -> SPU
3. Áo thun -> SPU

=> 1 SPU = N SKU

### SKU (Stock Keeping Unit)

- Table sd_product:

* Bảng này lưu trữ thông tin chung về một sản phẩm -> Đại diện cho 1 loại sản phẩm chung (1 Áo sơ mi với nhiều màu sắc)
* Cấp sản phẩm chính trước khi phân chia thành các SKU cụ thể.

- Table sku:

* Bảng này lưu trữ thông tin về từng SKU cụ thể.
* SKU thường là phiên bản cụ thể của một SPU (ví dụ: áo sơ mi màu đỏ, size M).
* Mỗi SKU là một phiên bản cụ thể của một sản phẩm, và có thể có nhiều SKU ứng với một SPU.

- Table sku_attr:

* Bảng này lưu trữ các thuộc tính cụ thể của từng SKU, chẳng hạn như các thông tin chi tiết của sản phẩm dưới dạng JSON.

- Table: spu_to_sku

* Bảng này lưu trữ mối quan hệ giữa SPU và SKU. Mỗi SPU có thể có nhiều SKU khác nhau, và bảng này giúp ánh xạ SPU với SKU tương ứng.
* Bảng này giúp liên kết SPU với các SKU tương ứng, cho phép quản lý tốt hơn các phiên bản khác nhau của một sản phẩm.

### Tổng quan

1. SPU (sd_product): Quản lý thông tin chung về sản phẩm.
2. SKU (sku): Quản lý từng phiên bản cụ thể của sản phẩm (mỗi SPU có nhiều SKU).
3. SKU Attributes (sku_attr): Lưu trữ thông tin chi tiết về từng SKU, giúp dễ dàng quản lý sản phẩm với nhiều thuộc tính khác nhau.
4. SPU to SKU (spu_to_sku): Liên kết giữa SPU và SKU, quản lý mối quan hệ nhiều SKU ứng với một SPU.
