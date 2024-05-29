-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(256),
  email VARCHAR(256),
  password VARCHAR(256)
);

-- +migrate StatementEnd