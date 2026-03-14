-- name: CreateItem :one
INSERT INTO items 
(
    active,
    description,
    name,
    has_variants,
    price_amount,
    price_currency,
    type,
    account_id,
    group_id,
    inventory_id
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
    $10
)
RETURNING
    id,
    created_at,
    updated_at,
    active,
    description,
    group_id,
    has_variants,
    inventory_id,
    name,
    price_amount,
    price_currency,
    type,
    variant;
--

-- name: CreateItemVariants :copyfrom
INSERT INTO items 
(
    active,
    description,
    name,
    price_amount,
    price_currency,
    type,
    variant,
    account_id,
    group_id,
    inventory_id,
    parent_item_id
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
);
--

-- name: DeleteItem :exec
DELETE FROM items
WHERE id = $1 AND account_id = $2;
--

-- name: GetItem :one
SELECT
    items.id,
    items.created_at,
    items.updated_at,
    items.active,
    items.description,
    items.group_id,
    items.has_variants,
    item_identifiers.id AS item_identifiers_id,
    items.inventory_id,
    items.name,
    items.parent_item_id,
    items.price_amount,
    items.price_currency,
    items.type,
    items.variant
FROM items
LEFT JOIN item_identifiers ON items.id = item_identifiers.item_id
WHERE items.id = $1 AND items.account_id = $2;
--

-- name: ListItemIdsByInventory :many
SELECT id
FROM
(
    SELECT
        items.id,
        items.created_at
    FROM items
    WHERE items.account_id = sqlc.arg('account_id') AND items.inventory_id = sqlc.arg('inventory_id')::uuid
    AND
        (
            sqlc.narg('starting_after')::uuid IS NULL
            OR 
            (
                (
                    sqlc.narg('starting_after_date')::timestamp,
                    sqlc.narg('starting_after')::uuid
                ) > (items.created_at, items.id)
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
                ) < (items.created_at, items.id)
            )
        )
    ORDER BY 
    (
        CASE 
            WHEN sqlc.narg('ending_before')::uuid IS NOT NULL 
            THEN (items.created_at, items.id)
        END
    ) ASC,
    (items.created_at, items.id) DESC
    LIMIT COALESCE(sqlc.narg('limit')::integer, 10) + 1
)
ORDER BY created_at DESC, id DESC;
--

