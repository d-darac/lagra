-- name: CreateAccountUserReference :exec
INSERT INTO accounts_users 
(
    created_at,
    updated_at,
    account_id,
    user_id
)
VALUES 
(
    $1,
    $2,
    $3,
    $4
);