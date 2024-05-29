-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE feeds (
  id SERIAL PRIMARY KEY,
  user_id INT,
  message VARCHAR(256),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

-- +migrate StatementEnd