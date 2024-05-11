CREATE DATABASE moneytransfer;

CREATE TABLE transaction (
  id SERIAL PRIMARY KEY,
  from_user_id INT NOT NULL,
  to_user_id INT NOT NULL,
  amount INT NOT NULL,
  status VARCHAR(32) NOT NULL,
  created_at TIMESTAMP NOT NULL, 
  updated_at TIMESTAMP NOT NULL
);