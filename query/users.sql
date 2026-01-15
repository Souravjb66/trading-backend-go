-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    password_hash,
    balance
)
VALUES ($1, $2, $3, $4)
RETURNING *;


-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;


-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = $1
  AND deleted_at IS NULL;


-- name: GetAllUsers :many
SELECT *
FROM users
WHERE deleted_at IS NULL;


-- name: UpdateUserBalanceDelta :one
UPDATE users
SET balance = balance + $1
WHERE id = $2
RETURNING *;

-- name: DebitUserBalance :one
UPDATE users
SET balance = balance - $1
WHERE id = $2
  AND balance >= $1
RETURNING *;

