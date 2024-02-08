-- name: HighestRalpvUser :one
SELECT * FROM GithubUsers
ORDER BY ralpv DESC
LIMIT 1