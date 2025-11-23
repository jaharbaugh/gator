-- name: GetUser :one

SELECT *
From users
WHERE name = $1;




