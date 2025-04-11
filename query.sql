-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
  email, name, password
) VALUES (
$1,$2,$3
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET
  email = COALESCE($1, email),
  name = COALESCE($2, name),
  password = COALESCE($3, password)
WHERE id = $4
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: CheckUserExists :one
SELECT EXISTS (
  SELECT 1
  FROM users
  WHERE id = $1
);

-- name: CountUsers :one
SELECT count(*) FROM users;