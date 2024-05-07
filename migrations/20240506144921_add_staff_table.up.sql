CREATE TABLE staffs (
  id uuid PRIMARY KEY, 
  phone_number VARCHAR  UNIQUE NOT NULL, 
  password VARCHAR (72) NOT NULL,
  name VARCHAR (50) NOT NULL, 
  created_at timestamptz NOT NULL
);

CREATE INDEX staffs_id ON staffs (id);
CREATE INDEX staffs_phone_number ON staffs (phone_number);