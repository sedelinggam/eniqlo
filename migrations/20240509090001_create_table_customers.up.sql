CREATE TABLE customers (
  id uuid PRIMARY KEY, 
  phone_number VARCHAR(16)  UNIQUE NOT NULL, 
  name VARCHAR (50) NOT NULL, 
  created_at timestamptz NOT NULL
);

CREATE INDEX customers_phone_number ON customers (phone_number);