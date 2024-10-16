-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS files(
  id UUID PRIMARY KEY,
  document_id UUID,
  file_name varchar(255)
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS files;
-- +goose StatementEnd
