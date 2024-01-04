-- +goose Up
-- +goose StatementBegin
CREATE TABLE templates (
	id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	content TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE templates;
-- +goose StatementEnd
