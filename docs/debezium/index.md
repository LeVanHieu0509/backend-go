### New feeds

-- Khi thay đổi trong mysql thì nó sẽ lắng nghe sự thay đổi đó thì nó sẽ đẩy về kafka
-- Tốc độ realtime thời gian thực
-- Kafka sẽ đẩy về các service đang lắng nghe.

curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" localhost:8083/connectors/ -d @register-mysql.json
