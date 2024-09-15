### Tạo schema sử dụng goose.

goose -dir sql/schema create pre_go_crm_user_c sql

### Sau đó sẽ apply vào database tạo version

make upse
