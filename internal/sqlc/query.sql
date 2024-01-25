-- name: GetGithubUser :one
SELECT * FROM GithubUsers
WHERE name = ? LIMIT 1;

-- name: ListGithubUsers :many
SELECT * FROM GithubUsers
ORDER BY name;

-- name: CreateGithubUser :one
INSERT INTO GithubUsers (
  name
) VALUES (
  ?
)
RETURNING *;

-- name: UpdateGithubUser :one
UPDATE GithubUsers
set name = ?
WHERE id = ?
RETURNING *;

-- name: DeleteGithubUser :exec
DELETE FROM GithubUsers
WHERE id = ?;