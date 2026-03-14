-- name: CreateGroup :one
INSERT INTO groups 
(
    description,
    name,
    account_id,
    parent_group_id
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
    description,
    name,
    parent_group_id;
--

-- name: DeleteGroup :exec
DELETE FROM groups
WHERE id = $1 AND account_id = $2;
--

-- name: GetGroup :one
SELECT
    id,
    created_at,
    updated_at,
    description,
    name,
    parent_group_id
FROM groups
WHERE id = $1 AND account_id = $2;
--

-- name: ListGroups :many
SELECT *
FROM 
(
    SELECT
        groups.id,
        groups.created_at,
        groups.updated_at,
        groups.description,
        groups.name,
        groups.parent_group_id
    FROM groups
    WHERE groups.account_id = sqlc.arg('account_id')
    AND
        (
            sqlc.narg('created_at_gt')::timestamp IS NULL 
            OR groups.created_at > sqlc.narg('created_at_gt')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_lt')::timestamp IS NULL 
            OR groups.created_at < sqlc.narg('created_at_lt')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_gte')::timestamp IS NULL 
            OR groups.created_at >= sqlc.narg('created_at_gte')::timestamp
        )
    AND 
        (
            sqlc.narg('created_at_lte')::timestamp IS NULL 
            OR groups.created_at <= sqlc.narg('created_at_lte')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_gt')::timestamp IS NULL 
            OR groups.updated_at > sqlc.narg('updated_at_gt')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_lt')::timestamp IS NULL 
            OR groups.updated_at < sqlc.narg('updated_at_lt')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_gte')::timestamp IS NULL 
            OR groups.updated_at >= sqlc.narg('updated_at_gte')::timestamp
        )
    AND 
        (
            sqlc.narg('updated_at_lte')::timestamp IS NULL 
            OR groups.updated_at <= sqlc.narg('updated_at_lte')::timestamp
        )
    AND 
        (
            sqlc.narg('starting_after')::typeid IS NULL
            OR 
            (
                (
                    sqlc.narg('starting_after_date')::timestamp,
                    sqlc.narg('starting_after')::typeid
                ) > (groups.created_at, groups.id)
            )
        )
    AND 
        (
            sqlc.narg('ending_before')::typeid IS NULL
            OR 
            (
                (
                    sqlc.narg('ending_before_date')::timestamp,
                    sqlc.narg('ending_before')::typeid
                ) < (groups.created_at, groups.id)
            )
        )
    AND 
        (
            sqlc.narg('description')::text IS NULL 
            OR groups.description ~~* CONCAT('%', sqlc.narg('description')::text, '%')
        )
    AND 
        (
            sqlc.narg('name')::text IS NULL 
            OR groups.name ~~* CONCAT('%', sqlc.narg('name')::text, '%')
        )
    AND 
        (
            sqlc.narg('parent_group_id')::typeid IS NULL 
            OR groups.parent_group_id = sqlc.narg('parent_group_id')::typeid
        )
    ORDER BY 
    (
        CASE 
            WHEN sqlc.narg('ending_before')::typeid IS NOT NULL 
            THEN (groups.created_at, groups.id)
        END
    ) ASC,
    (groups.created_at, groups.id) DESC
    LIMIT COALESCE(sqlc.narg('limit')::integer, 10) + 1
)
ORDER BY created_at DESC, id DESC;
--

-- name: ListGroupsByIds :many
SELECT
	id,
	created_at,
	updated_at,
	description,
	name,
	parent_group_id
FROM groups
WHERE account_id = sqlc.arg('account_id')
    AND id = ANY(sqlc.arg('ids')::typeid[])
ORDER BY created_at DESC, id DESC;
--

-- name: UpdateGroup :one
UPDATE groups
SET
    updated_at = sqlc.arg('updated_at')::timestamp,
    description = COALESCE(sqlc.narg('description'), description),
    name = COALESCE(sqlc.narg('name'), name),
    parent_group_id = COALESCE(sqlc.narg('parent_group_id'), parent_group_id)
WHERE id = sqlc.arg('id') AND account_id = sqlc.arg('account_id')
RETURNING
    id,
    created_at,
    updated_at,
    description,
    name,
    parent_group_id;
--
