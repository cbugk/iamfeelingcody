-- name: HighestRalpvUser :one
SELECT * FROM GithubUsers
ORDER BY alph DESC
LIMIT 1