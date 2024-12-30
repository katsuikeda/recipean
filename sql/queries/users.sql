-- name: CreateUser :one
INSERT INTO
    users (
        email,
        password_hash,
        name
    )
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: DeleteAllUsers :exec
DELETE FROM users;