-- name: DeleteGithubUser :exec
DELETE FROM GithubUsers
WHERE id = ?;