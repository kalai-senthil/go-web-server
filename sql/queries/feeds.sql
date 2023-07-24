-- name: CreateFeed :execresult
INSERT INTO feeds(id,created_at,updated_at,name,url,user_id) VALUES (?,?,?,?,?,?);

-- name: DeleteFeed :exec
DELETE FROM feeds
WHERE id = ?;

-- name: GetFeeds :many
SELECT * FROM feeds
WHERE user_id = ?;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at IS NULL,last_fetched_at ASC LIMIT ?;

-- name: MarkFeedAsFetched :execresult
UPDATE feeds SET last_fetched_at = NOW(),updated_at=NOW() WHERE id = ?;