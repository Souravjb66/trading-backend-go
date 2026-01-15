CREATE TABLE trades (
    id BIGSERIAL PRIMARY KEY,
    buy_order_id BIGINT NOT NULL,
    sell_order_id BIGINT NOT NULL,
    asset VARCHAR(20) NOT NULL,
    price NUMERIC(20,8) NOT NULL,
    quantity NUMERIC(20,8) NOT NULL,
    executed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_trades_buy_order
        FOREIGN KEY (buy_order_id)
        REFERENCES orders(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_trades_sell_order
        FOREIGN KEY (sell_order_id)
        REFERENCES orders(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_trades_asset_time
ON trades(asset, executed_at);