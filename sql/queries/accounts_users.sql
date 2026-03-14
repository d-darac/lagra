-- name: CreateAccountUserReference :exec
INSERT INTO accounts_users 
(
    account_id,
    user_id
)
VALUES 
(
    $1,
    $2
);