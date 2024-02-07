-- name: ListGithubUsers :many
SELECT * FROM GithubUsers
ORDER BY name;