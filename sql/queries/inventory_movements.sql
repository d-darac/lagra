-- name: CreateInventoryMovement :one
INSERT INTO inventory_movements
(
    type,     
    account_id,
    inventory_id,
    item_id
)
VALUES 
(
    $1,
    $2,
    $3,
    $4
)
RETURNING
    id,
    created_at,
    updated_at,
    inventory_id,
    item_id,
    type;
--

-- name: GetInventoryMovement :one
SELECT
    inventory_movements.id,
    inventory_movements.created_at,
    inventory_movements.updated_at,
    inventory_movements.inventory_id,
    inventory_movements.item_id,
    inventory_movement_references.id AS inventory_movement_reference_id,
    inventory_movements.type
FROM inventory_movements
LEFT JOIN inventory_movement_references ON inventory_movements.id = inventory_movement_references.inventory_movement_id
WHERE inventory_movements.id = $1 AND inventory_movements.account_id = $2;
--

-- name: ListInventoryMovements :many
SELECT *
FROM 
(
    SELECT
        inventory_movements.id,
        inventory_movements.created_at,
        inventory_movements.updated_at,
        inventory_movements.inventory_id,
        inventory_movements.item_id,
        inventory_movement_references.id AS inventory_movement_reference_id,
        inventory_movements.type
    FROM inventory_movements
    LEFT JOIN inventory_movement_references ON inventory_movements.id = inventory_movement_references.inventory_movement_id
    WHERE inventory_movements.account_id = sqlc.arg('account_id')
    AND
        (
            sqlc.narg('created_at_gt')::timestamp IS NULL 
            OR inventory_movements.created_at > sqlc.narg('created_at_gt')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_lt')::timestamp IS NULL 
            OR inventory_movements.created_at < sqlc.narg('created_at_lt')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_gte')::timestamp IS NULL 
            OR inventory_movements.created_at >= sqlc.narg('created_at_gte')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_lte')::timestamp IS NULL 
            OR inventory_movements.created_at <= sqlc.narg('created_at_lte')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_gt')::timestamp IS NULL 
            OR inventory_movements.updated_at > sqlc.narg('updated_at_gt')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_lt')::timestamp IS NULL 
            OR inventory_movements.updated_at < sqlc.narg('updated_at_lt')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_gte')::timestamp IS NULL 
            OR inventory_movements.updated_at >= sqlc.narg('updated_at_gte')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_lte')::timestamp IS NULL 
            OR inventory_movements.updated_at <= sqlc.narg('updated_at_lte')::timestamp
        )
    AND 
        (
            sqlc.narg('starting_after')::uuid IS NULL
            OR 
            (
                (
                    sqlc.narg('starting_after_date')::timestamp,
                    sqlc.narg('starting_after')::uuid
                ) > (inventory_movements.created_at, inventory_movements.id)
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
                ) < (inventory_movements.created_at, inventory_movements.id)
            )
        )
    AND 
        (
            sqlc.narg('item_id')::uuid IS NULL 
            OR inventory_movements.item_id = sqlc.narg('item_id')::uuid
        )
    AND 
        (
            sqlc.narg('inventory_id')::uuid IS NULL 
            OR inventory_movements.inventory_id = sqlc.narg('inventory_id')::uuid
        )
    AND 
        (
            sqlc.narg('type')::inventory_movement_type IS NULL 
            OR inventory_movements.type = sqlc.narg('type')::inventory_movement_type
        )
    ORDER BY 
    (
        CASE 
            WHEN sqlc.narg('ending_before')::uuid IS NOT NULL 
            THEN (inventory_movements.created_at, inventory_movements.id)
        END
    ) ASC,
    (inventory_movements.created_at, inventory_movements.id) DESC
    LIMIT COALESCE(sqlc.narg('limit')::integer, 10) + 1
)
ORDER BY created_at DESC, id DESC;
--