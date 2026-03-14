-- name: CreateInventoryMovementReference :one
INSERT INTO inventory_movement_references 
(
    type,  
    value,
    account_id,
    inventory_movement_id
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
    inventory_movement_id,
    type,
    value,
    account_id;
--

-- name: GetInventoryMovementReference :many
SELECT
    id,
    created_at,
    updated_at,
    inventory_movement_id,
    type,
    value
FROM inventory_movement_references
WHERE id = $1 AND account_id = $2;

--

-- name: ListInventoryMovementReferences :many
SELECT *
FROM 
(
    SELECT
        inventory_movement_references.id,
        inventory_movement_references.created_at,
        inventory_movement_references.updated_at,
        inventory_movement_references.inventory_movement_id,
        inventory_movement_references.type,
        inventory_movement_references.value
    FROM inventory_movement_references
    WHERE inventory_movement_references.account_id = sqlc.arg('account_id')
    AND
        (
            sqlc.narg('created_at_gt')::timestamp IS NULL 
            OR inventory_movement_references.created_at > sqlc.narg('created_at_gt')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_lt')::timestamp IS NULL 
            OR inventory_movement_references.created_at < sqlc.narg('created_at_lt')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_gte')::timestamp IS NULL 
            OR inventory_movement_references.created_at >= sqlc.narg('created_at_gte')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_lte')::timestamp IS NULL 
            OR inventory_movement_references.created_at <= sqlc.narg('created_at_lte')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_gt')::timestamp IS NULL 
            OR inventory_movement_references.updated_at > sqlc.narg('updated_at_gt')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_lt')::timestamp IS NULL 
            OR inventory_movement_references.updated_at < sqlc.narg('updated_at_lt')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_gte')::timestamp IS NULL 
            OR inventory_movement_references.updated_at >= sqlc.narg('updated_at_gte')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_lte')::timestamp IS NULL 
            OR inventory_movement_references.updated_at <= sqlc.narg('updated_at_lte')::timestamp
        )
    AND 
        (
            sqlc.narg('starting_after')::uuid IS NULL
            OR 
            (
                (
                    sqlc.narg('starting_after_date')::timestamp,
                    sqlc.narg('starting_after')::uuid
                ) > (inventory_movement_references.created_at, inventory_movement_references.id)
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
                ) < (inventory_movement_references.created_at, inventory_movement_references.id)
            )
        )
    AND 
        (
            sqlc.narg('inventory_movement_id')::uuid IS NULL 
            OR inventory_movement_references.inventory_movement_id = sqlc.narg('inventory_movement_id')::uuid
        )
    AND 
        (
            sqlc.narg('type')::inventory_movement_reference_type IS NULL 
            OR inventory_movement_references.type = sqlc.narg('type')::inventory_movement_reference_type
        )
    AND 
        (
            sqlc.narg('value')::text IS NULL 
            OR inventory_movement_references.value ~~* CONCAT('%', sqlc.narg('value')::text, '%')
        )
    ORDER BY 
    (
        CASE 
            WHEN sqlc.narg('ending_before')::uuid IS NOT NULL 
            THEN (inventory_movement_references.created_at, inventory_movement_references.id)
        END
    ) ASC,
    (inventory_movement_references.created_at, inventory_movement_references.id) DESC
    LIMIT COALESCE(sqlc.narg('limit')::integer, 10) + 1
)
ORDER BY created_at DESC, id DESC;
--