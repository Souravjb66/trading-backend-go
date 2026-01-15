-- name: CreateOrder :one
INSERT INTO orders (
    user_id,
    asset,
    type,
    price,
    quantity,
    remaining_quantity
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetOpenOrdersByAsset :many
SELECT *
FROM orders
WHERE asset = $1
  AND status = 'OPEN'
ORDER BY type DESC, price DESC, created_at ASC;


-- name: UpdateOrderStatus :one
UPDATE orders
SET remaining_quantity = $1,
    status = $2
WHERE id = $3
RETURNING *;



-- name: GetHighestPriceBuyer :many
SELECT *
FROM orders
WHERE type = 'BUY'
  AND status = 'OPEN'
  AND asset = $1
  AND price >= $2
ORDER BY price DESC;


-- name: GetOrdersByOrderId :one
SELECT *
FROM orders
WHERE id = $1
  AND status = 'OPEN';


-- name: GetMaxBuy :many
SELECT asset, MAX(price) AS max_buy_price
FROM orders
WHERE type = 'BUY'
  AND status IN ('OPEN', 'PARTIAL')
GROUP BY asset;


-- name: GetMaxBuyPriceOfaAsset :one
SELECT id,
       asset,
       user_id,
       price AS max_buy_price,
       quantity,
       remaining_quantity,
       status,
       created_at
FROM orders
WHERE type = 'BUY'
  AND status IN ('OPEN', 'PARTIAL')
  AND asset = $1
ORDER BY price DESC
LIMIT 1;


-- name: GetOpenBuyOrders :many
SELECT * FROM orders
WHERE type = 'buy' 
  AND remaining_quantity > 0
ORDER BY price DESC, created_at ASC;

-- name: GetOpenSellOrders :many
SELECT * FROM orders
WHERE type = 'sell' 
  AND remaining_quantity > 0
ORDER BY price ASC, created_at ASC;

