-- +goose Up
-- +goose StatementBegin
CREATE TABLE templates (
	id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	name VARCHAR(255) NOT NULL UNIQUE,
	content TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE templates;
-- +goose StatementEnd
