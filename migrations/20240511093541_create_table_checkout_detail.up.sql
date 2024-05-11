CREATE TABLE checkout_details (
    id uuid PRIMARY KEY,
    checkout_id uuid NOT NULL,
    product_id uuid NOT NULL,
    quantity int NOT NULL,

    CONSTRAINT fk_checkout_detail_checkout_id FOREIGN KEY (checkout_id) REFERENCES checkouts(id),
    CONSTRAINT fk_checkout_detail_product_id FOREIGN KEY (product_id) REFERENCES products(id)
);