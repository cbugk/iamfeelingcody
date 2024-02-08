-- name: CreateGithubUser :one
INSERT INTO GithubUsers (
  name,
  ralpv,
  present
) VALUES (
  ?,
  ?,
  ?
)
RETURNING *;