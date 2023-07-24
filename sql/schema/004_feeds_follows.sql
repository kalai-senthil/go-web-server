-- +goose Up
CREATE TABLE feeds_follows(
    id varchar(36) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id varchar(36) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id varchar(36) NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    UNIQUE(user_id,feed_id)
);


-- +goose Down

DROP TABLE feeds_follows;