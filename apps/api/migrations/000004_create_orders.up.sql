CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    invoice_number TEXT NOT NULL,
    cashier_id BIGINT NOT NULL,
    total_amount NUMERIC(12,2) NOT NULL DEFAULT 0,
    status TEXT NOT NULL DEFAULT 'pending',

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_orders_cashier
        FOREIGN KEY (cashier_id)
        REFERENCES users(id)
);