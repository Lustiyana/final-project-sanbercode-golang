-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE comments (
  id SERIAL PRIMARY KEY,
  user_id INT,
  feed_id INT,
  message VARCHAR(256),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

-- +migrate StatementEnd