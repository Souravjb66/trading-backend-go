-- 1️⃣ Create ENUMs first
CREATE TYPE order_type AS ENUM ('BUY', 'SELL');
CREATE TYPE order_status AS ENUM ('OPEN', 'PARTIAL', 'FILLED', 'CANCELLED');


-- 2️⃣ Now create table
CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    asset VARCHAR(20) NOT NULL,
    type order_type NOT NULL,
    price NUMERIC(20,8) NOT NULL,
    quantity NUMERIC(20,8) NOT NULL,
    remaining_quantity NUMERIC(20,8) NOT NULL,
    status order_status NOT NULL DEFAULT 'OPEN',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_orders_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);


-- 3️⃣ Indexes
CREATE INDEX idx_orders_asset_type_price
ON orders(asset, type, price);

CREATE INDEX idx_orders_status
ON orders(status);
