-- +goose Up
-- +goose StatementBegin
CREATE TABLE emails (
	id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	template_name VARCHAR(255) NOT NULL,
	recipients VARCHAR(255)[] NOT NULL CHECK (cardinality(recipients) > 0),
	subject VARCHAR(255) NOT NULL,
	context jsonb
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE emails;
-- +goose StatementEnd
