-- name: UpdateGithubUser :one
UPDATE GithubUsers
set name = ?
WHERE id = ?
RETURNING *;