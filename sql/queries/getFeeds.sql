-- name: GetFeeds :many
SELECT
  f.name,      -- feed name
  f.url,       -- feed URL
  u.name       -- user name
FROM feeds AS f
JOIN users AS u ON f.user_id = u.id;