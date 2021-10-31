-- name: FindUser :one
SELECT *
FROM "user"
WHERE id = $1
LIMIT 1;

-- name: CountUsersByUsername :one
SELECT count(id)
FROM "user"
WHERE username = $1;

-- name: CountUsersByEmail :one
SELECT count(id)
FROM "user"
WHERE email = $1;

-- name: CreateUser :one
INSERT INTO "user" (username, first_name, last_name, email, phone)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateUser :one
UPDATE "user"
SET username = $2, first_name = $3, last_name = $4, email = $5, phone = $6
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM "user"
WHERE id = $1;
