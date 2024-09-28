-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS documents(
  id UUID PRIMARY KEY,
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS documents;
-- +goose StatementEnd
