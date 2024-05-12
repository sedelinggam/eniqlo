CREATE TYPE product_category AS ENUM ('Clothing','Accessories','Footwear','Beverages');
CREATE TABLE products (
  id uuid PRIMARY KEY,
  name  VARCHAR(30),
  sku  VARCHAR(30),
  category product_category,
  image_url VARCHAR,
  notes VARCHAR,
  price BIGINT ,
  stock INT,
  location VARCHAR,
  is_avaiable BOOLEAN,
  created_at timestamptz NOT NULL
);
CREATE INDEX products_id ON products (id);
CREATE INDEX products_sku ON products (sku);