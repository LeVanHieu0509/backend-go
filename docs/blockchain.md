# Lợi ích khi sử dụng kiến trúc BlockChain

- Khả năng hoàn tất giao dịch nhanh
- Giảm chi phí cho doanh nghiệp, loại bỏ trung gian
- Xóa bỏ gian lận, tấn công mạng hoặc các tội phạm điện tử khác
- minh bạch và độ rõ ràng tối đa

# Khái niệm

1. mọi người sử dụng nó đều có một bản sao đầy đủ hoặc một phần của nó. Và một bản ghi mới chỉ có thể được thêm vào khi có sự đồng ý của những người giữ cơ sở dữ liệu khác. Ngoài ra, chính blockchain đã tạo ra tiền điện tử và hợp đồng thông minh

### Khối và chuỗi

- tính toán băm là một hoạt động khó khăn về mặt tính toán, mất một thời gian ngay cả trên máy tính nhanh (đó là lý do tại sao mọi người mua GPU mạnh để khai thác Bitcoin)
- Chain: là một danh sách được sắp xếp theo thứ tự, được liên kết ngược
- Đào bitcoin: Trong blockchain, một số người tham gia (thợ đào) của mạng lưới làm việc để duy trì mạng lưới, thêm các khối mới vào đó và nhận phần thưởng cho công việc của họ
- Thuật toán Proof-of-Work: thực hiện công việc thì khó, nhưng xác minh bằng chứng thì dễ. Bằng chứng thường được giao cho người khác, vì vậy đối với họ, việc xác minh bằng chứng không mất nhiều thời gian.

- Hàm HASH:

* Không thể khôi phục dữ liệu gốc từ hàm băm
* Một số dữ liệu chỉ có thể có một giá trị băm và giá trị băm đó là duy nhất
* Chỉ cần thay đổi một byte trong dữ liệu đầu vào cũng sẽ tạo ra một giá trị băm hoàn toàn khác.

## Lưu trữ

- blocks: lưu trữ siêu dữ liệu mô tả tất cả các khối trong một chuỗi.
- chainstate: lưu trữ trạng thái của chuỗi, bao gồm tất cả các đầu ra giao dịch chưa sử dụng và một số siêu dữ liệu.

## Transaction

- 1. Giao dịch là cốt lõi của Bitcoin và mục đích duy nhất của blockchain là lưu trữ giao dịch theo cách an toàn và đáng tin cậy
- 2. Đầu vào của một giao dịch mới tham chiếu đến đầu ra của một giao dịch trước đó

* Có những đầu ra không liên quan đến đầu vào.
* Trong một giao dịch, đầu vào có thể tham chiếu đến đầu ra từ nhiều giao dịch.
* Đầu vào phải tham chiếu đến đầu ra.

### Đào BITCOIN

1.  Khi một thợ đào bắt đầu đào một khối, họ sẽ thêm một giao dịch coinbase vào khối đó
2.  Giao dịch coinbase là một loại giao dịch đặc biệt, không yêu cầu các đầu ra đã tồn tại trước đó
3.  Đây là phần thưởng mà thợ đào nhận được khi đào các khối mới.

### Bitcoin Halving

- Khai thác khối genesis tạo ra 50 BTC và mỗi 210000 khối phần thưởng sẽ giảm một nửa
- mỗi khối phải lưu trữ ít nhất một giao dịch và không thể khai thác khối mà không có giao dịch nữa
- tất cả các giao dịch trong một khối được xác định duy nhất bằng một băm duy nhất (Để đạt được điều này, chúng tôi lấy các băm của mỗi giao dịch, nối chúng lại và lấy một băm của tổ hợp đã nối.)
