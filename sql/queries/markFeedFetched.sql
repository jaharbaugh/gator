-- name: MarkFetchedAt :exec

UPDATE feeds
SET 
last_fetched_at = $1, 
updated_at = $2

WHERE ID = $3;