-- name: ListItems :many
SELECT *
FROM
(
    SELECT
        items.id,
        items.created_at,
        items.updated_at,
        items.active,
        items.description,
        items.group_id,
        items.has_variants,
        item_identifiers.id AS item_identifiers_id,
        items.inventory_id,
        items.name,
        items.parent_item_id,
        items.price_amount,
        items.price_currency,
        items.type,
        items.variant
    FROM items
    LEFT JOIN item_identifiers ON items.id = item_identifiers.item_id
    WHERE items.account_id = sqlc.arg('account_id')
    AND
        (
            sqlc.narg('created_at_gt')::timestamp IS NULL 
            OR items.created_at > sqlc.narg('created_at_gt')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_lt')::timestamp IS NULL 
            OR items.created_at < sqlc.narg('created_at_lt')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_gte')::timestamp IS NULL 
            OR items.created_at >= sqlc.narg('created_at_gte')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_lte')::timestamp IS NULL 
            OR items.created_at <= sqlc.narg('created_at_lte')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_gt')::timestamp IS NULL 
            OR items.updated_at > sqlc.narg('updated_at_gt')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_lt')::timestamp IS NULL 
            OR items.updated_at < sqlc.narg('updated_at_lt')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_gte')::timestamp IS NULL 
            OR items.updated_at >= sqlc.narg('updated_at_gte')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_lte')::timestamp IS NULL 
            OR items.updated_at <= sqlc.narg('updated_at_lte')::timestamp
        )
    AND 
        (
            sqlc.narg('starting_after')::uuid IS NULL
            OR 
            (
                (
                    sqlc.narg('starting_after_date')::timestamp,
                    sqlc.narg('starting_after')::uuid
                ) > (items.created_at, items.id)
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
                ) < (items.created_at, items.id)
            )
        )
    AND
        (
            sqlc.narg('active')::boolean IS NULL
            OR items.active = sqlc.narg('active')::boolean
        )
    AND 
        (
            sqlc.narg('description')::text IS NULL 
            OR items.description ~~* CONCAT('%', sqlc.narg('description')::text, '%')
        )
    AND 
        (
            sqlc.narg('group_id')::uuid IS NULL 
            OR items.group_id = sqlc.narg('group_id')::uuid
        )
    AND
        (
            sqlc.narg('has_variants')::boolean IS NULL 
            OR items.has_variants = sqlc.narg('has_variants')::boolean
        )
    AND 
        (
            sqlc.narg('inventory_id')::uuid IS NULL 
            OR items.inventory_id = sqlc.narg('inventory_id')::uuid
        )
    AND 
        (
            sqlc.narg('name')::text IS NULL 
            OR items.name ~~* CONCAT('%', sqlc.narg('name')::text, '%')
        )
    AND 
        (
            sqlc.narg('parent_item_id')::uuid IS NULL 
            OR items.parent_item_id = sqlc.narg('parent_item_id')::uuid
        )
    AND 
        (
            sqlc.narg('price_amount')::integer IS NULL 
            OR items.price_amount = sqlc.narg('price_amount')::integer
        )
    AND 
        (
            sqlc.narg('price_currency')::currency IS NULL 
            OR items.price_currency = sqlc.narg('price_currency')::currency
        )
    AND
        (
            sqlc.narg('type')::item_type IS NULL 
            OR items.type = sqlc.narg('type')::item_type
        )
    AND
        (
            sqlc.narg('variant')::boolean IS NULL 
            OR items.variant = sqlc.narg('variant')::boolean
        )
    ORDER BY 
    (
        CASE 
            WHEN sqlc.narg('ending_before')::uuid IS NOT NULL 
            THEN (items.created_at, items.id)
        END
    ) ASC,
    (items.created_at, items.id) DESC
    LIMIT COALESCE(sqlc.narg('limit')::integer, 10) + 1
)
ORDER BY created_at DESC, id DESC;
--

-- name: ListItemsByIds :many
SELECT
    items.id,
    items.created_at,
    items.updated_at,
    items.active,
    items.description,
    items.group_id,
    items.has_variants,
    item_identifiers.id AS item_identifiers_id,
    items.inventory_id,
    items.name,
    items.parent_item_id,
    items.price_amount,
    items.price_currency,
    items.type,
    items.variant
FROM items
LEFT JOIN item_identifiers ON items.id = item_identifiers.item_id
WHERE items.account_id = sqlc.arg('account_id')
    AND items.id = ANY(sqlc.arg('ids')::uuid[])
ORDER BY items.created_at DESC, items.id DESC;
--

-- name: UpdateItem :one
UPDATE items
SET
    updated_at = sqlc.arg('updated_at')::timestamp,
    active = COALESCE(sqlc.narg('active')::boolean, active),
    description = COALESCE(sqlc.narg('description'), description),
    group_id = COALESCE(sqlc.narg('group_id'), group_id),
    has_variants = COALESCE(sqlc.narg('has_variants'), has_variants),
    inventory_id = COALESCE(sqlc.narg('inventory_id'), inventory_id),
    name = COALESCE(sqlc.narg('name'), name),
    price_amount = COALESCE(sqlc.narg('price_amount'), price_amount),
    price_currency = COALESCE(sqlc.narg('price_currency'), price_currency)
FROM
    items AS i
    LEFT JOIN item_identifiers ON i.id = item_identifiers.item_id
WHERE
    i.id = sqlc.arg('id')
    AND i.account_id = sqlc.arg('account_id')
    AND i.id = items.id
RETURNING
    items.id,
    items.created_at,
    items.updated_at,
    items.active,
    items.description,
    items.group_id,
    items.has_variants,
    item_identifiers.id AS item_identifiers_id,
    items.inventory_id,
    items.name,
    items.parent_item_id,
    items.price_amount,
    items.price_currency,
    items.type,
    items.variant;
--
