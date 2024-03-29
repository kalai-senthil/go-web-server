// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: feed_follows.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createFeedFollow = `-- name: CreateFeedFollow :execresult
INSERT INTO feeds_follows(id,created_at,updated_at,user_id,feed_id)
VALUES (?,?,?,?,?)
`

type CreateFeedFollowParams struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    string
	FeedID    string
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createFeedFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
}

const getFeedFollows = `-- name: GetFeedFollows :many
SELECT id, created_at, updated_at, user_id, feed_id FROM feeds_follows WHERE user_id = ?
`

func (q *Queries) GetFeedFollows(ctx context.Context, userID string) ([]FeedsFollow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollows, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedsFollow
	for rows.Next() {
		var i FeedsFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const unfollowFeed = `-- name: UnfollowFeed :exec
DELETE FROM feeds_follows WHERE id = ? AND user_id = ?
`

type UnfollowFeedParams struct {
	ID     string
	UserID string
}

func (q *Queries) UnfollowFeed(ctx context.Context, arg UnfollowFeedParams) error {
	_, err := q.db.ExecContext(ctx, unfollowFeed, arg.ID, arg.UserID)
	return err
}
