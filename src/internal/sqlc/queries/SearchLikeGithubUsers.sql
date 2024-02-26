-- name: SearchLikeGithubUsers :many
SELECT * FROM GithubUsers
WHERE name LIKE '%' || ? || '%'