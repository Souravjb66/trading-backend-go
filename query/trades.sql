-- name: InsertTrade :one
INSERT INTO trades (
    buy_order_id,
    sell_order_id,
    asset,
    price,
    quantity
)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;


-- name: GetTradesByUserID :many
SELECT t.*
FROM trades t
JOIN orders o
  ON t.buy_order_id = o.id
  OR t.sell_order_id = o.id
WHERE o.user_id = $1;

