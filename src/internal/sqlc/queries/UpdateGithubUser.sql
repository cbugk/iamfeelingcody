-- name: UpdateGithubUser :one
UPDATE GithubUsers
SET
name = ?,
present = ?
WHERE id = ?
RETURNING *;