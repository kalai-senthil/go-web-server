-- name: CreateUser :execresult
INSERT INTO users(id,created_at,updated_at,name,api_key) VALUES (?,?,?,?,?);

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;

-- name: GetUser :one
SELECT * FROM users
WHERE api_key = ? LIMIT 1;