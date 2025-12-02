-- name: GetPostsForUsers :many

SELECT
    p.id,
    p.created_at,
    p.updated_at,
    p.title,
    p.url,
    p.description,
    p.published_at,
    p.feed_id
FROM posts p
JOIN feeds f ON p.feed_id = f.id
WHERE f.user_id = $1
ORDER BY p.published_at DESC NULLS LAST,
         p.created_at DESC
LIMIT $2;