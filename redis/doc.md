### Redis

1.  Standalone (6K) vs Sentinel (15K) vs Cluster
2.  Đặt redis trước mysql nên tốc độ gấp hàng trăm ngàn lần, nếu nhiều request quá thì lock lại trển redis chứ ko được để cho mysql bị die
3.  Sentinel: Chia ra 3 redis. R -> R -> W, phân tán dữ liệu ra.
4.  Cluster: Master - slave 01 - slave 02 -> Đồng bộ hoá dữ liệu slave để chia ra Đọc và Ghi.

root@bc78807762ae:/data# redis-cli INFO replication

# Replication

role:master
connected_slaves:2
slave0:ip=172.19.0.3,port=6379,state=online,offset=182,lag=0
slave1:ip=172.19.0.4,port=6379,state=online,offset=182,lag=0
master_replid:2070e14e12eb45f973495f6c67ea2bc93550dbd2
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:182
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:182

=> redis-cli INFO replication

- Nếu master ghi bao nhiều thì slave sẽ đọc được bấy nhiêu.
- Slave chỉ được đọc thôi chứ không được phép ghi.
- Nếu master chết thì sẽ không tự chuyển đổi được slave1 hoặc slave2 lên thay thế master mà phải nhờ đến sentinel

### SENTINEL

bind 0.0.0.0
port 26379

sentinel monitor redis-master redis-master 6379 2
sentinel down-after-milliseconds redis-master 1000

note: Mạng lan cục bộ của master slave -> redis-data_default

Mạng này chính là mạng trong sentinel -> đọc được ip và đè lên file config.
