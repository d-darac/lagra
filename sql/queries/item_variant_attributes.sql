-- name: CreateItemVariantAttribute :copyfrom
INSERT INTO item_variant_attributes
(
    name,
    account_id,
    item_id
)
VALUES
(
    $1,
    $2,
    $3
);
--

-- name: DeleteItemVariantAttribute :exec
DELETE FROM item_variant_attributes
WHERE id = $1 AND account_id = $2;
--

-- name: GetItemVariantAttribute :one
SELECT
    id,
    created_at,
    updated_at,
    name,
    item_id
FROM item_variant_attributes
WHERE id = $1 AND account_id = $2;
--

-- name: ListItemVariantAttributes :many
SELECT *
FROM
(
    SELECT
        item_variant_attributes.id,
        item_variant_attributes.created_at,
        item_variant_attributes.updated_at,
        item_variant_attributes.name,
        item_variant_attributes.item_id
    FROM item_variant_attributes
    WHERE item_variant_attributes.account_id = sqlc.arg('account_id')
    AND
        (
            sqlc.narg('created_at_gt')::timestamp IS NULL 
            OR item_variant_attributes.created_at > sqlc.narg('created_at_gt')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_lt')::timestamp IS NULL 
            OR item_variant_attributes.created_at < sqlc.narg('created_at_lt')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_gte')::timestamp IS NULL 
            OR item_variant_attributes.created_at >= sqlc.narg('created_at_gte')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_lte')::timestamp IS NULL 
            OR item_variant_attributes.created_at <= sqlc.narg('created_at_lte')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_gt')::timestamp IS NULL 
            OR item_variant_attributes.updated_at > sqlc.narg('updated_at_gt')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_lt')::timestamp IS NULL 
            OR item_variant_attributes.updated_at < sqlc.narg('updated_at_lt')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_gte')::timestamp IS NULL 
            OR item_variant_attributes.updated_at >= sqlc.narg('updated_at_gte')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_lte')::timestamp IS NULL 
            OR item_variant_attributes.updated_at <= sqlc.narg('updated_at_lte')::timestamp
        )
    AND 
        (
            sqlc.narg('starting_after')::uuid IS NULL
            OR 
            (
                (
                    sqlc.narg('starting_after_date')::timestamp,
                    sqlc.narg('starting_after')::uuid
                ) > (item_variant_attributes.created_at, item_variant_attributes.id)
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
                ) < (item_variant_attributes.created_at, item_variant_attributes.id)
            )
        )
    AND 
        (
            sqlc.narg('name')::text IS NULL 
            OR item_variant_attributes.name ~~* CONCAT('%', sqlc.narg('name')::text, '%')
        )
    ORDER BY 
    (
        CASE 
            WHEN sqlc.narg('ending_before')::uuid IS NOT NULL 
            THEN (item_variant_attributes.created_at, item_variant_attributes.id)
        END
    ) ASC,
    (item_variant_attributes.created_at, item_variant_attributes.id) DESC
    LIMIT COALESCE(sqlc.narg('limit')::integer, 10) + 1
)
ORDER BY created_at DESC, id DESC;
--

-- name: UpdateItemVariantAttribute :one
UPDATE item_variant_attributes
SET
    updated_at = sqlc.arg('updated_at')::timestamp,
    name = COALESCE(sqlc.narg('name'), name)
WHERE id = sqlc.arg('id') AND account_id = sqlc.arg('account_id')
RETURNING
    id,
    created_at,
    updated_at,
    name,
    item_id;
--
