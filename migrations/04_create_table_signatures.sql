-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS signatures(
  id UUID PRIMARY KEY,
  file_id UUID,
  signer_id UUID,
  signature_file_name varchar(255),
  date DATE,
  delete BOOLEAN DEFAULT false
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS signatures;
-- +goose StatementEnd
