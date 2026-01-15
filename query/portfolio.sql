-- name: GetPortfolioByUserID :many
SELECT *
FROM portfolio
WHERE user_id = $1;

-- name: GetPortfolioByUserIdAndAsset :one
SELECT *
FROM portfolio
WHERE user_id = $1 AND asset=$2;

-- name: UpdatePortfolioBalance :one
UPDATE portfolio
SET quantity = $1,
    updated_at = CURRENT_TIMESTAMP
WHERE user_id = $2
  AND asset = $3
RETURNING *;

-- name: UpdatePortfolioDebitQuantity :one
UPDATE portfolio
SET quantity = quantity-$1,
    updated_at = CURRENT_TIMESTAMP
WHERE user_id = $2
  AND asset = $3
RETURNING *;

-- name: UpdatePortfolioCreditQuantity :one
UPDATE portfolio
SET quantity = quantity+$1,
    updated_at = CURRENT_TIMESTAMP
WHERE user_id = $2
  AND asset = $3
RETURNING *;

-- name: InsertPortfolio :one
INSERT INTO portfolio (user_id, asset, quantity)
VALUES ($1, $2, $3)
RETURNING *;

