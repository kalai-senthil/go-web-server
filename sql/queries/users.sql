-- name: CreateUser :execresult
INSERT INTO users(id,created_at,updated_at,name) VALUES (?,?,?,?);

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;