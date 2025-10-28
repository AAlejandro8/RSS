-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
      $1,
      $2,
      $3,
      $4,
      $5,
      $6
      ) 
RETURNING *;
--


-- name: GetFeeds :many
SELECT feeds.name AS feed_name,
      feeds.url AS feed_url, 
      users.name AS user_name
FROM feeds
JOIN users ON users.id = feeds.user_id;
--

-- name: GetQueryByURL :one
SELECT * FROM feeds
WHERE url = $1;
--

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = NOW(), updated_at = NOW()
WHERE feeds.id = $1;
--

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST, id ASC
LIMIT 1;


