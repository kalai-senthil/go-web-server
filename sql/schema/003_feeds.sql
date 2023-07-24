-- +goose Up
CREATE TABLE feeds(
    id varchar(36) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT  NOT NULL,
    user_id varchar(36) NOT NULL REFERENCES users(id) ON DELETE CASCADE
);


-- +goose Down

DROP TABLE feeds;