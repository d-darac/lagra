-- name: CreateApiKey :one
INSERT INTO api_keys
(
    expires_at,
    name,
    note,
    redacted_secret,
    secret,
    account_id
)
VALUES
(
    COALESCE(sqlc.narg('expires_at')::timestamp, NULL),
    sqlc.arg('name'),
    sqlc.arg('note'),
    sqlc.arg('redacted_secret'),
    sqlc.arg('secret'),
    sqlc.arg('account_id')
)
RETURNING
    id,
    created_at,
    updated_at,
    expires_at,
    name,
    note,
    redacted_secret;
--

-- name: DeleteApiKey :exec
DELETE FROM api_keys
WHERE id = $1 AND account_id = $1;
--

-- name: ExpireApiKey :exec
UPDATE api_keys
SET
    updated_at = sqlc.arg('updated_at')::timestamp,
    expires_at = sqlc.arg('expires_at')::timestamp
WHERE id = sqlc.arg('id') AND account_id = sqlc.arg('account_id');
--

-- name: GetApiKey :one
SELECT
    id,
    created_at,
    updated_at,
    expires_at,
    name,
    note,
    redacted_secret
FROM api_keys
WHERE id = $1 AND account_id = $2;
--

-- name: GetApiKeyAccAndExp :one
SELECT
    expires_at,
    account_id
FROM api_keys
WHERE secret = $1;
--

-- name: ListApiKeys :many
SELECT *
FROM
(
    SELECT
        api_keys.id,
        api_keys.created_at,
        api_keys.updated_at,
        api_keys.expires_at,
        api_keys.name,
        api_keys.note,
        api_keys.redacted_secret
    FROM api_keys
    WHERE api_keys.account_id = sqlc.arg('account_id')
    AND
        (
            sqlc.narg('created_at_gt')::timestamp IS NULL 
            OR api_keys.created_at > sqlc.narg('created_at_gt')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_lt')::timestamp IS NULL 
            OR api_keys.created_at < sqlc.narg('created_at_lt')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_gte')::timestamp IS NULL 
            OR api_keys.created_at >= sqlc.narg('created_at_gte')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_lte')::timestamp IS NULL 
            OR api_keys.created_at <= sqlc.narg('created_at_lte')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_gt')::timestamp IS NULL 
            OR api_keys.updated_at > sqlc.narg('updated_at_gt')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_lt')::timestamp IS NULL 
            OR api_keys.updated_at < sqlc.narg('updated_at_lt')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_gte')::timestamp IS NULL 
            OR api_keys.updated_at >= sqlc.narg('updated_at_gte')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_lte')::timestamp IS NULL 
            OR api_keys.updated_at <= sqlc.narg('updated_at_lte')::timestamp
        )
    AND 
        (
            sqlc.narg('starting_after')::uuid IS NULL
            OR 
            (
                (
                    sqlc.narg('starting_after_date')::timestamp,
                    sqlc.narg('starting_after')::uuid
                ) > (api_keys.created_at, api_keys.id)
            )
        )
    AND 
        (
            sqlc.narg('ending_before')::uuid IS NULL
            OR 
            (
                (
                    sqlc.narg('ending_before_date')::timestamp,
                    sqlc.narg('ending_before')::uuid
                ) < (api_keys.created_at, api_keys.id)
            )
        )
    AND 
        (
            sqlc.narg('name')::text IS NULL 
            OR api_keys.name ~~* CONCAT('%', sqlc.narg('name')::text, '%')
        )
    ORDER BY 
    (
        CASE 
            WHEN sqlc.narg('ending_before')::uuid IS NOT NULL 
            THEN (api_keys.created_at, api_keys.id)
        END
    ) ASC,
    (api_keys.created_at, api_keys.id) DESC
    LIMIT COALESCE(sqlc.narg('limit')::integer, 10) + 1
)
ORDER BY created_at DESC, id DESC;
--

-- name: UpdateApiKey :one
UPDATE api_keys
SET
    updated_at = sqlc.arg('updated_at')::timestamp,
    name = COALESCE(sqlc.narg('name'), name),
    note = COALESCE(sqlc.narg('note'), note)
WHERE id = sqlc.arg('id') AND account_id = sqlc.arg('account_id')
RETURNING
    id,
    created_at,
    updated_at,
    expires_at,
    name,
    note,
    redacted_secret;
--
