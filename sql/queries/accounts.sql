-- name: CreateAccount :one
INSERT INTO accounts 
(
    country,
    deleted,
    nickname,
    owner_id
)
VALUES 
(
    $1,
    FALSE,
    $2,
    $3
)
RETURNING
    id,
    created_at,
    updated_at,
    country,
    nickname,
    owner_id;
--

-- name: DeleteAccount :exec
UPDATE accounts
SET
    updated_at = $1,
    deleted = TRUE
WHERE id = $2 AND owner_id = $3;
--

-- name: GetAccount :one
SELECT
    id,
    created_at,
    updated_at,
    country,
    nickname,
    owner_id
FROM accounts
WHERE id = $1 AND owner_id = $2;
--

-- name: ListAccountsByOwnerId :many
SELECT
    id,
    created_at,
    updated_at,
    country,
    nickname,
    owner_id
FROM accounts
WHERE owner_id = $1
ORDER BY created_at DESC
LIMIT COALESCE(sqlc.narg('limit'), 10);
--

-- name: ListAccountsByUserId :many
SELECT
    accounts.id,
    accounts.created_at,
    accounts.updated_at,
    accounts.country,
    accounts.nickname,
    accounts.owner_id
FROM accounts
JOIN accounts_users 
ON accounts.id = accounts_users.account_id
WHERE accounts_users.user_id = $1
ORDER BY accounts.created_at DESC
LIMIT COALESCE(sqlc.narg('limit'), 10);;
--

-- name: UpdateAccount :one
UPDATE accounts
SET
    updated_at = sqlc.arg('updated_at')::timestamp,
    country = COALESCE(sqlc.narg('country'), country),
    nickname = COALESCE(sqlc.narg('nickname'), nickname)
WHERE id = sqlc.arg('id') AND owner_id = sqlc.arg('owner_id')
RETURNING
    id,
    created_at,
    updated_at,
    country,
    nickname,
    owner_id; 
--
