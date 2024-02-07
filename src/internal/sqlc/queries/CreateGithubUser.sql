-- name: CreateGithubUser :one
INSERT INTO GithubUsers (
  name,
  alph,
  present
) VALUES (
  ?,
  ?,
  ?
)
RETURNING *;