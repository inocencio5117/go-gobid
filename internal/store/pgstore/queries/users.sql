-- name: CreateUser :one
INSERT INTO users ("user_name", "email", "bio", "password_hash")
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: GetUserById :one
SELECT
    id,
    user_name,
    password_hash,
    email,
    bio,
    created_at,
    updated_at
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT
    id,
    user_name,
    password_hash,
    email,
    bio,
    created_at,
    updated_at
FROM users
WHERE email = $1;