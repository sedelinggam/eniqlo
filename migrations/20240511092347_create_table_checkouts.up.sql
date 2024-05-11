CREATE TABLE checkouts (
    id uuid PRIMARY KEY,
    customer_id uuid NOT NULL,
    paid bigint NOT NULL,
    change int NOT NULL,
    created_at timestamptz NOT NULL,

    CONSTRAINT fk_checkkout_customer_id FOREIGN KEY (customer_id) REFERENCES customers(id)
);