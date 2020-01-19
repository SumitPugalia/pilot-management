-- +goose Up
-- +goose StatementBegin
CREATE TABLE pilots (
	id uuid PRIMARY KEY,
	user_id TEXT,
	supplier_id TEXT,
	market_id TEXT,
	service_id TEXT,
	code_name TEXT,
	state TEXT,
	created_at timestamp,
	updated_at timestamp,
	deleted boolean
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pilots;
-- +goose StatementEnd
