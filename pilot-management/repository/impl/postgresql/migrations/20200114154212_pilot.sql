-- +goose Up
-- +goose StatementBegin
CREATE TABLE pilots (id UUID PRIMARY KEY,
	user_id VARCHAR,
	supplier_id VARCHAR,
	market_id VARCHAR,
	service_id VARCHAR,
	code_name VARCHAR,
	created_at BIGINT DEFAULT 0,
	updated_at BIGINT DEFAULT 0,
	deleted_at BIGINT DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pilots;
-- +goose StatementEnd
