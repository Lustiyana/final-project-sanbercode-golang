-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE likes (
  id SERIAL PRIMARY KEY,
  user_id INT,
  feed_id INT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

-- +migrate StatementEnd