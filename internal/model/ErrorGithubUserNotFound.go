package model

import (
)

type ErrorGithubUserNotFound struct{}

func (e *ErrorGithubUserNotFound) Error() string {
	return "Error: Github user not found."
}
