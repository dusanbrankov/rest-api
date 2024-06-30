-- name: GetUser :one
SELECT * FROM user
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM user
ORDER BY first_name;

-- name: CreateUser :execresult
INSERT INTO user (
	first_name, last_name, email, password
) VALUES (?, ?, ?, ?);

-- name: DeleteUser :exec
DELETE FROM user
WHERE id = ?;

