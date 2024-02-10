-- name: CreateGithubUser :one
INSERT INTO GithubUsers (
  name,
  present
) VALUES (
  ?,
  ?
)
RETURNING *;