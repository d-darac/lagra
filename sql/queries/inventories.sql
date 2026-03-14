-- name: CreateInventory :one
INSERT INTO inventories 
(
    in_stock,
    orderable,
    account_id
)
VALUES 
(
    $1,
    $2,
    $3
)
RETURNING
    id,
    created_at,
    updated_at,
    in_stock,
    orderable,
    reserved;
--

-- name: DeleteInventory :exec
DELETE FROM inventories
WHERE id = $1 AND account_id = $2;
--

-- name: GetInventory :one
SELECT
    id,
    created_at,
    updated_at,
    in_stock,
    orderable,
    reserved
FROM inventories
WHERE id = $1 AND account_id = $2;
--

-- name: ListInventories :many
SELECT *
FROM
(
    SELECT
        inventories.id,
        inventories.created_at,
        inventories.updated_at,
        inventories.in_stock,
        inventories.orderable,
        inventories.reserved
    FROM inventories
    WHERE inventories.account_id = sqlc.arg('account_id')
    AND
        (
            sqlc.narg('created_at_gt')::timestamp IS NULL 
            OR inventories.created_at > sqlc.narg('created_at_gt')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_lt')::timestamp IS NULL 
            OR inventories.created_at < sqlc.narg('created_at_lt')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_gte')::timestamp IS NULL 
            OR inventories.created_at >= sqlc.narg('created_at_gte')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_lte')::timestamp IS NULL 
            OR inventories.created_at <= sqlc.narg('created_at_lte')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_gt')::timestamp IS NULL 
            OR inventories.updated_at > sqlc.narg('updated_at_gt')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_lt')::timestamp IS NULL 
            OR inventories.updated_at < sqlc.narg('updated_at_lt')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_gte')::timestamp IS NULL 
            OR inventories.updated_at >= sqlc.narg('updated_at_gte')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_lte')::timestamp IS NULL 
            OR inventories.updated_at <= sqlc.narg('updated_at_lte')::timestamp
        )
    AND 
        (
            sqlc.narg('starting_after')::uuid IS NULL
            OR 
            (
                (
                    sqlc.narg('starting_after_date')::timestamp,
                    sqlc.narg('starting_after')::uuid
                ) > (inventories.created_at, inventories.id)
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
                ) < (inventories.created_at, inventories.id)
            )
        )
    AND
        (
            sqlc.narg('in_stock_gt')::integer IS NULL 
            OR inventories.in_stock > sqlc.narg('in_stock_gt')::integer
        )
    AND 
        (
            sqlc.narg('in_stock_lt')::integer IS NULL 
            OR inventories.in_stock < sqlc.narg('in_stock_lt')::integer
        )
    AND 
        (
            sqlc.narg('in_stock_gte')::integer IS NULL 
            OR inventories.in_stock >= sqlc.narg('in_stock_gte')::integer
        )
    AND 
        (
            sqlc.narg('in_stock_lte')::integer IS NULL 
            OR inventories.in_stock <= sqlc.narg('in_stock_lte')::integer
        )
    AND
        (
            sqlc.narg('orderable_gt')::integer IS NULL 
            OR inventories.orderable > sqlc.narg('orderable_gt')::integer
        )
    AND 
        (
            sqlc.narg('orderable_lt')::integer IS NULL 
            OR inventories.orderable < sqlc.narg('orderable_lt')::integer
        )
    AND 
        (
            sqlc.narg('orderable_gte')::integer IS NULL 
            OR inventories.orderable >= sqlc.narg('orderable_gte')::integer
        )
    AND 
        (
            sqlc.narg('orderable_lte')::integer IS NULL 
            OR inventories.orderable <= sqlc.narg('orderable_lte')::integer
        )
    AND
        (
            sqlc.narg('reserved_gt')::integer IS NULL 
            OR inventories.reserved > sqlc.narg('reserved_gt')::integer
        )
    AND 
        (
            sqlc.narg('reserved_lt')::integer IS NULL 
            OR inventories.reserved < sqlc.narg('reserved_lt')::integer
        )
    AND 
        (
            sqlc.narg('reserved_gte')::integer IS NULL 
            OR inventories.reserved >= sqlc.narg('reserved_gte')::integer
        )
    AND 
        (
            sqlc.narg('reserved_lte')::integer IS NULL 
            OR inventories.reserved <= sqlc.narg('reserved_lte')::integer
        )
    ORDER BY 
    (
        CASE 
            WHEN sqlc.narg('ending_before')::uuid IS NOT NULL 
            THEN (inventories.created_at, inventories.id)
        END
    ) ASC,
    (inventories.created_at, inventories.id) DESC
    LIMIT COALESCE(sqlc.narg('limit')::integer, 10) + 1
)
ORDER BY created_at DESC, id DESC;
--

-- name: ListInventoriesByIds :many
SELECT
    id,
    created_at,
    updated_at,
    in_stock,
    orderable,
    reserved
FROM inventories
WHERE account_id = sqlc.arg('account_id')
    AND id = ANY(sqlc.arg('ids')::uuid[])
ORDER BY created_at DESC, id DESC;
--

-- name: UpdateInventory :one
UPDATE inventories
SET
    updated_at = sqlc.arg('updated_at')::timestamp,
    in_stock = COALESCE(sqlc.narg('in_stock'), in_stock),
    orderable = COALESCE(sqlc.narg('orderable'), orderable),
    reserved = COALESCE(sqlc.narg('reserved'), reserved)
WHERE id = sqlc.arg('id') AND account_id = sqlc.arg('account_id')
RETURNING
    id,
    created_at,
    updated_at,
    in_stock,
    orderable,
    reserved;
--
