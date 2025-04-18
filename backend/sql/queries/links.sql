
-- name: CreateLink :one
INSERT INTO links(
    token,
    created_at,
    updated_at,
    doc_id,
    permission,
    expires_at
)VALUES(
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4
)
RETURNING *;