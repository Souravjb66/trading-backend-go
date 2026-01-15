CREATE TABLE portfolio (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    asset VARCHAR(20) NOT NULL,
    quantity BIGINT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_portfolio_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,
    CONSTRAINT unique_user_asset UNIQUE (user_id, asset)
);
