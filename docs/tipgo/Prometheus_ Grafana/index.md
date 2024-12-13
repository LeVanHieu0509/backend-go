GO 40: Triển khai giám sát thời gian thực với

- Thu thập và xem những warining

### Prometheus

- Truy vấn hoá dữ liệu

- Golang (application), redis, mysql sẽ đẩy dữ liệu vào Prometheus => thu nhập dữ liệu đẩy về Grafana

### Grafana

- trực quan hoá dữ liệu khi app bắt đầu start -> nếu stop thì phải tính lại từ đầu

- http://localhost:3002/?orgId=1
- admin/ admin

1. connections: add data sources
   http://host.docker.internal:9093

2. Khám sức khoẻ cho application
   http://localhost:8001/metrics
   Dashboard: Send 5k request: ab -n 5000 -c 100 http://localhost:8001/ping/200

3. Sử dụng template để xem sức khoẻ của application

- grafana.com/grafana/dashboards/14061-go-runtime-metrics/
- import 14061 vào dashboard thì sẽ check được go routing
