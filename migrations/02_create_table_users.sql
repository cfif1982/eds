-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
  id UUID PRIMARY KEY,
  email varchar(20) NOT NULL,
  name varchar(255),
  telephone varchar(11),
  open_key varchar(255)
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
