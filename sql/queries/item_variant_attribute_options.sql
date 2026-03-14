-- name: CreateItemVariantAttributeOption :copyfrom
INSERT INTO item_variant_attribute_options
(
    name,
    account_id,
    item_variant_attribute_id
)
VALUES
(
    $1,
    $2,
    $3
);
--

-- name: DeleteItemVariantAttributeOption :exec
DELETE FROM item_variant_attribute_options
WHERE id = $1 AND account_id = $2;
--

-- name: GetItemVariantAttributeOption :one
SELECT
    id,
    created_at,
    updated_at,
    name,
    item_variant_attribute_id
FROM item_variant_attribute_options
WHERE id = $1 AND account_id = $2;
--

-- name: ListItemVariantAttributeOptions :many
SELECT *
FROM
(
    SELECT
        item_variant_attribute_options.id,
        item_variant_attribute_options.created_at,
        item_variant_attribute_options.updated_at,
        item_variant_attribute_options.name,
        item_variant_attribute_options.item_variant_attribute_id
    FROM item_variant_attribute_options
    WHERE item_variant_attribute_options.account_id = sqlc.arg('account_id')
    AND
        (
            sqlc.narg('created_at_gt')::timestamp IS NULL 
            OR item_variant_attribute_options.created_at > sqlc.narg('created_at_gt')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_lt')::timestamp IS NULL 
            OR item_variant_attribute_options.created_at < sqlc.narg('created_at_lt')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_gte')::timestamp IS NULL 
            OR item_variant_attribute_options.created_at >= sqlc.narg('created_at_gte')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_lte')::timestamp IS NULL 
            OR item_variant_attribute_options.created_at <= sqlc.narg('created_at_lte')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_gt')::timestamp IS NULL 
            OR item_variant_attribute_options.updated_at > sqlc.narg('updated_at_gt')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_lt')::timestamp IS NULL 
            OR item_variant_attribute_options.updated_at < sqlc.narg('updated_at_lt')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_gte')::timestamp IS NULL 
            OR item_variant_attribute_options.updated_at >= sqlc.narg('updated_at_gte')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_lte')::timestamp IS NULL 
            OR item_variant_attribute_options.updated_at <= sqlc.narg('updated_at_lte')::timestamp
        )
    AND 
        (
            sqlc.narg('starting_after')::uuid IS NULL
            OR 
            (
                (
                    sqlc.narg('starting_after_date')::timestamp,
                    sqlc.narg('starting_after')::uuid
                ) > (item_variant_attribute_options.created_at, item_variant_attribute_options.id)
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
                ) < (item_variant_attribute_options.created_at, item_variant_attribute_options.id)
            )
        )
    AND 
        (
            sqlc.narg('name')::text IS NULL 
            OR item_variant_attribute_options.name ~~* CONCAT('%', sqlc.narg('name')::text, '%')
        )
    ORDER BY 
    (
        CASE 
            WHEN sqlc.narg('ending_before')::uuid IS NOT NULL 
            THEN (item_variant_attribute_options.created_at, item_variant_attribute_options.id)
        END
    ) ASC,
    (item_variant_attribute_options.created_at, item_variant_attribute_options.id) DESC
    LIMIT COALESCE(sqlc.narg('limit')::integer, 10) + 1
)
ORDER BY created_at DESC, id DESC;
--

-- name: UpdateItemVariantAttributeOption :one
UPDATE item_variant_attribute_options
SET
    updated_at = sqlc.arg('updated_at')::timestamp,
    name = COALESCE(sqlc.narg('name'), name)
WHERE id = sqlc.arg('id') AND account_id = sqlc.arg('account_id')
RETURNING
    id,
    created_at,
    updated_at,
    name,
    item_variant_attribute_id;
--
