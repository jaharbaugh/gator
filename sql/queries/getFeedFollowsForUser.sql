-- name: GetFeedFollowsForUser :many

SELECT
    feed_follows.*,
    f.name  AS feed_name,
    u.name  AS user_name
FROM feed_follows
INNER JOIN feeds AS f
    ON f.id = feed_follows.feed_id
INNER JOIN users AS u
    ON u.id = feed_follows.user_id
WHERE feed_follows.user_id = $1;