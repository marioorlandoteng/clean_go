CREATE DATABASE moneytransfer;

CREATE TABLE transaction (
  id SERIAL PRIMARY KEY,
  ref_id VARCHAR(255) NOT NULL,
  from_account_no BIGINT NOT NULL,
  to_account_no BIGINT NOT NULL,
  amount INT NOT NULL,
  status INT NOT NULL,
  created_at BIGINT NOT NULL, 
  updated_at BIGINT NOT NULL
);