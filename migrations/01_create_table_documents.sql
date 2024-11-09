-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS documents(
  id UUID PRIMARY KEY,
  creator_id UUID NOT NULL,
  date DATE NOT NULL,
  approve BOOLEAN DEFAULT false
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS documents;
-- +goose StatementEnd
