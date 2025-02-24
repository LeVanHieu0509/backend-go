dev_gateway:
	cd gateway && air

dev_orders:
	cd orders && air

dev_payments:
	cd payments && air

dev:
	# Chạy các dịch vụ trong nền
	$(MAKE) dev_gateway &
	$(MAKE) dev_orders &
	$(MAKE) dev_payments &

	# Đợi tất cả các tiến trình hoàn thành
	wait

kill_service:
	kill $(lsof -t -i :8080)
	kill $(lsof -t -i :8081)
	kill $(lsof -t -i :2000)
	kill $(lsof -t -i :2001)
