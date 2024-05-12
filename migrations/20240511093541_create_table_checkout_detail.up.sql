CREATE TABLE checkout_details (
    id uuid PRIMARY KEY,
    checkout_id uuid NOT NULL,
    product_id uuid NOT NULL,
    quantity int NOT NULL,

    CONSTRAINT fk_checkout_detail_checkout_id FOREIGN KEY (checkout_id) REFERENCES checkouts(id),
    CONSTRAINT fk_checkout_detail_product_id FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE INDEX checkout_details_id ON checkout_details (id);
CREATE INDEX checkout_details_checkout_id ON checkout_details (checkout_id);
CREATE INDEX checkout_details_product_id ON checkout_details (product_id);