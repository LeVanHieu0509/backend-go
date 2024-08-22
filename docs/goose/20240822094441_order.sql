-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
                orderId int PRIMARY KEY,
                userId int NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
