-- name: GetGithubUser :one
SELECT * FROM GithubUsers
WHERE name = ? LIMIT 1;