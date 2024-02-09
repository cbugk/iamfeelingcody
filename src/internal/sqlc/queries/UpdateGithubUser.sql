-- name: UpdateGithubUser :one
UPDATE GithubUsers
SET
name = ?,
ralpv = ?,
present = ?
WHERE id = ?
RETURNING *;