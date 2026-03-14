-- name: CreateItemIdentifier :one
INSERT INTO item_identifiers 
(
    ean,
    gtin,
    isbn,
    jan,
    mpn,
    nsn,
    upc,
    qr,
    sku,
    account_id,
    item_id
)
VALUES 
(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11
)
RETURNING
    id,
    created_at,
    updated_at,
    ean,
    gtin,
    isbn,
    jan,
    mpn,
    nsn,
    upc,
    qr,
    sku,
    item_id;
--

-- name: DeleteItemIdentifier :exec
DELETE FROM item_identifiers
WHERE id = $1 AND account_id = $2;
--

-- name: GetItemIdentifier :one
SELECT
    id,
    created_at,
    updated_at,
    ean,
    gtin,
    isbn,
    jan,
    mpn,
    nsn,
    upc,
    qr,
    sku,
    item_id
FROM item_identifiers
WHERE id = $1 AND account_id = $2;
--

-- name: ListItemIdentifiers :many
SELECT *
FROM
(
    SELECT
        item_identifiers.id,
        item_identifiers.created_at,
        item_identifiers.updated_at,
        item_identifiers.ean,
        item_identifiers.gtin,
        item_identifiers.isbn,
        item_identifiers.jan,
        item_identifiers.mpn,
        item_identifiers.nsn,
        item_identifiers.upc,
        item_identifiers.qr,
        item_identifiers.sku,
        item_identifiers.item_id
    FROM item_identifiers
    WHERE item_identifiers.account_id = sqlc.arg('account_id')
    AND
        (
            sqlc.narg('created_at_gt')::timestamp IS NULL 
            OR item_identifiers.created_at > sqlc.narg('created_at_gt')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_lt')::timestamp IS NULL 
            OR item_identifiers.created_at < sqlc.narg('created_at_lt')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_gte')::timestamp IS NULL 
            OR item_identifiers.created_at >= sqlc.narg('created_at_gte')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_lte')::timestamp IS NULL 
            OR item_identifiers.created_at <= sqlc.narg('created_at_lte')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_gt')::timestamp IS NULL 
            OR item_identifiers.updated_at > sqlc.narg('updated_at_gt')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_lt')::timestamp IS NULL 
            OR item_identifiers.updated_at < sqlc.narg('updated_at_lt')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_gte')::timestamp IS NULL 
            OR item_identifiers.updated_at >= sqlc.narg('updated_at_gte')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_lte')::timestamp IS NULL 
            OR item_identifiers.updated_at <= sqlc.narg('updated_at_lte')::timestamp
        )
    AND 
        (
            sqlc.narg('starting_after')::uuid IS NULL
            OR 
            (
                (
                    sqlc.narg('starting_after_date')::timestamp,
                    sqlc.narg('starting_after')::uuid
                ) > (item_identifiers.created_at, item_identifiers.id)
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
                ) < (item_identifiers.created_at, item_identifiers.id)
            )
        )
    ORDER BY 
    (
        CASE 
            WHEN sqlc.narg('ending_before')::uuid IS NOT NULL 
            THEN (item_identifiers.created_at, item_identifiers.id)
        END
    ) ASC,
    (item_identifiers.created_at, item_identifiers.id) DESC
    LIMIT COALESCE(sqlc.narg('limit')::integer, 10) + 1
)
ORDER BY created_at DESC, id DESC;
--

-- name: ListItemIdentifiersByIds :many
SELECT
	id,
    created_at,
    updated_at,
    ean,
    gtin,
    isbn,
    jan,
    mpn,
    nsn,
    upc,
    qr,
    sku,
    item_id
FROM item_identifiers
WHERE account_id = sqlc.arg('account_id')
    AND id = ANY(sqlc.arg('ids')::uuid[])
ORDER BY created_at DESC, id DESC;
--

-- name: UpdateItemIdentifier :one
UPDATE item_identifiers
SET
    updated_at = sqlc.arg('updated_at')::timestamp,
    ean = COALESCE(sqlc.narg('ean'), ean),
    gtin = COALESCE(sqlc.narg('gtin'), gtin),
    isbn = COALESCE(sqlc.narg('isbn'), isbn),
    jan = COALESCE(sqlc.narg('jan'), jan),
    mpn = COALESCE(sqlc.narg('mpn'), mpn),
    nsn = COALESCE(sqlc.narg('nsn'), nsn),
    upc = COALESCE(sqlc.narg('upc'), upc),
    qr = COALESCE(sqlc.narg('qr'), qr),
    sku = COALESCE(sqlc.narg('sku'), sku)
WHERE id = sqlc.arg('id') AND account_id = sqlc.arg('account_id')
RETURNING
    id,
    created_at,
    updated_at,
    ean,
    gtin,
    isbn,
    jan,
    mpn,
    nsn,
    upc,
    qr,
    sku,
    item_id;
--
